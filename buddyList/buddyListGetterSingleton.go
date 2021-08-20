package buddyList

import (
	"spotify/entity"
	"spotify/services"
)

var (
	cookie string
)

//
// Set the cookie for the singleton implementation
//
func SetCookie(c string) {
	cookie = c
}

//
// Gets the buddyList using the singleton cookie provided
//
func GetSingletonBuddyList() entity.BuddyList {
	bl := services.GetBuddyList(cookie)
	return bl
}
