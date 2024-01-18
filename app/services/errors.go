package services

import (
	"fmt"
)

type CreateError struct {
	Model string
}

func (e CreateError) Error() string {
	return fmt.Sprintf("Error creating %s", e.Model)
}

type UpdateError struct {
	Model string
	ID    string
}

func (e UpdateError) Error() string {
	return fmt.Sprintf("Error updating %s with id: %s", e.Model, e.ID)
}

type FindError struct {
	Model string
	ID    string
}

func (e FindError) Error() string {
	return fmt.Sprintf("failed to find %s with id '%s'", e.Model, e.ID)
}
