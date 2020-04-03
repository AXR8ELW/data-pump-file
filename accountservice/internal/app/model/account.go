package model

package model
type Account struct {
        Id string `json:"id"`
        Name string  `json:"name"`
}

// ToString is a somewhat generic ToString method.
func (a *Account) ToString() string {
	return a.ID + " " + a.Name
}
