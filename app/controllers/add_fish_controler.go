package controllers

import (
	"api/app/services"
	http_types "api/app/types/http_types"
	"api/pkg/errors"
	"api/pkg/ports/logic"
	"api/pkg/ports/types"
	"fmt"
)

type addFishControler struct {
	svc services.IAddFishService
}

func NewAddFishControler(fishService services.IAddFishService) addFishControler {
	return addFishControler{svc: fishService}
}

func (fc *addFishControler) AddFish(rd types.RequestData) (interface{}, *errors.HttpError) {
	add_fish_req, _ := logic.Unmarshal[http_types.UpsertFishRequest](rd.BodyByte, rd.Ctx)

	id, err := fc.svc.AddFish(add_fish_req)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("internal error trying to create new fish"), 500)
	}

	return id, nil
}
