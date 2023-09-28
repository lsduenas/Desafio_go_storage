package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/products"
	"github.com/gin-gonic/gin"
)

type Products struct {
	s products.Service
}

func NewHandlerProducts(s products.Service) *Products {
	return &Products{s}
}

func (p *Products) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, products)
	}
}

func (p *Products) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products := domain.Product{}
		err := ctx.ShouldBindJSON(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = p.s.Create(&products)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": products})
	}
}

func (p *Products) ValidateDB() ([]*domain.Product, error) {
	products, err := p.s.ReadAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p *Products) InsertProducts() (bool, error) {
	// utils pkg
	data, err := os.Open("./datos/products.json")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	products := []domain.Product{}
	json.Unmarshal(dataRead, &products)

	if err != nil {
		err = errors.New("Can not get products list")
		return false, err
	}

	for _, product := range products {
		// service
		var pr domain.Product
		pr = product
		err = p.s.Create(&pr)

		if err != nil {
			err = errors.New("Can not insert product into DB")
			return false, err
		}
		//fmt.Println("customer", customer)
	}
	return true, nil
}
