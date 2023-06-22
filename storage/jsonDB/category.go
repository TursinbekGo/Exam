package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/uuid"
)

type CategoryRepo struct {
	fileName string
	file     *os.File
}

func NewCategoryRepo(fileName string, file *os.File) *CategoryRepo {
	return &CategoryRepo{
		fileName: fileName,
		file:     file,
	}
}
func (u *CategoryRepo) Create(req *models.CreateCategory) (*models.Category, error) {

	Categories, err := u.read()
	if err != nil {
		return nil, err
	}

	var (
		id       = uuid.New().String()
		Category = models.Category{
			Id:       id,
			Name:     req.Name,
			ParentId: req.ParentId,
		}
	)
	Categories[id] = Category

	err = u.write(Categories)
	if err != nil {
		return nil, err
	}

	return &Category, nil
}
func (u *CategoryRepo) GetById(req *models.CategoryPrimaryKey) (*models.Category, error) {

	Categories, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Categories[req.Id]; !ok {
		return nil, errors.New("Category not found")
	}
	Category := Categories[req.Id]

	return &Category, nil
}
func (u *CategoryRepo) GetAll(req *models.CategoryGetListRequest) (*models.CategoryGetListResponse, error) {

	var resp = &models.CategoryGetListResponse{}
	resp.Categories = []*models.Category{}

	CategoryMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(CategoryMap)
	for _, val := range CategoryMap {
		Categories := val
		resp.Categories = append(resp.Categories, &Categories)
	}

	return resp, nil
}
func (u *CategoryRepo) Update(req *models.UpdateCategory) (*models.Category, error) {

	Categories, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Categories[req.Id]; !ok {
		return nil, errors.New("Category not found")
	}

	Categories[req.Id] = models.Category{
		Id:       req.Id,
		Name:     req.Name,
		ParentId: req.ParentId,
	}

	err = u.write(Categories)
	if err != nil {
		return nil, err
	}
	Category := Categories[req.Id]

	return &Category, nil
}
func (u *CategoryRepo) Delete(req *models.CategoryPrimaryKey) error {

	Categories, err := u.read()
	if err != nil {
		return err
	}

	delete(Categories, req.Id)

	err = u.write(Categories)
	if err != nil {
		return err
	}

	return nil
}
func (u *CategoryRepo) read() (map[string]models.Category, error) {
	var (
		Categorys   []*models.Category
		CategoryMap = make(map[string]models.Category)
	)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &Categorys)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, Category := range Categorys {
		CategoryMap[Category.Id] = *Category
	}

	return CategoryMap, nil
}
func (u *CategoryRepo) write(CategoryMap map[string]models.Category) error {

	var Categorys []models.Category

	for _, val := range CategoryMap {
		Categorys = append(Categorys, val)
	}

	body, err := json.MarshalIndent(Categorys, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
