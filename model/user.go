package model

import "encoding/json"

type User struct {
	Name      string
	Age       int
	Address   string
	Education string
	Married   string
	Salary    string
	Height    int
	Sex       string
}

func FormatJson(o interface{}) (user User, err error) {
	data, err := json.Marshal(o)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(data, &user)
	return user, err
}
