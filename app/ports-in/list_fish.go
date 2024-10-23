package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/errors"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
	"fmt"
	"strconv"
)

var ListFish types.HttpServerPort

func init() {
	listFishController := controllers.NewListFishControler(
		services.NewGetListFishService(
			repository.FishRepoInstance(),
		),
	)

	ListFish = types.HttpServerPort{
		Name:       "list_fish",
		Path:       "fish",
		Verb:       types.GET,
		Adapter:    adapters.Fiber,
		Controller: listFishController.ListFish,
		Validator:  ListDetailsValidator,
	}

	ListFish.Start()
}

func ListDetailsValidator(rd types.RequestData) *errors.HttpError {

	limit, exists := rd.Query["limit"]

	if exists {
		err := validateNumberFromStr(limit)
		if err != nil {
			return errors.NewHttpError(fmt.Errorf("invalid 'limit' in query: %s", limit), 400)
		}
	}
	page, exists := rd.Query["page"]

	if exists {
		err := validateNumberFromStr(page)
		if err != nil {
			return errors.NewHttpError(fmt.Errorf("invalid 'page' in query: %s", page), 400)
		}
	}

	sortParam, exists := rd.Query["sort"]
	if exists {
		err := validateSortQuery(sortParam)
		if err != nil {
			return errors.NewHttpError(fmt.Errorf("query param '%s' is not supported", sortParam), 400)
		}
	}

	return nil
}

func validateNumberFromStr(numberAsStr string) error {
	numberAsInt, err := strconv.Atoi(numberAsStr)

	if err != nil {
		return fmt.Errorf("invalid number in query")
	}

	if numberAsInt <= 0 {
		return fmt.Errorf("invalid  number for limit in query")
	}

	return nil
}

func validateSortQuery(sortQuery string) error {
	if !(sortQuery == "created_at" || sortQuery == "-created_at") {
		return fmt.Errorf("this query param is not suported")
	}
	return nil
}
