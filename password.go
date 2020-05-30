package golang

import "golang.org/x/crypto/bcrypt"

func PasswordHash(s string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost+1)
	if err != nil {
		return ""
	}
	return string(b)
}

func PasswordVerify(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
