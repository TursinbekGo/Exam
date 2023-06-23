package controller

import (
	"app/models"
	"errors"
	"log"
)

func (c *Controller) ShopCartCreate(req *models.CreateShopCart) (*models.ShopCart, error) {
	log.Printf("ShopCart create req: %+v\n", req)
	resp, err := c.Strg.ShopCart().Create(req)
	if err != nil {
		log.Printf("error while ShopCart Create: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) ShopCartGetById(req *models.ShopCartprimarykey) (*models.ShopCart, error) {
	log.Printf("ShopCart GetById req: %+v\n", req)
	resp, err := c.Strg.ShopCart().GetById(req)
	if err != nil {
		log.Printf("error while ShopCart GetById: %+v\n", err)
		return nil, errors.New("invalid data")
	}

	return resp, nil
}

func (c *Controller) ShopCartGetAll(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {
	log.Printf("ShopCart  GetAll req: %+v\n", req)
	resp, err := c.Strg.ShopCart().GetAll(req)
	if err != nil {
		log.Printf("error while ShopCart GetAll: %+v\n", err)
		return nil, errors.New("invalid data")
	}
	return resp, nil

}

func (c *Controller) ShopCartUpdate(req *models.UpdateShopCart) (*models.ShopCart, error) {
	log.Printf("ShopCart  ShopCartUpdate req: %+v\n", req)
	resp, err := c.Strg.ShopCart().Update(req)
	if err != nil {
		log.Printf("error while ShopCart ShopCartUpdate: %+v\n", err)
		return nil, errors.New("invalid data")
	}
	return resp, nil

}

func (c *Controller) Delete_ShopCard(req *models.ShopCartprimarykey) error {
	log.Printf("ShopCart  Delete_ShopCard req: %+v\n", req)
	err := c.Strg.ShopCart().Delete(req)
	if err != nil {
		log.Printf("error while ShopCart Delete_ShopCard: %+v\n", err)
		return errors.New("invalid data")
	}
	return nil

}
