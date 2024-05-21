package entity

type User struct {
	Id            string `json:"id" db:"id"`
	Name          string `json:"name" binding:"required" db:"name"`
	Surname       string `json:"surname" binding:"required" db:"surname"`
	Email         string `json:"email" binding:"required" db:"email"`
	Phone         string `json:"phone" binding:"required" db:"phone"`
	Password_hash string `json:"password_hash" binding:"required" db:"password_hash"`
}

type SignUpInput struct {
	Name     string `json:"name" binding:"required" default:"example"`
	Surname  string `json:"surname" binding:"required" default:"example"`
	Email    string `json:"email" binding:"required" default:"example@gmail.com"`
	Phone    string `json:"phone" binding:"required" default:"+79376343481"`
	Password string `json:"password" binding:"required" default:"example"`
}

type UserDisplay struct {
	Id      string `json:"id" db:"id"`
	Name    string `json:"name" binding:"required" default:"example"`
	Surname string `json:"surname" binding:"required" default:"example"`
	Email   string `json:"email" binding:"required" default:"example@gmail.com"`
	Phone   string `json:"phone" binding:"required" default:"+79376343481"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	Token string `json:"token" binding:"required"`
}
