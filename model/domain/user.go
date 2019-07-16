package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct save information of an User
type User struct {
	Model
	Fullname string `json:"fullname,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
	Active   bool   `json:"active,omitempty" sql:"default:true"`
}

// BeforeCreate prepare data before create data
func (u *User) BeforeCreate(scope *gorm.Scope) error {
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		scope.SetColumn("Password", string(hash))
	}
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// BeforeUpdate prepare data before update date
func (u *User) BeforeUpdate(scope *gorm.Scope) error {
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MaxCost)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}

	return nil
}

// ComparePassword func
func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
