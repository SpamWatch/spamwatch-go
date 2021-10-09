package main

import (
	"fmt"

	"github.com/SpamWatch/spamwatch-go"
)

var client = spamwatch.Client("API_KEY", nil)

func main() {
	// Getting a specific ban
	ban, err := client.GetBan(777000)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ban)

	// Getting all bans
	bans, err := client.GetBans()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(bans)

	// Getting a list of banned ids
	bannedIds, err := client.GetBansMin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(bannedIds)

	// Adding a ban
	client.AddBan(777000, "reason", "message")

	// Deleting a ban
	client.DeleteBan(777000)
}
