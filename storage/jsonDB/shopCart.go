package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type ShopCartRepo struct {
	fileName string
	file     *os.File
}

func NewShopCartRepo(fileName string, file *os.File) *ShopCartRepo {
	return &ShopCartRepo{
		fileName: fileName,
		file:     file,
	}
}

func (o *ShopCartRepo) Create(req *models.CreateShopCart) (*models.ShopCart, error) {
	items, err := o.read()
	if err != nil {
		return nil, err
	}
	var (
		item = models.ShopCart{
			ProductId: req.ProductId,
			UserId:    req.UserId,
			Count:     req.Count,
			Status:    req.Status,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
	)
	items[item.UserId] = append(items[item.UserId], item)
	err = o.write(items)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
func (u *ShopCartRepo) GetById(req *models.ShopCartprimarykey) (*models.ShopCart, error) {
	var item models.ShopCart

	items, err := u.read()
	if err != nil {
		return nil, err
	}
	if _, ok := items[req.UserId]; !ok {
		return nil, errors.New("item not  found")
	}

	for _, val := range items[req.UserId] {
		if val.ProductId == req.ProductId {
			item = val
		}

	}
	return &item, nil
}

func (u *ShopCartRepo) GetAll(req *models.ShopCartGetListRequest) (*models.ShopCartGetListResponse, error) {
	var resp = &models.ShopCartGetListResponse{}
	resp.Items = []*models.ShopCart{}

	ShopCartMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(ShopCartMap) //len
	for _, val := range ShopCartMap {
		user := val
		resp.Items = append(resp.Items, &user)
	}
}

// 	return resp, nil
// }
// func (u *ShopCartRepo) Update(req *models.UpdateShopCart) (*models.ShopCart, error) {
// 	var resp models.ShopCart
// 	items, err := u.read()
// 	if err != nil {
// 		return nil, err
// 	}
// 	if _, ok := items[req.UserId]; !ok {
// 		return nil, errors.New("ShopCart not found!")
// 	}

// 	items[req.UserId] = models.ShopCart{
// 		ProductId: req.ProductId,
// 		UserId:    req.UserId,
// 		Count:     req.Count,
// 		Status:    req.Status,
// 		Time:
// 	}
// }
// func (u *ShopCartRepo) Delete(req *models.ShopCartprimarykey) error {
// 	items, err := u.read()
// 	if err != nil {
// 		return err
// 	}

// 	err = u.write(items)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
func (u *ShopCartRepo) read() (map[string][]models.ShopCart, error) {
	var (
		items       []*models.ShopCart
		ShopCartMap = make(map[string][]models.ShopCart)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v\n", err)
		return nil, err
	}

	err = json.Unmarshal(data, &items)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	for _, item := range items {
		ShopCartMap[item.UserId] = append(ShopCartMap[item.UserId], *item)
	}

	return ShopCartMap, nil
}
func (u *ShopCartRepo) write(ShopCartMap map[string][]models.ShopCart) error {
	var items []models.ShopCart

	for _, val := range ShopCartMap {
		items = append(items, val...)
	}
	body, err := json.MarshalIndent(items, "", "   ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}
