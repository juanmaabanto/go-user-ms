package response

import "time"

type UserResponse struct {
	Id          string    `json:"id"`
	UserName    string    `json:"userName"`
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	BirthDate   time.Time `json:"birthDate"`
}
