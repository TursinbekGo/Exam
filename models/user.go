package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type CreateUser struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UpdateUser struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type UserGetListRequest struct {
	Offset int
	Limit  int
}

type UserGetListResponse struct {
	Count int
	Users []*User
}

type UserProducts struct {
	UserName     string
	UserProducts []ProductUser
}

type ProductUser struct {
	Name         string
	ProductPrice int
	ProductCount int
	TotalPrice   int
	Time         string
}
