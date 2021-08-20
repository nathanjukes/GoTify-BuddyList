package main

import (
	"fmt"
	"log"
	"spotify/buddyList"
)

func main() {
	fmt.Println("start")

	// Set Cookie
	cookie := "AQAZcV6w6B20P-ohE5V-nfFUH7TR29bGImJFSIOqwzWEOCDu0KBKdyjjgfMCOxEFqhazAteezaDpYlBSOYwqroOB3NGViMqK-g1r8qJZeo1XE-DBO7WXU0oktHjgvjUevu0TfJZYgZ4uRcgtM5KTT-I-s7FB6JE8pLJbX1tdXA"

	// Scoped
	sbl := buddyList.NewScopedInstance(cookie)
	bl, err := sbl.BuddyList()

	if err != nil {
		log.Print(err)
	}

	// Singleton
	buddyList.SetCookie(cookie)
	bl2, err := buddyList.GetSingletonBuddyList()

	if err != nil {
		log.Print(err)
	}

	for _, i := range bl2.Friends {
		fmt.Printf("Friend: %s last listened to %s at %v by %s\n", i.User.Name, i.Track.Name, i.Time, i.Track.Artist.Name)
	}

	for _, i := range bl.Friends {
		fmt.Printf("Friend: %s last listened to %s at %v by %s\n", i.User.Name, i.Track.Name, i.Time, i.Track.Artist.Name)
	}

}
