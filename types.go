package spamwatch

import "time"

type SpamWatch struct {
	Token   Tokens
	ApiUrl  string
	TimeOut time.Duration
}

// JSON object returned on /version method
// https://docs.spamwat.ch/?go#misc
type Version struct {
	Major   string `json:"major"`
	Minor   string `json:"minor"`
	Patch   string `json:"patch"`
	Version string `json:"version"`
}

// JSON object returned on /stats method
// https://docs.spamwat.ch/?go#misc
type Stats struct {
	TotalBansCount int `json:"total_ban_count"`
}

// JSON object returned on /banlist methods
// https://docs.spamwat.ch/?go#banlist
type BanList struct {
	Admin   int    `json:"admin,omitempty"`
	Date    int64  `json:"date,omitempty"`
	Id      int64  `json:"id"`
	Reason  string `json:"reason"`
	Message string `json:"message,omitempty"`
}

// JSON object used while adding multiple bans
type AddBans struct {
	Id      int64  `json:"id"`
	Reason  string `json:"reason"`
	Message string `json:"message,omitempty"`
}

// JSON object returned in /tokens methods
// https://docs.spamwat.ch/?go#tokens
type Tokens struct {
	Id         int    `json:"id"`
	Permission string `json:"permission"`
	Retired    bool   `json:"retired"`
	Token      string `json:"token"`
	Userid     int64  `json:"userid"`
}

// JSON object returned when an error occurs
type Error struct {
	Code        int    `json:"code"`
	Description string `json:"error"`
	Until       int64  `json:"until,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

// Struct that handles errors
type ErrorHandler struct {
	Err       *Error
	Spamwatch *SpamWatch
	Method    string
}

// These are the optional parameters of SpamWatch Client
type ClientOpts struct {
	ApiUrl  string
	TimeOut time.Duration
}
