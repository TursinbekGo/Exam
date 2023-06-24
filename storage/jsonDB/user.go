package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

type UserRepo struct {
	fileName string
	file     *os.File
}

func NewUserRepo(fileName string, file *os.File) *UserRepo {
	return &UserRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *UserRepo) Create(req *models.CreateUser) (*models.User, error) {
	users, err := u.read()
	if err != nil {
		return nil, err
	}
	var (
		id   = uuid.New().String()
		user = models.User{
			Id:      id,
			Name:    req.Name,
			Surname: req.Surname,
			Balance: req.Balance,
		}
	)
	users[id] = user

	err = u.write(users)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (u *UserRepo) GetById(req *models.UserPrimaryKey) (*models.User, error) {
	users, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := users[req.Id]; !ok {
		return nil, errors.New("User not found")
	}
	user := users[req.Id]
	return &user, nil
}
func (u *UserRepo) GetAll(req *models.UserGetListRequest) (*models.UserGetListResponse, error) {
	var resp = &models.UserGetListResponse{}
	resp.Users = []*models.User{}
	userMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(userMap)
	for _, val := range userMap {
		users := val
		resp.Users = append(resp.Users, &users)
	}
	fmt.Println(resp.Users)
	return resp, nil
}
func (u *UserRepo) Update(req *models.UpdateUser) (*models.User, error) {
	users, err := u.read()
	if err != nil {
		return nil, err
	}
	if _, ok := users[req.Id]; !ok {
		return nil, errors.New("User not found")
	}
	users[req.Id] = models.User{
		Id:      req.Id,
		Name:    req.Name,
		Surname: req.Surname,
		Balance: req.Balance,
	}

	err = u.write(users)
	if err != nil {
		return nil, err
	}
	user := users[req.Id]
	return &user, nil
}
func (u *UserRepo) Delete(req *models.UserPrimaryKey) error {
	users, err := u.read()
	if err != nil {
		return err
	}
	delete(users, req.Id)
	err = u.write(users)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserRepo) read() (map[string]models.User, error) {
	var (
		users   []*models.User
		userMap = make(map[string]models.User)
	)

	data, err := ioutil.ReadFile(u.fileName)

	if err != nil {
		log.Printf("Error while Reading  data : %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Printf("Error while Unmashaling data : %+v", err)
		return nil, err
	}

	for _, user := range users {
		userMap[user.Id] = *user
	}

	return userMap, nil
}
func (u *UserRepo) write(userMap map[string]models.User) error {
	var users []models.User
	for _, val := range userMap {
		users = append(users, val)
	}

	body, err := json.MarshalIndent(users, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (o *UserRepo) GetByUserId(req *models.UserPrimaryKey) (resp []models.ShopCart, err error) {
	var items []models.ShopCart
	body, err := ioutil.ReadFile("./data/order.json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Printf("error while Unmarshal ")
		return nil, err
	}

	for _, item := range items {
		if item.UserId == req.Id {
			resp = append(resp, item)
		}
	}

	return resp, nil
}
