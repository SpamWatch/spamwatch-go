package main

import (
	"fmt"

	"github.com/SpamWatch/spamwatch-go"
)

var client = spamwatch.Client("API_KEY", nil)

func main() {
	// Getting the API version
	v, err := client.Version()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(v)

	// Getting some stats
	stats, err := client.Stats()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(stats.TotalBansCount)
}
