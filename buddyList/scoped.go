package buddyList

import (
	"spotify/entity"
	"spotify/services"
)

type spotifyBuddyList struct {
	cookie string
}

//
// Instantiating a new object with the provided cookie for a scoped implementation
//
func NewScopedInstance(cookie string) spotifyBuddyList {
	return spotifyBuddyList{cookie: cookie}
}

//
// Returns the BuddyList using the current instantiation's cookie
//
func (sbl *spotifyBuddyList) BuddyList() (entity.BuddyList, error) {
	bl, err := services.GetBuddyList(sbl.cookie)
	return bl, err
}
