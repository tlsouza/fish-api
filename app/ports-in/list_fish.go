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

	Addfish = types.HttpServerPort{
		Name:       "add_fish",
		Path:       "fish",
		Verb:       types.GET,
		Adapter:    adapters.Fiber,
		Controller: listFishController.ListFish,
		Validator:  ListDetailsValidator,
	}

	Addfish.Start()
}

func ListDetailsValidator(rd types.RequestData) *errors.HttpError {

	limit, exists := rd.Query["limit"]

	if exists {
		err := validateNumberFromStr(limit)
		if err != nil {
			return errors.NewHttpError(err, 400)
		}
	}
	page, exists := rd.Query["page"]

	if exists {
		err := validateNumberFromStr(page)
		if err != nil {
			return errors.NewHttpError(err, 400)
		}
	}

	return nil
}

func validateNumberFromStr(numberAsStr string) error {
	numberAsInt, err := strconv.Atoi(numberAsStr)

	if err != nil {
		return fmt.Errorf("invalid limit in query: %s", numberAsStr)
	}

	if numberAsInt <= 0 {
		return fmt.Errorf("invalid  range for limit in query: %s", numberAsStr)
	}

	return nil
}
