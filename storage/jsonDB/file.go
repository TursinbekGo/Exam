package jsondb

import (
	"os"

	"app/config"
	"app/storage"
)

type StoreJSON struct {
	user     *UserRepo
	product  *ProductRepo
	category *CategoryRepo
	shopCart *ShopCartRepo
}

func NewconnectionJSON(cfg *config.Config) (storage.StorageI, error) {
	userFile, err := os.Open(cfg.Path + cfg.UserFileName)
	if err != nil {
		return nil, err
	}
	productFile, err := os.Open(cfg.Path + cfg.ProductFileName)
	if err != nil {
		return nil, err
	}
	categoryFile, err := os.Open(cfg.Path + cfg.CategoryFileName)
	if err != nil {
		return nil, err
	}
	shopCartFile, err := os.Open(cfg.Path + cfg.ShopCartFileName)
	if err != nil {
		return nil, err
	}
	return &StoreJSON{
		user:     NewUserRepo(cfg.Path+cfg.UserFileName, userFile),
		product:  NewProductRepo(cfg.Path+cfg.ProductFileName, productFile),
		category: NewCategoryRepo(cfg.Path+cfg.CategoryFileName, categoryFile),
		shopCart: NewShopCartRepo(cfg.Path+cfg.ShopCartFileName, shopCartFile),
	}, nil
}

func (u *StoreJSON) User() storage.UserRepoI {
	return u.user
}
func (p *StoreJSON) Product() storage.ProductRepoI {
	return p.product
}
func (u *StoreJSON) Category() storage.CategoryRepoI {
	return u.category
}
func (u *StoreJSON) ShopCart() storage.ShopCartRepoI {
	return u.shopCart
}
