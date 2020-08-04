package usecase

import "fmt"

func (u *userUsecase) CreateUser() error {
	user, err := u.postgresRepo.GetUser("test@test.com")
	fmt.Println(user)
	if err != nil {
		return err
	}
	return nil
}
