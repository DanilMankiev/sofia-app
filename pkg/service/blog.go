package service

import (

	entity "github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type BlogService struct {
	repo repository.Blog
}

func newBlogService( repo repository.Blog) *BlogService{
	return &BlogService{repo:repo}
}

func (bs *BlogService) CreateBlog(input entity.CreateBlog) (int,error){
	return bs.repo.CreateBlog(input)
}

func (bs *BlogService) GetAllBlog() ([]entity.Blog,error) {
	return bs.repo.GetAllBlog()
}

func (bs *BlogService) DeleteBlog(id int) error{
   return bs.repo.DeleteBlog(id)
}

func (bs *BlogService) UpdateBlog(id int,input entity.UpdateBlog) error {
   return bs.repo.UpdateBlog(id,input)
}

func (bs *BlogService) GetBlogById(id int) (entity.Blog,error){
	return bs.repo.GetBlogById(id)
}