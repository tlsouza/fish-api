package services

import http_types "api/app/types/http_requests"

type IFishService interface {
	AddFish(http_types.CreateFishRequest) (*string, error)
}
