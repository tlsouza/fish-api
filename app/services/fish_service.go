package services

import (
	"api/app/repository"
	db_types "api/app/types/db"
	http_types "api/app/types/http_requests"
	"time"

	"github.com/google/uuid"
)

type fishService struct {
	repo repository.IFishRepository
}

func (fs *fishService) AddFish(fish_request http_types.CreateFishRequest) (*string, error) {
	fish := MapCreateFishRequestToFish(&fish_request)
	id, err := fs.repo.Save(fish)

	return id, err
}

func MapCreateFishRequestToFish(request *http_types.CreateFishRequest) db_types.Fish {
	currentTime := time.Now().Format(time.RFC3339)
	return db_types.Fish{
		ID:          uuid.New().String(), // Generate a new UUID for the ID
		SpeciesName: request.SpeciesName,
		Description: request.Description,
		Lifespan:    request.Lifespan,
		Length:      request.Length,
		CreatedAt:   currentTime, // Set Createdat to the current time
		UpdatedAt:   currentTime, // Set UpdatedAt to the current time (same as CreatedAt initialy)
		IsDeleted:   false,       // Set Isdeleted to false
	}
}

func NewFishService(fish_repo repository.IFishRepository) *fishService {
	return &fishService{
		repo: fish_repo,
	}
}
