package model

type Accounts struct {
	Account []Account `json:"accounts"`
}
type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
