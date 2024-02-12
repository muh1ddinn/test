package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Producct struct {
	Id                 int     `json:"id"`
	Title              string  `json:"title"`
	Description        string  `json:"description"`
	Price              float64 `json:"price"`
	DiscountPercentage float64 `json:"discountpercentage"`
	Rating             float64 `json:"rating"`
	Stock              float64 `json:"stock"`
	Brand              string  `json:"brand"`
	Category           string  `json:"category"`
	Thumbnail          string  `json:"thumbnail"`
}

func readfuncfromjson(filenaame string) []Producct {

	ddta, err := os.ReadFile(filenaame)
	if err != nil {

		fmt.Println("you have err while reading json file")
		return nil
	}

	var produ []Producct
	if err := json.Unmarshal(ddta, &produ); err != nil {
		fmt.Println("you have some issues while reading json file")

		return nil
	}

	return produ
}

func writeresultfile(filenaame, resultt string) {

	fil, err := os.Create(filenaame)
	if err != nil {

		fmt.Println("err with creating text file ", err)

	}
	defer fil.Close()

	_, err = fil.WriteString(resultt)
	if err != nil {

		fmt.Println("err with wrting to file", err)
	}

}

func main() {

	var max float64

	max = 0
	prod := readfuncfromjson("dataa.json")

	for i := 0; i < len(prod); i++ {

		if prod[i].DiscountPercentage > max {
			max = prod[i].DiscountPercentage
		}
	}
	fmt.Println("Maximum discount percentage:", max)

}
