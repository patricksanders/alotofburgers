package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

type Burger struct {
	Name          string
	Restaurant    string
	Price         float32
	TwitterHandle string
}

func readBurgers(filename string) []Burger {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var burgers []Burger

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		burgers = append(burgers, Burger{
			Name:          line[0],
			Restaurant:    line[1],
			Price:         parsePrice(line[2]),
			TwitterHandle: line[3],
		})
	}

	return burgers
}

func parsePrice(priceString string) float32 {
	price, err := strconv.ParseFloat(priceString, 32)
	price = math.Round(price*100) / 100
	if err != nil {
		log.Fatal(err)
	}

	return float32(price)
}

func main() {
	burgers := readBurgers("burgers.csv")
	fmt.Println(burgers)
}
