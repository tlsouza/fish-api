package controllers

import (
	"api/app/services"
	"api/pkg/errors"
	"api/pkg/ports/types"
	"fmt"
)

type deleteFishControler struct {
	svc services.IDeleteFishService
}

func NewDeleteFishControler(svc services.IDeleteFishService) deleteFishControler {
	return deleteFishControler{svc: svc}
}

func (df *deleteFishControler) DeleteFish(rd types.RequestData) (interface{}, *errors.HttpError) {
	idStr := rd.PathParams["id"]
	err := df.svc.DeleteFish(idStr)

	if err != nil {
		return nil, errors.NewHttpError(fmt.Errorf("fish %s not found", idStr), 404)
	}

	return "deleted", nil
}
