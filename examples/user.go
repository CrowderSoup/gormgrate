package examples

import (
	"github.com/jinzhu/gorm"
)

// User the user of our application
type User struct {
	gorm.Model

	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string
	Profile  Profile
}

// Profile the users profile
type Profile struct {
	gorm.Model

	UserID    uint
	NickName  string `gorm:"size:128"`
	FirstName string `gorm:"size:128"`
	LastName  string `gorm:"size:128"`
	PhotoURL  string `gorm:"size:2000"`
}
