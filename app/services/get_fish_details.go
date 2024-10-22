package services

import (
	"api/app/repository"
	db_types "api/app/types/db"
	http_types "api/app/types/http_types"
)

type getFishDetailService struct {
	repo repository.IFishRepository
}

func (gfd *getFishDetailService) GetFishDetails(id string) (*http_types.FishDetailResponse, error) {
	fish, err := gfd.repo.GetFishDetail(id)

	if err != nil {
		return nil, err
	}

	return mapFishToFishDetailResponse(fish), nil
}

func NewGetFishDetailService(fish_repo repository.IFishRepository) *getFishDetailService {
	return &getFishDetailService{
		repo: fish_repo,
	}
}

func mapFishToFishDetailResponse(fish *db_types.Fish) *http_types.FishDetailResponse {
	return &http_types.FishDetailResponse{
		ID:          fish.ID,
		SpeciesName: fish.SpeciesName,
		Description: fish.Description,
		Lifespan:    fish.Lifespan,
		Length:      fish.Length,
		CreatedAt:   fish.CreatedAt,
		UpdatedAt:   fish.UpdatedAt,
		ImageURL:    fish.ImageURL,
		IsVerified:  fish.IsVerified,
	}
}
