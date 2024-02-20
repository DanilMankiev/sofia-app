package sofia

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type List struct {
	List_id  int    `json:"-" db:"id"`
	Listname string `json:"listname" binding:"required"`
}
