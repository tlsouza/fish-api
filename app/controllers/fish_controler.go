package controllers

import (
	"api/app/services"
	http_types "api/app/types/http_requests"
	"api/pkg/errors"
	"api/pkg/ports/logic"
	"api/pkg/ports/types"
	"fmt"
)

type fishControler struct {
	svc services.IFishService
}

func NewFishControler(fishService services.IFishService) fishControler {
	return fishControler{svc: fishService}
}

func (fc *fishControler) AddFish(rd types.RequestData) (interface{}, *errors.HttpError) {
	add_fish_req, _ := logic.Unmarshal[http_types.CreateFishRequest](rd.BodyByte, rd.Ctx)

	id, err := fc.svc.AddFish(add_fish_req)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("internal error trying to create new fish"), 500)
	}

	return id, nil
}
