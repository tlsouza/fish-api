package http_types

import "time"

type CreateFishRequest struct {
	SpeciesName string   `json:"species_name" validate:"required"` // Required species name of the fish
	Description *string  `json:"description,omitempty"`            // Optional description, omitempty allows it to be omitted if nil
	Lifespan    *int     `json:"lifespan,omitempty"`               // Optional lifespan in years
	Length      *float64 `json:"length,omitempty"`                 // Optional length in mm
}

type FishDetailResponse struct {
	ID          string    `json:"id"`           // Unique identifier for the fish
	SpeciesName string    `json:"species_name"` // Required species name of the fish
	Description *string   `json:"description"`  // Optional description, omitempty allows it to be omitted if nil
	Lifespan    *int      `json:"lifespan"`     // Optional lifespan in years
	Length      *float64  `json:"length"`       // Optional length in mm
	CreatedAt   time.Time `json:"created_at"`   // Creation timestamp
	UpdatedAt   time.Time `json:"updated_at"`   // Last update timestamp
	ImageURL    *string   `json:"image_url"`    // url for the fish image
}

type FishListItemResponse struct {
	ID          string   `json:"id"`           // Unique identifier for the fish
	SpeciesName string   `json:"species_name"` // Required species name of the fish
	Lifespan    *int     `json:"lifespan"`     // Optional lifespan in years
	Length      *float64 `json:"length"`       // Optional length in mm
}

type FishListResponse struct {
	Fish  []FishListItemResponse `json:"fish"`
	Page  int                    `json:"page"`
	Limit int                    `json:"limit"`
}

type QueryParams struct {
	Limit            int
	Page             int
	OrderByCreatedAt bool
	Asc              bool
}
