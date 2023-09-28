package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/invoices"
	"github.com/gin-gonic/gin"
)

type Invoices struct {
	s invoices.Service
}

func NewHandlerInvoices(s invoices.Service) *Invoices {
	return &Invoices{s}
}

func (i *Invoices) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices, err := i.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, invoices)
	}
}

func (i *Invoices) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices := domain.Invoices{}
		err := ctx.ShouldBindJSON(&invoices)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = i.s.Create(&invoices)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": invoices})
	}
}

func (i *Invoices) UpdateTotalField() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	
		err := i.s.UpdateTotalField()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("err", err)
		ctx.JSON(201, gin.H{"data": "Total field of invoice was updated"})
	}
}

func (i *Invoices) ValidateDB() ([]*domain.Invoices, error) {
	invoices, err := i.s.ReadAll()
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (i *Invoices) InsertInvoices() (bool, error) {
	// utils pkg
	data, err := os.Open("./datos/invoices.json")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	invoices := []domain.Invoices{}
	json.Unmarshal(dataRead, &invoices)

	if err != nil {
		err = errors.New("Can not get invoices list")
		return false, err
	}

	for _, invoice := range invoices {
		// service
		var inv domain.Invoices
		inv = invoice
		err = i.s.Create(&inv)

		if err != nil {
			err = errors.New("Can not insert invoice into DB")
			return false, err
		}
	}
	return true, nil
}
