package Model

import (
	"fmt"
)

type User struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Age       uint16  `json:"age"`
	Email     string  `json:"email"`
	Money     float64 `json:"money"`
	AvgGrades float64 `json:"grades"`
	Happiness float64 `json:"happiness"`
}

func (u *User) GetFullInformation() string {
	return fmt.Sprintf("Welcome, %s, to the HomePage!\nyour info:\nAge: %d\nEmail: %s\n",
		u.Name, u.Age, u.Email)
}

func (u *User) ShowBalance() string {
	return fmt.Sprintf("%s balance equals %.2f\n", u.Name, u.Money)
}

func (u *User) SetNewName(newName string) {
	u.Name = newName
}
