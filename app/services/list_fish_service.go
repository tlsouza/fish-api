package services

import (
	"api/app/repository"
	db_types "api/app/types/db"
	http_types "api/app/types/http_types"
)

type listFishService struct {
	repo repository.IFishRepository
}

func (lfs *listFishService) ListFish(query http_types.QueryParams) *http_types.FishListResponse {
	fishList := lfs.repo.ListFish(query)
	return &http_types.FishListResponse{
		Limit: query.Limit,
		Page:  query.Page,
		Fish:  MapFishToFishListItemResponse(fishList),
	}
}

func NewGetListFishService(fish_repo repository.IFishRepository) *listFishService {
	return &listFishService{
		repo: fish_repo,
	}
}

func MapFishToFishListItemResponse(fishSlice []db_types.Fish) []http_types.FishListItemResponse {

	responseSlice := make([]http_types.FishListItemResponse, 0, len(fishSlice))
	for _, fish := range fishSlice {

		response := http_types.FishListItemResponse{
			ID:          fish.ID,
			SpeciesName: fish.SpeciesName,
			Lifespan:    fish.Lifespan,
			Length:      fish.Length,
			IsVerified:  fish.IsVerified,
		}

		responseSlice = append(responseSlice, response)
	}
	return responseSlice
}
