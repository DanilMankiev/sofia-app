package service

import (
	"github.com/DanilMankiev/sofia-app/entities"
	"github.com/DanilMankiev/sofia-app/pkg/repository"
)

type ReviewService struct {
	repo repository.Review
}

func newReviewService (repo repository.Review) *ReviewService {
	return &ReviewService{repo: repo}
}

func(rs * ReviewService) CreateReview(input entity.CreateReview) (int,error){
	return rs.repo.CreateReview(input)
}

func (rs *ReviewService) GetAllReview() ([]entity.Review,error) {
 	return rs.repo.GetAllReview()
}

func (rs *ReviewService) DeleteReview(id int) error{
	return rs.repo.DeleteReview(id)
}

func (rs *ReviewService) UpdateReview(id int,input entity.UpdateReview) error {
	return rs.repo.UpdateReview(id,input)
}