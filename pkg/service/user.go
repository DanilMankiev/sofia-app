package service

import (
	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func newUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func(us *UserService) GetUser(uid string) (entity.UserDisplay,error){
	return us.repo.GetUser(uid)
}

func (us *UserService) CreateFavorites(uid string,id int) error {
	return us.repo.CreateFavorites(uid, id)
}

func (us * UserService) GetAllFavorites(uid string) ([]entity.Product,error){
	return us.repo.GetAllFavorites(uid)
}

func (us * UserService) DeleteFavorites(uid string, id int) error{
	return us.repo.DeleteFavorites(uid,id)
}
