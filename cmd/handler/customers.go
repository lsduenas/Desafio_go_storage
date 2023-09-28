package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/customers"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"

	"github.com/gin-gonic/gin"
)

type Customers struct {
	s customers.Service
}

func NewHandlerCustomers(s customers.Service) *Customers {
	return &Customers{s}
}

func (c *Customers) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customers, err := c.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, customers)
	}
}

func (c *Customers) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer := domain.Customers{}
		err := ctx.ShouldBindJSON(&customer)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = c.s.Create(customer)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": customer})
	}
}

func (c *Customers) ValidateDB() ([]*domain.Customers, error) {
	customers, err := c.s.ReadAll()
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (c *Customers) InsertCustomers() (bool, error) {
	// utils pkg
	data, err := os.Open("./datos/customers.json")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	customers := []domain.Customers{}
	json.Unmarshal(dataRead, &customers)

	if err != nil {
		err = errors.New("Can not get customers list")
		return false, err
	}

	for _, customer := range customers {
		// service
		var cus domain.Customers
		cus = customer
		err = c.s.Create(cus)
		if err != nil {
			err = errors.New("Can not insert customer into DB")
			return false, err
		}
	}
	return true, nil
}
