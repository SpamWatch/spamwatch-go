package main

import (
	"fmt"

	"github.com/SpamWatch/spamwatch-go"
)

var client = spamwatch.Client("API_KEY", nil)

func main() {
	// Getting your own Token
	myToken, err := client.GetSelf()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(myToken)

	// Getting all Tokens
	tokens, err := client.GetTokens()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tokens)

	// Getting a specific Token
	Token, err := client.GetToken(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Token)

	// Getting a Users tokens
	tokens, err = client.GetUserTokens(777000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tokens)

	// Creating a Token
	client.CreateToken(777000, spamwatch.UserPermission)

	// Retiring a specific Token
	client.DeleteToken(1)
}
