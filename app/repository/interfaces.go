package repository

import db_types "api/app/types/db"

type IFishRepository interface {
	Save(fish db_types.Fish) (*string, error)
}
