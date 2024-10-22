package controllers

import (
	"api/app/services"
	"api/pkg/errors"
	"api/pkg/ports/types"
	"strconv"
)

var DEFEAULT_PAGE, DEFAULT_LIMIT = 1, 5

type listControler struct {
	svc services.IListFishService
}

func NewListFishControler(svc services.IListFishService) listControler {
	return listControler{svc: svc}
}

func (fc *listControler) ListFish(rd types.RequestData) (interface{}, *errors.HttpError) {
	limitAsInt := getIntFromQuery(rd.Query, "limit", DEFAULT_LIMIT)
	pageAsInt := getIntFromQuery(rd.Query, "page", DEFEAULT_PAGE)

	fishResponseList := fc.svc.ListFish(limitAsInt, pageAsInt)

	return fishResponseList, nil
}

func getIntFromQuery(query map[string]string, key string, defaultValue int) int {
	numberAsStr, exists := query[key]
	if exists {
		numberAsInt, _ := strconv.Atoi(numberAsStr)
		return numberAsInt
	}
	return defaultValue
}
