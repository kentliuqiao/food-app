package entity

import (
	"food-app/infrastucture/security"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string     `gorm:"size:100;not null" json:"first_name"`
	LastName  string     `gorm:"size:100;not null" json:"last_name"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Password  string     `gorm:"size:100;not null" json:"password"`
	CreatedAt time.Time  `gorm:"default:default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at;omitempty"`
}

type PublicUser struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:100;not null" json:"first_name"`
	LastName  string `gorm:"size:100;not null" json:"last_name"`
}

// BeforeSave is a gorm hook
func (u *User) BeforeSave() error {
	hashPassword, err := security.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) map[string]string {
	errMsgs := map[string]string{}

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errMsgs["email_required"] = "email required"
		}
		if u.Email != "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				errMsgs["invalid_email"] = "invalid email"
			}
		}
	case "login":
		if u.Password == "" {
			errMsgs["password_required"] = "password required"
		}
		if u.Email == "" {
			errMsgs["email_required"] = "email required"
		}
		if u.Email != "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				errMsgs["invalid_email"] = "invalid email"
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			errMsgs["email_required"] = "email required"
		}
		if u.Email != "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				errMsgs["invalid_email"] = "invalid email"
			}
		}
	default:
		if u.FirstName == "" {
			errMsgs["firstname_required"] = "first name is required"
		}
		if u.LastName == "" {
			errMsgs["lastname_required"] = "last name is required"
		}
		if u.Password == "" {
			errMsgs["password_required"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 6 {
			errMsgs["invalid_password"] = "password should be at least 6 characters"
		}
		if u.Email == "" {
			errMsgs["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err := checkmail.ValidateFormat(u.Email); err != nil {
				errMsgs["invalid_email"] = "please provide a valid email"
			}
		}
	}

	return errMsgs
}

type Users []User

func (users Users) PublicUsers() []*PublicUser {
	publicUsers := make([]*PublicUser, len(users))
	for i, u := range users {
		publicUsers[i] = u.PublicUser()
	}

	return publicUsers
}

func (u *User) PublicUser() *PublicUser {
	return &PublicUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
