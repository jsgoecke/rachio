package rachio

import (
	"encoding/json"
)

// Person data structure - https://rachio.readme.io/docs/publicpersoninfo
type Person struct {
	Email    string `json:"email"`
	ID       string `json:"id"`
	Username string `json:"username"`
	Devices  []struct {
	} `json:"devices"`
	CreateDate int64  `json:"createDate"`
	Deleted    bool   `json:"deleted"`
	FullName   string `json:"fullName"`
}

// The response when a person info is requested
type Response struct {
	ID string `json:"id"`
}

// Returns if the vehicle is mobile enabled for Tesla API control
func (c *Client) GetPerson() (*Person, error) {
	body, err := c.get(BaseURL + "/person/info")
	if err != nil {
		return nil, err
	}

	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	body, err = c.get(BaseURL + "/person/" + response.ID)
	if err != nil {
		return nil, err
	}

	person := &Person{}
	err = json.Unmarshal(body, person)
	if err != nil {
		return nil, err
	}

	return person, nil
}
