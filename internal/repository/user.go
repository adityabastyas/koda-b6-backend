package repository

import "koda-b6-backend1/internal/models"

type UserRepository struct {
	users []models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetAll() []models.User {
	return r.users
}

func (r *UserRepository) FindByEmail(email string) *models.User {
	for _, x := range r.users {
		if x.Email == email {
			return &x
		}
	}
	return nil
}

func (r *UserRepository) Save(user models.User) {
	r.users = append(r.users, user)
}
