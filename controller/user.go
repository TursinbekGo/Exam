package controller

import (
	"app/models"
	"errors"
	"fmt"

	"log"
)

func (c *Controller) UserCreate(req *models.CreateUser) (*models.User, error) {
	log.Printf("User create req : %+v\n", req)
	resp, err := c.Strg.User().Create(req)
	if err != nil {
		log.Printf("error while creating User : %+v\n ", err)
		return nil, errors.New("Invalid Data")
	}
	return resp, nil
}
func (c *Controller) UserGetById(req *models.UserPrimaryKey) (*models.User, error) {
	log.Printf("User GetById req : %+v\n", req)
	resp, err := c.Strg.User().GetById(req)
	if err != nil {
		log.Printf("error while  user GetById: %+v\n", err)
		return nil, err
	}
	return resp, nil
}
func (c *Controller) UserGetAll(req *models.UserGetListRequest) (*models.UserGetListResponse, error) {
	resp, err := c.Strg.User().GetAll(req)
	if err != nil {
		log.Printf("Error while user GetAll: %+v\n", err)
		return nil, err
	}
	return resp, nil
}
func (c *Controller) UserUpdate(req *models.UpdateUser) (*models.User, error) {
	resp, err := c.Strg.User().Update(req)
	if err != nil {
		log.Printf("error while user Update: %+v\n", err)
		return nil, err
	}
	return resp, nil
}
func (c *Controller) UserDelete(req *models.UserPrimaryKey) error {
	err := c.Strg.User().Delete(req)
	if err != nil {
		log.Printf("error while user delete : %+v\n", err)
		return err
	}
	return nil
}
func (c *Controller) GetUserProducts(req *models.UserPrimaryKey) (resp *models.UserProducts, err error) {
	// get user by id
	user, err := c.Strg.User().GetById(req)
	// logic for getting users order
	orders, err := c.Strg.ShopCart().GetAll(&models.ShopCartGetListRequest{
		Offset: 0,
		Limit:  100,
	})
	var userOrders []models.ShopCart
	for _, order := range orders.Items {
		if order.UserId == req.Id {
			userOrders = append(userOrders, *order)
		}
	}

	resp = &models.UserProducts{}
	resp.UserName = user.Name + " " + user.Surname

	productsCount := map[string]int{}

	for _, order := range userOrders {
		productsCount[order.ProductId] += order.Count
	}

	for key, value := range productsCount {
		// get product
		product, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{
			Id: key,
		})
		if err != nil {
			return nil, err
		}

		resp.UserProducts = append(resp.UserProducts, &models.ProductUser{
			Name:         product.Name,
			ProductCount: value,
			ProductPrice: product.Price,
			TotalPrice:   product.Price * value,
		})

	}

	for _, item := range resp.UserProducts {
		fmt.Println(item)
	}
	return
}
func (c *Controller) UserMoneySpent(req *models.UserPrimaryKey) (resp *models.UserProducts, err error) {
	user, err := c.Strg.User().GetById(req)
	// logic for getting users order
	orders, err := c.Strg.ShopCart().GetAll(&models.ShopCartGetListRequest{
		Offset: 0,
		Limit:  100,
	})
	var userOrders []models.ShopCart
	for _, order := range orders.Items {
		if order.UserId == req.Id {
			userOrders = append(userOrders, *order)
		}
	}

	resp = &models.UserProducts{}
	resp.UserName = user.Name + " " + user.Surname

	productsCount := map[string]int{}

	for _, order := range userOrders {
		productsCount[order.ProductId] += order.Count
	}

	for key, value := range productsCount {
		// get product
		product, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{
			Id: key,
		})
		if err != nil {
			return nil, err
		}
		resp.UserProducts = append(resp.UserProducts, &models.ProductUser{
			Name:       product.Name,
			TotalPrice: product.Price * value,
		})

	}
	for _, item := range resp.UserProducts {
		fmt.Println(item)
	}
	return
}
