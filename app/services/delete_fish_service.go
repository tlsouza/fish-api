package services

import (
	"api/app/repository"
)

type deleteFishService struct {
	repo repository.IFishRepository
}

func (ds *deleteFishService) DeleteFish(id string) error {
	return ds.repo.DeleteFish(id)
}

func NewDeleteFishService(fish_repo repository.IFishRepository) *deleteFishService {
	return &deleteFishService{
		repo: fish_repo,
	}
}
