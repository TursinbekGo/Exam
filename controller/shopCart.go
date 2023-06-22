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
