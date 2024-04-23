package entity

import "errors"

type Category struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCategoryInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

func (up UpdateCategoryInput) Validate() error {
	if up.Name == nil && up.Description == nil {
		return errors.New("update table no validate")
	}
	return nil
}
