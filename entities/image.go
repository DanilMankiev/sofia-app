package entity

type ImageInput struct {
	Product_id int    `json:"product_id"`
	Image      string `json:"img"`
}

type ImageInputBlog struct {
	Blog_id int `json:"blog_id"`
	Image string `json:"img"`
}

type ImageOutput struct {
	Url string `json:"url"`
}
