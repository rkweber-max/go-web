package models

import "github.com/rkweber/db"

type Products struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Products {
	db := db.ConnectionDB()

	getAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	p := Products{}
	products := []Products{}

	for getAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = getAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}
