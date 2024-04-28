package entity
import (

"github.com/lib/pq"
)

type ImageInputProduct struct {
	Product_id int    `json:"product_id"`
	Image      pq.StringArray `json:"img"`
}

type ImageInputBlog struct {
	Blog_id int    `json:"blog_id"`
	Image   pq.StringArray `json:"img"`
}
