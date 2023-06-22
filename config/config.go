package config

type Config struct {
	Path             string
	UserFileName     string
	ProductFileName  string
	CategoryFileName string
	ShopCartFileName string
	DefaultOffset    int
	DefaultLimit     int
}

func Load() Config {

	cfg := Config{}
	cfg.DefaultOffset = 0
	cfg.DefaultLimit = 10
	cfg.Path = "./data"
	cfg.UserFileName = "/user.json"
	cfg.ProductFileName = "/product.json"
	cfg.CategoryFileName = "/category.json"
	cfg.ShopCartFileName = "/shop_cart.json"
	return cfg
}
