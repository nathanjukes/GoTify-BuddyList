package buddyList

import (
	"spotify/entity"
	"spotify/services"
)

type spotifyBuddyList struct {
	buddyList entity.BuddyList
	cookie    string
}

//
// Instantiating a new object with the provided cookie for a scoped implementation
//
func NewScopedBuddyList(cookie string) spotifyBuddyList {
	return spotifyBuddyList{cookie: cookie}
}

//
// Returns the BuddyList using the current instantiation's cookie
//
func (sbl *spotifyBuddyList) BuddyList() entity.BuddyList {
	sbl.buddyList = services.GetBuddyList(sbl.cookie)

	return sbl.buddyList
}
