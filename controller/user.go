package controller

import (
	"app/models"
	"errors"
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
