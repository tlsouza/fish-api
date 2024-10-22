package repository

import (
	db_types "api/app/types/db"
	http_types "api/app/types/http_types"
)

type IFishRepository interface {
	Save(fish db_types.Fish) (*string, error)
	GetFishDetail(string) (*db_types.Fish, error)
	ListFish(http_types.QueryParams) []db_types.Fish
}
