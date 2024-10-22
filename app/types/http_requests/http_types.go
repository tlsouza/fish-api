package http_types

type CreateFishRequest struct {
	SpeciesName string   `json:"species_name" validate:"required"` // Required species name of the fish
	Description *string  `json:"description,omitempty"`            // Optional description, omitempty allows it to be omitted if nil
	Lifespan    *int     `json:"lifespan,omitempty"`               // Optional lifespan in years
	Length      *float64 `json:"length,omitempty"`                 // Optional length in mm
}
