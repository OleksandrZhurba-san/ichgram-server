package user

import "golang.org/x/crypto/bcrypt"

func (u *User) BeforeSave() error {
	if u.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password),
			bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		u.Password = string(hashed)
	}
	return nil
}
