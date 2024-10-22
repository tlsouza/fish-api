package services

import (
	"api/app/configs"
	"api/app/repository"
	db_types "api/app/types/db"
	http_client "api/app/types/http_client_types"
	http_types "api/app/types/http_types"
	"api/pkg/ports/logic"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type addFishService struct {
	repo repository.IFishRepository
}

func (fs *addFishService) AddFish(fish_request http_types.UpsertFishRequest) (*string, error) {
	fish := MapCreateFishRequestToFish(&fish_request)
	fish.ImageURL, fish.IsVerified = veriryFishNameAndImage(fish.SpeciesName)

	id, err := fs.repo.Save(fish)

	return id, err
}

func MapCreateFishRequestToFish(request *http_types.UpsertFishRequest) db_types.Fish {
	currentTime := time.Now()
	return db_types.Fish{
		ID:          uuid.NewString(),
		SpeciesName: request.SpeciesName,
		Description: request.Description,
		Lifespan:    request.Lifespan,
		Length:      request.Length,
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
		IsDeleted:   false,
	}
}

func NewAddFishService(fish_repo repository.IFishRepository) *addFishService {
	return &addFishService{
		repo: fish_repo,
	}
}

func veriryFishNameAndImage(name string) (*string, bool) {
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("https://fish-species.p.rapidapi.com/fish_api/fish/%s", name),
		nil,
	)
	req.Header.Add("x-rapidapi-key", configs.X_RAPIDAPI_KEY)
	req.Header.Add("x-rapidapi-host", configs.X_RAPIDAPI_HOST)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, false
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fishes, err := logic.Unmarshal[[]http_client.FishExternalApi](body, req.Context())
	if err != nil {
		return nil, false
	}

	if len(fishes) > 0 {
		for _, url := range fishes[0].ImgSrcSet {
			return &url, true
		}
		return nil, true
	}

	return nil, false
}
