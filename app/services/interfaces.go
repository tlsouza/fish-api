package services

import (
	http_types "api/app/types/http_requests"
)

type IAddFishService interface {
	AddFish(http_types.CreateFishRequest) (*string, error)
}

type IGetFishDetailsService interface {
	GetFishDetails(string) (*http_types.FishDetailResponse, error)
}
