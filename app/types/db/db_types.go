package db_types

type Fish struct {
	ID          string   `json:"id"`                    // Unique identifier for the fish
	SpeciesName string   `json:"species_name"`          // Required species name of the fish
	Description *string  `json:"description,omitempty"` // Optional description, omitempty allows it to be omitted if nil
	Lifespan    *int     `json:"lifespan,omitempty"`    // Optional lifespan in years
	Length      *float64 `json:"length,omitempty"`      // Optional length in mm
	CreatedAt   string   `json:"created_at"`            // Creation timestamp
	UpdatedAt   string   `json:"updated_at"`            // Last update timestamp
	IsDeleted   bool     `json:"is_deleted"`            // Indicates if the fish has been soft-deleted
}
