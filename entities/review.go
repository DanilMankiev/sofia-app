package entity

type Review struct{
	Id int `json:"id" db:"id"`
	Topic string `json:"topic" binding:"required"` 
	Name string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateReview struct{
	Topic string `json:"topic" binding:"required"` 
	Name string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateReview struct {
	Topic *string `json:"topic"` 
	Name *string `json:"name"`
	Surname *string `json:"surname"`
	Description *string `json:"description"`
}
