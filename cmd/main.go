package main

import (
	"app/config"
	"app/controller"
	"app/models"
	jsondb "app/storage/jsonDB"
	"fmt"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewconnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)
	//User
	// CreateNewUser(con)
	//GetUserByID(con)
	// GetAll_Users(con)
	// Update_User(con)
	//Delete_User(con)
	//Product
	// CreateNewProduct(con)
	// GetProductByID(con)
	// Update_Product(con)
	// Delete_Product(con)
	// GetAll_Products(con)
	// //Category
	// CreateNewCategory(con)
	// GetCategoryByID(con)
	// GetAll_Categoties(con)
	// Update_Category(con)
	// Delete_Category(con)
	// //ShopCart
	// CreateNewShopCart(con)
	// GetShopCartByID(con)
	// GetAll_ShopCarts(con)
	// Update_ShopCart(con)
	// Delete_ShopCard(con)
	UserHistory(con)
	//UserMoney(con)

}

//User
func CreateNewUser(con *controller.Controller) {
	con.UserCreate(&models.CreateUser{
		Name:    "Abdu",
		Surname: "Arzubaev",
		Balance: 2_000_000,
	})
}
func GetUserByID(con *controller.Controller) {
	resp, _ := con.UserGetById(&models.UserPrimaryKey{
		Id: "e6ded598-675b-4de2-a1e9-00a876b8e719",
	})
	fmt.Println(resp)
}
func GetAll_Users(con *controller.Controller) {
	users, _ := con.UserGetAll(&models.UserGetListRequest{
		Offset: 0,
		Limit:  10,
	})
	for _, val := range users.Users {
		fmt.Println(val)
	}
}
func Update_User(con *controller.Controller) {
	con.UserUpdate(&models.UpdateUser{
		Id:      "05622de4-5be4-4254-8449-bcd3dd557631",
		Name:    "Asadbek",
		Surname: "Ergashev",
		Balance: 1_600_000,
	})
}
func Delete_User(con *controller.Controller) {
	con.UserDelete(&models.UserPrimaryKey{
		Id: "8b2ffa05-3b57-46b5-b00a-146628c721e6",
	})
}

//Product
func CreateNewProduct(con *controller.Controller) {
	con.ProductCreate(&models.CreateProduct{
		Name:       "Samsung Galaxy Book",
		Price:      15_000_000,
		CategoryID: "f4b72191-efb2-4b1a-ad78-5550aff87c34",
	})
}
func GetProductByID(con *controller.Controller) {
	resp, _ := con.GetByIdPoduct(&models.ProductPrimaryKey{
		Id: "a8d7239f-8348-4ffd-a53b-f6fea535ad56",
	})
	fmt.Println(resp)
}
func GetAll_Products(con *controller.Controller) {
	products, _ := con.ProductGetAll(&models.ProductGetListRequest{
		Offset: 0,
		Limit:  10,
	})
	for _, val := range products.Products {
		fmt.Println(val)
	}
}
func Update_Product(con *controller.Controller) {
	con.ProductUpdate(&models.UpdateProduct{
		Id:         "a8d7239f-8348-4ffd-a53b-f6fea535ad56",
		Name:       "xazinan",
		Price:      10_000,
		CategoryID: "bfca8c38-e85a-42c5-89af-8777c448e711",
	})
}
func Delete_Product(con *controller.Controller) {
	con.ProductDelete(&models.ProductPrimaryKey{
		Id: "2848cbc4-e5b2-4dd5-9b44-28adef456020",
	})
}

//category
func CreateNewCategory(con *controller.Controller) {
	con.CategoryCreate(&models.CreateCategory{
		Name:     "shimlar",
		ParentId: "2997b6c1-ea68-4237-9033-e17fba9af821",
	})
}
func GetCategoryByID(con *controller.Controller) {
	resp, _ := con.CategoryGetById(&models.CategoryPrimaryKey{
		Id: "2997b6c1-ea68-4237-9033-e17fba9af821",
	})
	fmt.Println(resp)
}
func GetAll_Categoties(con *controller.Controller) {
	categories, _ := con.CategoryGetAll(&models.CategoryGetListRequest{
		Offset: 0,
		Limit:  10,
	})
	for _, val := range categories.Categories {
		fmt.Println(val)
	}
}
func Update_Category(con *controller.Controller) {
	con.CategoryUpdate(&models.UpdateCategory{
		Id:       "152e6091-80cc-4fe1-b592-74c89bf6b480",
		Name:     "printer",
		ParentId: "117bc391-ce09-4976-b5e9-7fdde869895b",
	})
}
func Delete_Category(con *controller.Controller) {
	con.CategoryDelete(&models.CategoryPrimaryKey{
		Id: "152e6091-80cc-4fe1-b592-74c89bf6b480",
	})
}

//ShopCart
func CreateNewShopCart(con *controller.Controller) {
	con.ShopCartCreate(&models.CreateShopCart{
		ProductId: "9ca18640-35f9-46ae-8c02-e8ead811f0db",
		UserId:    "05622de4-5be4-4254-8449-bcd3dd557631",
		Count:     2,
		Status:    true,
	})
}
func GetShopCartByID(con *controller.Controller) {
	resp, _ := con.ShopCartGetById(&models.ShopCartprimarykey{
		ProductId: "a80cc924-fec3-4717-8289-f23604de45ae",
		UserId:    "ddc46ae9-6ccc-450a-ad74-50276f3c09f1",
	})
	fmt.Println(resp)
}
func GetAll_ShopCarts(con *controller.Controller) {
	shopCarts, _ := con.ShopCartGetAll(&models.ShopCartGetListRequest{
		Offset: 0,
		Limit:  10,
	})
	for _, val := range shopCarts.Items {
		fmt.Println(val)
	}
}
func Update_ShopCart(con *controller.Controller) {
	con.ShopCartUpdate(&models.UpdateShopCart{
		ProductId: "ffa888f7-e0cb-44e9-9cae-8c4ca2a115b9",
		UserId:    "e6ded598-675b-4de2-a1e9-00a876b8e719",
		Count:     5,
		Status:    true,
	})
}
func Delete_ShopCard(con *controller.Controller) {
	con.Delete_ShopCard(&models.ShopCartprimarykey{
		ProductId: "a80cc924-fec3-4717-8289-f23604de45ae",
		UserId:    "ddc46ae9-6ccc-450a-ad74-50276f3c09f1",
	})
}

//task-1
func UserHistory(con *controller.Controller) {
	con.GetUserProducts(&models.UserPrimaryKey{
		Id: "05622de4-5be4-4254-8449-bcd3dd557631",
	})
}

//task-4  ???
func UserMoney(con *controller.Controller) {
	con.UserMoneySpent(&models.UserPrimaryKey{
		Id: "05622de4-5be4-4254-8449-bcd3dd557631",
	})
}
