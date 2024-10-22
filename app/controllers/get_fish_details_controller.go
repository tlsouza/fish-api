package controllers

import (
	"api/app/services"
	"api/pkg/errors"
	"api/pkg/ports/types"
	"fmt"
)

type fishDetailControler struct {
	svc services.IGetFishDetailsService
}

func NewFishDetailControler(svc services.IGetFishDetailsService) fishDetailControler {
	return fishDetailControler{svc: svc}
}

func (fc *fishDetailControler) FishDetail(rd types.RequestData) (interface{}, *errors.HttpError) {

	idStr := rd.PathParams["id"]

	fish, err := fc.svc.GetFishDetails(idStr)

	if err != nil {
		if fish == nil {
			return nil, errors.NewHttpError(fmt.Errorf("fish %s not found", idStr), 404)
		}
		return nil, errors.NewHttpError(fmt.Errorf("error getting fish details"), 500)
	}

	return fish, nil
}
