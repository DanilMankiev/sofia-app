package entity

type User struct {
	Id       string `json:"-" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	Surname string `json:"surname" binding:"required" db:"surname"`
	Email string `json:"email" binding:"required" db:"email"`
	Phone string `json:"phone" binding:"required" db:"phone"`
	Password_hash string `json:"password_hash" binding:"required" db:"password_hash"`
}

type SignUpInput struct{
	Name     string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct{
	Email string `json:"email"`
	Password string `json:"password"`
}
