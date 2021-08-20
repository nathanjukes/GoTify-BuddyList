package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"spotify/entity"
	"time"
)

//
// Returns buddy list using passed in cookie
//
func GetBuddyList(cookie string) (entity.BuddyList, error) {
	a, err := getWebAccessToken(cookie)
	if err != nil {
		return entity.BuddyList{}, err
	}

	bl, err := getBuddyListObject(a)
	if err != nil {
		return entity.BuddyList{}, err
	}

	appendTime(&bl)

	return bl, nil
}

//
// Gets the web access token required for getting the buddy list
//
func getWebAccessToken(cookie string) (string, error) {
	url := "https://open.spotify.com/get_access_token?reason=transport&productType=web_player"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	if cookie == "" {
		return "", fmt.Errorf("no cookie provided")
	}

	// Provide 'Cookie' header
	req.Header.Add("Cookie", fmt.Sprintf("sp_dc=%s", cookie))

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var webAccessToken entity.WebAccessToken
	if err := json.Unmarshal(dat, &webAccessToken); err != nil {
		return "", err
	}

	return webAccessToken.AccessToken, nil
}

//
// Gets the buddy list using an internal Spotify API, requires a web access token with no additional scopes
//
func getBuddyListObject(a string) (entity.BuddyList, error) {
	var buddyList entity.BuddyList

	url := "https://guc-spclient.spotify.com/presence-view/v1/buddylist"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return buddyList, err
	}

	// Provide 'Authorization' header
	if a == "" {
		return buddyList, fmt.Errorf("no web access token provided")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a))

	dat, err := doRequest(req)
	if err != nil {
		return buddyList, err
	}

	// Unmarshal json data into a BuddyList object to return
	if err := json.Unmarshal(dat, &buddyList); err != nil {
		return buddyList, err
	}

	return buddyList, nil
}

//
// BuddyList originally has 'Timestamp', with this method it appends the converted readable time to it
//
func appendTime(bl *entity.BuddyList) {
	for i := 0; i < len(bl.Friends); i++ {
		bl.Friends[i].Time = convertTime(bl.Friends[i].Timestamp)
	}
}

//
// Converts a timestamp to Unix time
//
func convertTime(t int64) time.Time {
	return time.Unix(0, t*int64(time.Millisecond))
}

//
// Execute a given request and return the []byte data from it
//
func doRequest(req *http.Request) ([]byte, error) {
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return []byte{}, err
	}

	dat, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return dat, nil
}
