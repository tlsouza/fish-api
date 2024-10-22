package repository

import (
	db_types "api/app/types/db"
	http_types "api/app/types/http_requests"
	"fmt"
	"sort"
	"sync"
)

var lockProductRepository = &sync.Mutex{}
var fishRepoSingleton *FishRepository

// InMemoryRepository struct implementing Repository interface
type FishRepository struct {
	fish []db_types.Fish
	mu   sync.Mutex // to ensure thread-safe operations
}

// NewInMemoryRepository creates a new InMemoryRepository
func FishRepoInstance() *FishRepository {
	if fishRepoSingleton == nil {
		lockProductRepository.Lock()
		defer lockProductRepository.Unlock()
		if fishRepoSingleton == nil {
			fishRepoSingleton = &FishRepository{
				fish: []db_types.Fish{},
			}
		}
	}
	return fishRepoSingleton
}

// Save adds a new Product to the repository
func (fr *FishRepository) Save(fish db_types.Fish) (*string, error) {
	fr.mu.Lock()
	defer fr.mu.Unlock()

	// Add the Product to the slice
	fr.fish = append(fr.fish, fish)
	return &fish.ID, nil
}

func (fr *FishRepository) GetFishDetail(id string) (*db_types.Fish, error) {
	fr.mu.Lock()
	defer fr.mu.Unlock()

	for _, fish := range fr.fish {
		// Find by Id condering only non-deleted fish
		if fish.ID == id && !fish.IsDeleted {
			return &fish, nil
		}
	}
	return nil, fmt.Errorf("recordNotFound")
}

func (fr *FishRepository) ListFish(query http_types.QueryParams) []db_types.Fish {
	fr.mu.Lock()
	defer fr.mu.Unlock()

	// Filter out fish that are marked as deleted
	activeFish := []db_types.Fish{}
	for _, fish := range fr.fish {
		if !fish.IsDeleted {
			activeFish = append(activeFish, fish)
		}
	}

	// Calculate the offset based on the page
	offset := (query.Page - 1) * query.Limit

	// Handle out of bounds offsets
	if offset > len(activeFish) {
		return []db_types.Fish{} // return an empty slice if the offset is out of range
	}

	// Determine the end index based on limit and offset
	end := offset + query.Limit
	if end > len(activeFish) {
		end = len(activeFish) // adjust if end exceeds the number of active fish
	}

	result := orderFishListByDate(activeFish[offset:end], query.OrderByCreatedAt, query.Asc)
	// return the paginated and filtered slice

	return result
}

func orderFishListByDate(fishList []db_types.Fish, orderByDate bool, Asc bool) []db_types.Fish {
	if orderByDate {
		sort.Slice(fishList, func(i, j int) bool {
			a := fishList[i]
			b := fishList[j]

			if Asc {
				return a.CreatedAt.Before(b.CreatedAt)
			}

			return a.CreatedAt.After(b.CreatedAt)

		})
	}
	return fishList
}
