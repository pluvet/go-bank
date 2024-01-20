package services

import (
	"fmt"
)

type ErrorCreatingRecordInDB struct {
	Model string
}

// ErrorCreatingRecordInDB
func (e ErrorCreatingRecordInDB) Error() string {
	return fmt.Sprintf("Error creating %s", e.Model)
}

type ErrorUpdatingRecordInDB struct {
	Model string
	ID    string
}

func (e ErrorUpdatingRecordInDB) Error() string {
	return fmt.Sprintf("Error updating %s with id: %s", e.Model, e.ID)
}

type ErrorFindingOneRecordInDB struct {
	Model string
	ID    string
}

func (e ErrorFindingOneRecordInDB) Error() string {
	return fmt.Sprintf("failed to find %s with id '%s'", e.Model, e.ID)
}
