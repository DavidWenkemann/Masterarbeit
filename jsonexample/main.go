package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type DBProduct struct {
	ProductID int     `sql:"id"`
	EAN       string  `sql:"ean"`
	Name      string  `sql:"name"`
	Price     float64 `sql:"price"`
}

type BProduct struct {
	ProductID int     `json:"product_id,omitempty"`
	EAN       string  `json:"ean,omitempty"`
	Name      string  `json:"name,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

type B2Product struct {
	ProductID int     `json:"product_id,omitempty"`
	EAN       string  `json:"ean,omitempty"`
	Name      string  `json:"name,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

func main() {

	bdp := BProduct{
		ProductID: 1,
		EAN:       "1234567890123",
		Name:      "foobar",
		Price:     12.3,
	}

	sbdp, err := json.Marshal(bdp)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(sbdp))
	queue := make(chan []byte, 1024)

	queue <- sbdp

	go func() {

		bdp2 := BProduct{
			ProductID: 2,
			EAN:       "1234567890124",
			Name:      "foobarfoo",
			Price:     12.4,
		}

		sbdp, err := json.Marshal(bdp2)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(sbdp))

		nt := time.NewTimer(time.Second)

		<-nt.C

		queue <- sbdp

	}()

	go func() {

		bdp3 := BProduct{
			ProductID: 3,
			EAN:       "1234567890125",
			Name:      "foobarbar",
			Price:     12.5,
		}

		sbdp, err := json.Marshal(bdp3)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(sbdp))

		nt := time.NewTimer(time.Second * 2)

		<-nt.C

		queue <- sbdp

	}()

	rp := B2Product{}
	quit := time.NewTimer(time.Second * 3)

	go func() {
		for {

			select {
			case msg := <-queue:
				err = json.Unmarshal(msg, &rp)
				if err != nil {
					panic(err)
				}
				//SAVE TO DB
				fmt.Println("(msg)json :", string(msg), "; reporting product go :", rp)
			case <-quit.C:
				break
			}
		}
	}()

}

func SoldItem(it )

Verkaufen -> Anpassung am Stock -> anpassung an allen andern stocks


func SendProduct(p BProduct) {

	mp, _ := json.Marshal(BProduct)

	chanBaseData <- mp
	chanReport <- mp
	chanStore <- mp

}
