package services

import (
	"api/app/repository"
	db_types "api/app/types/db"
	http_types "api/app/types/http_types"
	"time"
)

type updateFishService struct {
	repo repository.IFishRepository
}

func (uf *updateFishService) UpdateFish(id string, fish_request http_types.UpsertFishRequest) error {
	currentFish, err := uf.repo.GetFishDetail(id)
	if err != nil {
		return err
	}

	updatedFish := BuildNewFish(currentFish, fish_request)
	if currentFish.SpeciesName != fish_request.SpeciesName {
		updatedFish.ImageURL, updatedFish.IsVerified = veriryFishNameAndImage(updatedFish.SpeciesName)
	}

	return uf.repo.UpdateFish(updatedFish)
}

func NewUpdateFishService(fish_repo repository.IFishRepository) *updateFishService {
	return &updateFishService{
		repo: fish_repo,
	}
}

func BuildNewFish(currentFish *db_types.Fish, newFish http_types.UpsertFishRequest) *db_types.Fish {

	newDbFish := db_types.Fish{
		ID:          currentFish.ID,
		SpeciesName: newFish.SpeciesName,
		Description: newFish.Description,
		CreatedAt:   currentFish.CreatedAt,
		UpdatedAt:   time.Now(),
		Lifespan:    newFish.Lifespan,
		Length:      newFish.Length,
		IsDeleted:   false,
		ImageURL:    currentFish.ImageURL,
		IsVerified:  currentFish.IsVerified,
	}

	return &newDbFish

}
