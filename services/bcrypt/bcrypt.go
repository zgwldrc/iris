package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func HashPassword(plainPass string) (hashedPassword string){
	tempHashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("err[bcrypt.GenerateFromPassword]:", err)
		return ""
	}
	hashedPassword = string(tempHashedPassword)
	return hashedPassword
}

func CompareHashAndPassword (hashedPassword , password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}
