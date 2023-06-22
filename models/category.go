package models

type CategoryPrimaryKey struct {
	Id string `json:"id"`
}

type Category struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type CreateCategory struct {
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type UpdateCategory struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
}

type CategoryGetListRequest struct {
	Offset int
	Limit  int
}

type CategoryGetListResponse struct {
	Count      int
	Categories []*Category
}
