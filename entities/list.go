package entity

import "errors"

type List struct {
	List_id     int    `json:"list_id" db:"list_id"`
	Listname    string `json:"listname" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateListInput struct {
	Listname    *string `json:"listname"`
	Description *string `json:"description"`
}

func (up UpdateListInput) Validate() error {
	if up.Listname == nil && up.Description == nil {
		return errors.New("update table no validate")
	}
	return nil
}