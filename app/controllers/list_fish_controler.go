package controllers

import (
	"api/app/services"
	http_types "api/app/types/http_requests"
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
	orderByDate, Asc := getOrderByDateAndAsc(rd.Query)
	params := http_types.QueryParams{
		Limit:            getIntFromQuery(rd.Query, "limit", DEFAULT_LIMIT),
		Page:             getIntFromQuery(rd.Query, "page", DEFEAULT_PAGE),
		OrderByCreatedAt: orderByDate,
		Asc:              Asc,
	}

	fishResponseList := fc.svc.ListFish(params)

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

func getOrderByDateAndAsc(query map[string]string) (bool, bool) {
	sortParam, exists := query["sort"]
	if exists {
		if sortParam == "created_at" {
			return true, true
		} else {
			return true, false
		}
	}

	return true, true
}
