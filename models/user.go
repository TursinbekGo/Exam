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

type UserHistory struct {
	UserName    string `json:"user_name"`
	Productname string `json:"product_name"`
	Price       int    `json:"price"`
	Count       int    `json:"count"`
	TotalPrice  int    `json:"total_price"`
	Time        string `json:"time"`
}
