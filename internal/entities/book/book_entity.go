package book

import "github.com/google/uuid"

// Database model
type Entity struct {
	Uuid uuid.UUID
	Name string
}