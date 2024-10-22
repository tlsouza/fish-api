package http_types

type CreateFishRequest struct {
	SpeciesName string   `json:"species_name" validate:"required"` // Required species name of the fish
	Description *string  `json:"description,omitempty"`            // Optional description, omitempty allows it to be omitted if nil
	Lifespan    *int     `json:"lifespan,omitempty"`               // Optional lifespan in years
	Length      *float64 `json:"length,omitempty"`                 // Optional length in mm
}

type FishDetailResponse struct {
	ID          string   `json:"id"`                    // Unique identifier for the fish
	SpeciesName string   `json:"species_name"`          // Required species name of the fish
	Description *string  `json:"description,omitempty"` // Optional description, omitempty allows it to be omitted if nil
	Lifespan    *int     `json:"lifespan,omitempty"`    // Optional lifespan in years
	Length      *float64 `json:"length,omitempty"`      // Optional length in mm
	CreatedAt   string   `json:"created_at"`            // Creation timestamp
	UpdatedAt   string   `json:"updated_at"`            // Last update timestamp
	ImageURL    *string  `json:"image_url,omitempty"`   // url for the fish image
}
