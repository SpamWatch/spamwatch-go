package spamwatch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// This returns the major, minor and patch version of the API. This endpoint doesn't need a Authorization Token.
// https://docs.spamwat.ch/?go#getting-the-api-version
func (s *SpamWatch) Version() (*Version, error) {
	b, err := s.MakeRequest(http.MethodGet, "version", nil)
	if err != nil {
		return nil, err
	}
	var a = &Version{}
	err = json.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// This returns general stats about the API. Right now this only returns the total ban count.
// https://docs.spamwat.ch/?go#getting-some-stats
func (s *SpamWatch) Stats() (*Stats, error) {
	b, err := s.MakeRequest(http.MethodGet, "stats", nil)
	if err != nil {
		return nil, err
	}
	var a = &Stats{}
	err = json.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Check the ban status of a specific User.
// https://docs.spamwat.ch/?go#getting-a-specific-ban
func (s *SpamWatch) GetBan(id int64) (*BanList, error) {
	b, err := s.MakeRequest(http.MethodGet, fmt.Sprintf("banlist/%d", id), nil)
	if err != nil {
		return nil, err
	}
	var a = &BanList{}
	err = json.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// This returns a list of all Bans.
// https://docs.spamwat.ch/?go#getting-all-bans
func (s *SpamWatch) GetBans() (*[]BanList, error) {
	b, err := s.MakeRequest(http.MethodGet, "banlist", nil)
	if err != nil {
		return nil, err
	}
	var a = []BanList{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// This returns a newline seperated list of all Bans. This method currently ignores the Accept header and will always return a newline seperated list. In the future it might return a JSON with the corresponding content type.
// https://docs.spamwat.ch/?go#getting-a-list-of-banned-ids
func (s *SpamWatch) GetBanMin() ([]int64, error) {
	b, err := s.MakeRequest(http.MethodGet, "banlist/all", nil)
	if err != nil {
		return nil, err
	}
	idstrs := strings.Fields(string(b))
	a := make([]int64, 0, len(idstrs))
	for _, x := range idstrs {
		n, err := strconv.ParseInt(x, 10, 0)
		if err != nil {
			return nil, err
		}
		a = append(a, n)
	}
	return a, nil
}

// This method can be used for adding a ban.
// https://docs.spamwat.ch/?go#adding-a-ban
func (s *SpamWatch) AddBan(id int64, reason string, message string) (bool, error) {
	_, err := s.MakeRequest(http.MethodPost, "banlist", bytes.NewBuffer([]byte(fmt.Sprintf(`[{"id":%d,"reason":"%s","message":"%s"}]`, id, reason, message))))
	if err != nil {
		return false, err
	}
	return true, nil
}

// This method can be used for adding multiple bans at a time.
// https://docs.spamwat.ch/?go#adding-a-ban
func (s *SpamWatch) AddBans(toBan []AddBans) (bool, error) {
	jsonBytes, err := json.Marshal(toBan)
	if err != nil {
		return false, err
	}
	_, err = s.MakeRequest(http.MethodPost, "banlist", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return false, err
	}
	return true, nil
}

// Deleting a ban
// https://docs.spamwat.ch/?go#deleting-a-ban
func (s *SpamWatch) DeleteBan(id int64) (bool, error) {
	_, err := s.MakeRequest(http.MethodDelete, fmt.Sprintf("banlist/%d", id), nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

// This returns the Token the request was made with. Useful for checking the permission Level of the token.
// https://docs.spamwat.ch/?go#getting-your-own-token
func (s *SpamWatch) GetSelf() (*Tokens, error) {
	b, err := s.MakeRequest(http.MethodGet, "tokens/self", nil)
	if err != nil {
		return nil, err
	}
	var a = &Tokens{}
	err = json.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// This returns a list of all Tokens.
// https://docs.spamwat.ch/?go#getting-all-tokens
func (s *SpamWatch) GetTokens() (*[]Tokens, error) {
	b, err := s.MakeRequest(http.MethodGet, "tokens", nil)
	if err != nil {
		return nil, err
	}
	var a = []Tokens{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// This returns a specific Tokens.
// https://docs.spamwat.ch/?go#getting-a-specific-token
func (s *SpamWatch) GetToken(id int) (*Tokens, error) {
	b, err := s.MakeRequest(http.MethodGet, fmt.Sprintf("tokens/%d", id), nil)
	if err != nil {
		return nil, err
	}
	var a = &Tokens{}
	err = json.Unmarshal(b, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// This returns a list of all tokens associated with the specified user id.
// https://docs.spamwat.ch/?go#getting-a-users-tokens
func (s *SpamWatch) GetUserTokens(id int64) (*[]Tokens, error) {
	b, err := s.MakeRequest(http.MethodGet, fmt.Sprintf("tokens/userid/%d", id), nil)
	if err != nil {
		return nil, err
	}
	var a = []Tokens{}
	err = json.Unmarshal(b, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

// Creating a Token
// https://docs.spamwat.ch/?go#creating-a-token
func (s *SpamWatch) CreateToken(userId int64, permission string) (bool, error) {
	_, err := s.MakeRequest(http.MethodPost, "tokens", bytes.NewBuffer([]byte(fmt.Sprintf(`{"id":%d,"permission":"%s"}`, userId, permission))))
	if err != nil {
		return false, err
	}
	return true, nil
}

// This retires a specific Token. The Token won't be able to make any requests anymore.
// https://docs.spamwat.ch/?go#retiring-a-specific-token
func (s *SpamWatch) DeleteToken(tokenId int) (bool, error) {
	_, err := s.MakeRequest(http.MethodDelete, fmt.Sprintf("tokens/%d", tokenId), nil)
	if err != nil {
		return false, err
	}
	return true, nil
}
