package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	"api/pkg/errors"
	"api/pkg/ports/adapters"
	"api/pkg/ports/types"
	"fmt"

	"github.com/google/uuid"
)

var FishDetails types.HttpServerPort

func init() {
	fishControler := controllers.NewFishDetailControler(
		services.NewGetFishDetailService(
			repository.FishRepoInstance(),
		),
	)

	Addfish = types.HttpServerPort{
		Name:       "fish_details",
		Path:       "fish/:id",
		Verb:       types.GET,
		Adapter:    adapters.Fiber,
		Controller: fishControler.FishDetail,
		Validator:  fishDetailsValidator,
	}

	Addfish.Start()
}

func fishDetailsValidator(rd types.RequestData) *errors.HttpError {

	idStr, exists := rd.PathParams["id"]

	if !exists {
		return errors.BadRequestError(fmt.Errorf("id missing in path"))
	}
	err := uuid.Validate(idStr)

	if err != nil {
		fmt.Println(err)
		return errors.BadRequestError(fmt.Errorf("provided id is not a valid UUID"))
	}

	return nil
}
