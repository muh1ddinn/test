package main

type Customer struct {
	name string
	cash int
	basket
}

type Basket struct {
}

type Store struct {
}

var (
	List = ProductList{
		{
			Name:          "bread",
			Count:         15,
			Price:         3000,
			OriginalPrice: 4000,
		},
		{
			Name:          "meat",
			Count:         5,
			Price:         120000,
			OriginalPrice: 90000,
		},
		{
			Name:          "milk",
			Count:         8,
			Price:         30000,
			OriginalPrice: 25000,
		},
		{
			Name:          "juice",
			Count:         20,
			Price:         13000,
			OriginalPrice: 11000,
		},
	}
)

type Product struct {
	Name          string
	Quantity      int
	Price         int
	OriginalPrice int
}
