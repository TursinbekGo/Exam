package controller

import (
	"app/models"
	"errors"
	"fmt"
	"sort"

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

//3 task
func (c *Controller) GetUserProducts(req *models.UserPrimaryKey) (resp *models.UserProducts, err error) {
	// get user by id
	user, err := c.Strg.User().GetById(req)
	// logic for getting users order
	orders, err := c.Strg.ShopCart().GetAll(&models.ShopCartGetListRequest{
		Offset: 0,
		Limit:  10,
	})
	// fmt.Println(orders)
	var userOrders []models.ShopCart
	for _, order := range orders.Items {
		//fmt.Println(order)
		if order.UserId == req.Id {
			//fmt.Println(order)
			if order.Status {
				userOrders = append(userOrders, order)
			}
		}
		//fmt.Println(order)
	}
	// fmt.Println(userOrders)

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

		add := models.ProductUser{
			Name:         product.Name,
			ProductCount: value,
			ProductPrice: product.Price,
			TotalPrice:   product.Price * value,
		}
		resp.UserProducts = append(resp.UserProducts, add)

	}

	return resp, nil
}

//1 task
func (c *Controller) Sort(req *models.ShopCartGetListRequest) ([]models.ShopCart, error) {

	var itemSortByDate []models.ShopCart

	SH_items, err := c.ShopCartGetAll(req)
	if err != nil {
		return nil, err
	}

	for _, item := range SH_items.Items {
		itemSortByDate = append(itemSortByDate, item)
	}
	//	fmt.Println(itemSortByDate)

	//fmt.Println(itemSortByDate)
	sort.Slice(itemSortByDate, func(i, j int) bool {
		return itemSortByDate[i].Time > itemSortByDate[j].Time
	})

	for key, v := range itemSortByDate {
		name, err := c.Strg.User().GetById(&models.UserPrimaryKey{v.UserId})
		if err != nil {
			return nil, err
		}
		product, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{v.ProductId})
		if err != nil {
			return nil, err
		}
		v.UserId = name.Name
		v.ProductId = product.Name
		fmt.Println(v)
		itemSortByDate[key] = v

	}
	fmt.Println(itemSortByDate)
	// sort.Slice(itemSortByDate, func(i, j int) bool {
	// 	return itemSortByDate[i].Time > itemSortByDate[j].Time
	// })
	// for _, value := range itemSortByDate {
	// 	fmt.Println(value)
	// }
	return itemSortByDate, nil
}

//2 task

func (c *Controller) DateSort(req *models.ShopCartGetListRequest) ([]*models.ShopCart, error) {
	var (
		itemFilterByDate []*models.ShopCart
	)
	items, err := c.Strg.ShopCart().GetAll(req)
	if err != nil {
		return nil, err
	}
	//	fmt.Println(items)
	for _, item := range items.Items {
		//fmt.Println(item)
		if item.Status == true {
			if item.Time >= req.From && item.Time < req.To {
				itemFilterByDate = append(itemFilterByDate, &item)
				fmt.Println(item)
			}
		}
	}

	return itemFilterByDate, nil
}

//4 task
func (c *Controller) UserMoneySpent(req *models.UserPrimaryKey) (string, int) {
	var (
		UserTotalSpendMoney int = 0
	)
	history, _ := c.GetUserProducts(&models.UserPrimaryKey{req.Id})
	userName, _ := c.Strg.User().GetById(req)

	for _, val := range history.UserProducts {
		UserTotalSpendMoney += val.TotalPrice
	}
	name := userName.Name + " " + userName.Surname
	return name, UserTotalSpendMoney
}

//5
func (c *Controller) TotalSoldProducts() {
	var (
		totalCount = make(map[string]int)
	)
	prod, err := c.Strg.ShopCart().GetAll(&models.ShopCartGetListRequest{})
	if err != nil {
		log.Printf("Error in get Shopcarts!!!")
		return
	}
	for _, val := range prod.Items {
		prodName, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{val.ProductId})
		if err != nil {
			log.Println("Error while in getbyid product!!!", err)
			return
		}
		if val.Status {
			totalCount[prodName.Name] += val.Count

		}
	}

	for key, val := range totalCount {
		fmt.Println(key, ":", val)
	}

}

//6
func (c *Controller) AvtiveProducts(limit int) (string, int) {
	var (
		totalCount = make(map[string]int)
	)
	prod, err := c.Strg.ShopCart().GetAll(&models.ShopCartGetListRequest{})
	if err != nil {
		log.Printf("Error in get Shopcarts!!!")
	}
	for _, val := range prod.Items {
		prodName, err := c.Strg.Product().GetById(&models.ProductPrimaryKey{val.ProductId})
		if err != nil {
			log.Println("Error while in getbyid product!!!", err)
		}
		if val.Status {
			totalCount[prodName.Name] += val.Count

		}
	}
	var active_products []int
	keys := make([]string, 0, len(totalCount))
	for key, val := range totalCount {
		keys = append(keys, key)
		active_products = append(active_products, val)
	}
	sort.Slice(active_products[:], func(i, j int) bool {
		return active_products[:][i] > active_products[:][j]
	})
	//sort.Ints(aktive_products)
	fmt.Println(active_products)

	return "", 0
}

//7
//passive
//8
//9
//10

func (c *Controller) ActiveClient() (string, error) {
	clients := make(map[string]int)
	items, err := c.ShopCartGetAll(&models.ShopCartGetListRequest{})
	if err != nil {
		return "", err
	}

	for _, val := range items.Items {
		if val.Status == true {
			itemsProduct, err := c.GetByIdPoduct(&models.ProductPrimaryKey{Id: val.ProductId})
			if err != nil {
				return " ", err
			}
			clients[val.UserId] += val.Count * itemsProduct.Price
		}
	}
	client, sum := "", 0
	for key, value := range clients {
		if sum < value {
			client = key
			sum = value
		}
	}

	itemsUser, err := c.UserGetById(&models.UserPrimaryKey{Id: client})
	if err != nil {
		return "", nil
	}
	name := itemsUser.Name + " " + itemsUser.Surname
	return name, nil
}
