package repository

import (
	db_types "api/app/types/db"
)

type IFishRepository interface {
	Save(fish db_types.Fish) (*string, error)
	GetFishDetail(string) (*db_types.Fish, error)
	ListFish(limit int, page int) []db_types.Fish
}
