package eventchannel

import "fmt"

// User - структура пользователя
type User struct {
	SubscriberDefault
	Username string
}

// NewUser -
func NewUser(username string) *User {
	return &User{
		Username: username,
	}
}

// OnReceive -
func (u *User) OnReceive(msg string) {
	fmt.Printf("Message Got: %s: %s", u.Username, msg)
}

// GetID -
func (u *User) GetID() string {
	return u.Username
}
