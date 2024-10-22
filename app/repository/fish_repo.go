package repository

import (
	db_types "api/app/types/db"
	"fmt"
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

func (fr *FishRepository) ListFish(limit int, page int) []db_types.Fish {
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
	offset := (page - 1) * limit

	// Handle out of bounds offsets
	if offset > len(activeFish) {
		return []db_types.Fish{} // return an empty slice if the offset is out of range
	}

	// Determine the end index based on limit and offset
	end := offset + limit
	if end > len(activeFish) {
		end = len(activeFish) // adjust if end exceeds the number of active fish
	}

	return activeFish[offset:end] // return the paginated and filtered slice
}
