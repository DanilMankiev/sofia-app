package entity

type Blog struct{
	Id int `json:"id" db:"id"`
	Topic string `json:"topic" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateBlog struct{
	Topic string `json:"topic" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateBlog struct{
	Topic *string `json:"topic"`
	Description *string `json:"description"`
}