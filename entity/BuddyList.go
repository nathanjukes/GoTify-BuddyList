package entity

import "time"

type BuddyList struct {
	Friends []Friends `json:"friends"`
}
type User struct {
	URI      string `json:"uri"`
	Name     string `json:"name"`
	ImageURL string `json:"imageUrl"`
}
type Album struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}
type Artist struct {
	URI  string `json:"uri"`
	Name string `json:"name"`
}
type Context struct {
	URI   string `json:"uri"`
	Name  string `json:"name"`
	Index int    `json:"index"`
}
type Track struct {
	URI      string  `json:"uri"`
	Name     string  `json:"name"`
	ImageURL string  `json:"imageUrl"`
	Album    Album   `json:"album"`
	Artist   Artist  `json:"artist"`
	Context  Context `json:"context"`
}
type Friends struct {
	Timestamp int64     `json:"timestamp"`
	Time      time.Time `json:"time"`
	User      User      `json:"user"`
	Track     Track     `json:"track"`
}
