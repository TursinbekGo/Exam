package models

type ShopCartprimarykey struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}
type ShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}
type CreateShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}
type UpdateShopCart struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}
type ShopCartGetListRequest struct {
	Offset int
	Limit  int
}
type ShopCartGetListResponse struct {
	Count int
	Items []*ShopCart
}
