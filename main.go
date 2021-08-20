package main

import (
	"fmt"
	"spotify/buddyList"
)

func main() {
	fmt.Println("start")

	// Set Cookie
	cookie := "AQAZcV6w6B20P-ohE5V-nfFUH7TR29bGImJFSIOqwzWEOCDu0KBKdyjjgfMCOxEFqhazAteezaDpYlBSOYwqroOB3NGViMqK-g1r8qJZeo1XE-DBO7WXU0oktHjgvjUevu0TfJZYgZ4uRcgtM5KTT-I-s7FB6JE8pLJbX1tdXA"

	// Scoped
	sbl := buddyList.NewScopedBuddyList(cookie)
	bl := sbl.BuddyList()

	// Singleton
	buddyList.SetCookie(cookie)
	//bl2 := buddyList.GetSingletonBuddyList()

	for _, i := range bl.Friends {
		fmt.Printf("Friend: %s last listened to %s at %v by %s\n", i.User.Name, i.Track.Name, i.Time, i.Track.Artist.Name)
	}

}
