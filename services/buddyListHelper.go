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
func GetBuddyList(cookie string) entity.BuddyList {
	a := getWebAccessToken(cookie)
	bl := getBuddyListObject(a)
	appendTime(&bl)
	return bl
}

//
// Gets the web access token required for getting the buddy list
//
func getWebAccessToken(cookie string) string {
	url := "https://open.spotify.com/get_access_token?reason=transport&productType=web_player"
	req, _ := http.NewRequest("GET", url, nil)

	// Provide 'Cookie' header
	req.Header.Add("Cookie", fmt.Sprintf("sp_dc=%s", cookie))

	res, _ := (&http.Client{}).Do(req)

	raw, _ := ioutil.ReadAll(res.Body)

	var webAccessToken entity.WebAccessToken
	json.Unmarshal(raw, &webAccessToken)

	return webAccessToken.AccessToken
}

//
// Gets the buddy list using an internal Spotify API, requires a web access token with no additional scopes
//
func getBuddyListObject(a string) entity.BuddyList {
	url := "https://guc-spclient.spotify.com/presence-view/v1/buddylist"
	req, _ := http.NewRequest("GET", url, nil)

	// Provide 'Authorization' header
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a))

	res, _ := (&http.Client{}).Do(req)

	raw, _ := ioutil.ReadAll(res.Body)

	var buddyList entity.BuddyList
	json.Unmarshal(raw, &buddyList)

	return buddyList
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
