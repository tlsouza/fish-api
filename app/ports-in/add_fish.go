package in

import (
	"api/app/controllers"
	"api/app/repository"
	"api/app/services"
	http_types "api/app/types/http_requests"
	"api/pkg/errors"
	"api/pkg/ports/adapters"
	"api/pkg/ports/logic"
	"api/pkg/ports/types"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var Addfish types.HttpServerPort
var validate = validator.New()

func init() {
	fishControler := controllers.NewFishControler(
		services.NewFishService(
			repository.FishRepoInstance(),
		),
	)

	Addfish = types.HttpServerPort{
		Name:       "add_fish",
		Path:       "fish",
		Verb:       types.POST,
		Adapter:    adapters.Fiber,
		Controller: fishControler.AddFish,
		Validator:  addFishValidator,
	}

	Addfish.Start()
}

func addFishValidator(rd types.RequestData) *errors.HttpError {
	req, err := logic.Unmarshal[http_types.CreateFishRequest](rd.BodyByte, rd.Ctx)

	if err != nil {
		return errors.NewHttpError(fmt.Errorf("invalid body structure"), 400)
	}

	err = validate.Struct(&req)
	if err != nil {

		return errors.NewHttpError(logic.GetRequiredFieldError[http_types.CreateFishRequest](
			err,
			&req,
		), 400)
	}

	return nil
}
