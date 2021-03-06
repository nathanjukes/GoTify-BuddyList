package buddyList

import (
	"github.com/nathanjukes/GoTify-BuddyList/entity"
	"github.com/nathanjukes/GoTify-BuddyList/services"
)

var (
	cookie string
)

//
// Gets the buddyList using the singleton cookie provided
//
func GetSingletonBuddyList() (entity.BuddyList, error) {
	bl, err := services.GetBuddyList(cookie)
	return bl, err
}

//
// Set the cookie for the singleton implementation
//
func SetCookie(c string) {
	cookie = c
}
