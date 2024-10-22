package controllers

import (
	"api/app/services"
	http_types "api/app/types/http_types"
	"api/pkg/errors"
	"api/pkg/ports/logic"
	"api/pkg/ports/types"
	"fmt"
)

type updateFishControler struct {
	svc services.IUpdateFishService
}

func NewUpdateFishControler(updateService services.IUpdateFishService) updateFishControler {
	return updateFishControler{svc: updateService}
}

func (uc *updateFishControler) UpdateFish(rd types.RequestData) (interface{}, *errors.HttpError) {
	upsert, _ := logic.Unmarshal[http_types.UpsertFishRequest](rd.BodyByte, rd.Ctx)
	idStr := rd.PathParams["id"]

	err := uc.svc.UpdateFish(idStr, upsert)

	if err != nil {
		if err.Error() == "recordNotFound" {
			return nil, errors.NewHttpError(fmt.Errorf("fish  %s not found", idStr), 404)
		}

		return nil, errors.NewHttpError(fmt.Errorf("internal error trying to update fish"), 500)

	}

	return "updated", nil
}
