<h1 align="center">
  GoTify-BuddyList
</h1>

<h4 align="center">A Go package for <a href="https://github.com/valeriangalliat/spotify-buddylist">Spotify-BuddyList</a></h4>

<p align="center">
  <a href="https://github.com/nathanjukes/GoTify-BuddyList">
      <img src="https://scrutinizer-ci.com/g/pH7Software/pH7-Social-Dating-CMS/badges/build.png?b=master">
  </a>
  <a href="https://pkg.go.dev/github.com/nathanjukes/GoTify-BuddyList">
    <img src="https://img.shields.io/badge/version-v1.0-blue">
  </a>
  <a href="https://github.com/nathanjukes/Sorting-Algorithm-Visualisation/blob/master/LICENSE.md">
    <img src="https://img.shields.io/github/license/Naereen/StrapDown.js.svg">
  </a>
</p>

## What is it & why?

GoTify-BuddyList is a Go package for the official <a href="https://github.com/valeriangalliat/spotify-buddylist">BuddyList package</a> that allows easy integration in Go systems. I created it to aid in a Go Spotify package I'm working on. Spotify-BuddyList adds a way to get the listening activity of the people you follow on Spotify, which isn't included in Spotify's official <a href="https://developer.spotify.com/documentation/web-api/">API</a>.

Instead, Spotify-BuddyList uses internal Spotify endpoints to achieve this in a better manner than plain web-scraping.
## Usage

- `go get github.com/nathanjukes/GoTify-BuddyList`

```go
import "github.com/nathanjukes/GoTify-BuddyList/buddyList"

func main() {
        // sd_pc cookie from Spotify web player, see Data needed section in docs
	myCookie := ""

	// Scoped instance version
	b := buddyList.NewScopedInstance(myCookie)
	bl, err := b.BuddyList()


	// Singleton version
	buddyList.SetCookie(myCookie)
	bl, err := buddyList.GetSingletonBuddyList()
}
```

- BuddyList object structure (inline version)
```go
type BuddyList struct {
	Friends []struct {
		Timestamp int64     `json:"timestamp"`
		Time      time.Time `json:"time"`
		User      struct {
			URI      string `json:"uri"`
			Name     string `json:"name"`
			ImageURL string `json:"imageUrl"`
		} `json:"user"`
		Track struct {
			URI      string `json:"uri"`
			Name     string `json:"name"`
			ImageURL string `json:"imageUrl"`
			Album    struct {
				URI  string `json:"uri"`
				Name string `json:"name"`
			} `json:"album"`
			Artist struct {
				URI  string `json:"uri"`
				Name string `json:"name"`
			} `json:"artist"`
			Context struct {
				URI   string `json:"uri"`
				Name  string `json:"name"`
				Index int    `json:"index"`
			} `json:"context"`
		} `json:"track"`
	} `json:"friends"`
}
```

## Data needed

- '<b>Cookie</b>' - sd_pc cookie from Spotify's Web Player
    - After you log into the Spotify web player, you need to find the cookie named 'sp_dc' and get the value of it
    - This cookie expires every year, so if used in production, it will be best to automate the collection of it

## Change Log

- 1.00 - Released 20/08/21


### License
[MIT](https://github.com/nathanjukes/GoTify-BuddyList/blob/master/LICENSE.md)
