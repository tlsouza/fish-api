package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	http_types "api/app/types/http_types"
	"api/pkg/errors"
	"api/pkg/ports/adapters"
	"api/pkg/ports/logic"
	"api/pkg/ports/types"
	"fmt"

	"github.com/google/uuid"
)

var UpdateFish types.HttpServerPort

func init() {
	updateFishControler := controllers.NewUpdateFishControler(
		services.NewUpdateFishService(
			repository.FishRepoInstance(),
		),
	)

	UpdateFish = types.HttpServerPort{
		Name:       "update_fish",
		Path:       "fish/:id",
		Verb:       types.PUT,
		Adapter:    adapters.Fiber,
		Controller: updateFishControler.UpdateFish,
		Validator:  UpdateFishValidator,
	}

	UpdateFish.Start()
}

func UpdateFishValidator(rd types.RequestData) *errors.HttpError {
	req, err := logic.Unmarshal[http_types.UpsertFishRequest](rd.BodyByte, rd.Ctx)

	if err != nil {
		return errors.NewHttpError(fmt.Errorf("invalid body structure"), 400)
	}

	err = validate.Struct(&req)
	if err != nil {

		return errors.NewHttpError(logic.GetRequiredFieldError[http_types.UpsertFishRequest](
			err,
			&req,
		), 400)
	}

	idStr, exists := rd.PathParams["id"]

	if !exists {
		return errors.BadRequestError(fmt.Errorf("id missing in path"))
	}
	err = uuid.Validate(idStr)

	if err != nil {
		fmt.Println(err)
		return errors.BadRequestError(fmt.Errorf("provided id is not a valid UUID"))
	}

	return nil
}
