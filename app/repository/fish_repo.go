package repository

import (
	db_types "api/app/types/db"
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
