package entity

import(
	"github.com/lib/pq"
)

type Blog struct{
	Id int `json:"id" db:"id"`
	Topic string `json:"topic" db:"topic" binding:"required"`
	Description_preview string `json:"description_preview" db:"description_preview" binding:"required"`
	Description_full string `json:"description_full" db:"description_full" binding:"required"`
	Image_preview string `json:"image_preview" db:"image_preview" binding:"required"`
	Image_all pq.StringArray `json:"image_all" db:"image_all" binding:"required"`
}

type CreateBlog struct{
	Topic string `json:"topic" binding:"required"`
	Description_preview string `json:"description_preview" binding:"required"`
	Description_full string `json:"description_full" binding:"required"`
	Image_preview string `json:"image_preview"`
	Image_all pq.StringArray `json:"image_all"`
}

type UpdateBlog struct{
	Topic *string `json:"topic"`
	Description_preview *string `json:"description_preview"`
	Description_full *string `json:"description_full"`
}