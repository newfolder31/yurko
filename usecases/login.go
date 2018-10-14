package usecases

import "fmt"

func ValidateLogin(email, pass string) error {
	fmt.Println("ValidateLogin ", email, pass)

	//UserRepo.FindByEmailAndPassword(email, pass)

	return nil
}
