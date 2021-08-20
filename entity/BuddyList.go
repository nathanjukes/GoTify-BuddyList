package entity

import "time"

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
