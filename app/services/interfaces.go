package services

import (
	http_types "api/app/types/http_types"
)

type IAddFishService interface {
	AddFish(http_types.UpsertFishRequest) (*string, error)
}

type IGetFishDetailsService interface {
	GetFishDetails(string) (*http_types.FishDetailResponse, error)
}

type IListFishService interface {
	ListFish(http_types.QueryParams) *http_types.FishListResponse
}

type IDeleteFishService interface {
	DeleteFish(string) error
}

type IUpdateFishService interface {
	UpdateFish(string, http_types.UpsertFishRequest) error
}
