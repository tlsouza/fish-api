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

var DeleteFish types.HttpServerPort

func init() {
	deleteFishControler := controllers.NewDeleteFishControler(
		services.NewDeleteFishService(
			repository.FishRepoInstance(),
		),
	)

	DeleteFish = types.HttpServerPort{
		Name:       "delete_fish",
		Path:       "fish/:id",
		Verb:       types.DELETE,
		Adapter:    adapters.Fiber,
		Controller: deleteFishControler.DeleteFish,
		Validator:  deleteFishValidator,
	}

	DeleteFish.Start()
}

func deleteFishValidator(rd types.RequestData) *errors.HttpError {

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
