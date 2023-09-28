package handler

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/sales"
	"github.com/gin-gonic/gin"
)

type Sales struct {
	s sales.Service
}

func NewHandlerSales(s sales.Service) *Sales {
	return &Sales{s}
}

func (s *Sales) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		invoices, err := s.s.ReadAll()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, invoices)
	}
}

func (s *Sales) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sale := domain.Sales{}
		err := ctx.ShouldBindJSON(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		err = s.s.Create(&sale)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(201, gin.H{"data": sale})
	}
}

func (s *Sales) ValidateDB() ([]*domain.Sales, error) {
	sales, err := s.s.ReadAll()
	if err != nil {
		return nil, err
	}
	return sales, nil
}

func (s *Sales) InsertSales() (bool, error) {
	// utils pkg
	data, err := os.Open("./datos/sales.json")
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	sales := []domain.Sales{}
	json.Unmarshal(dataRead, &sales)

	if err != nil {
		err = errors.New("Can not get sales list")
		return false, err
	}

	for _, sale := range sales {
		// service
		var sal domain.Sales
		sal = sale
		err = s.s.Create(&sal)

		if err != nil {
			err = errors.New("Can not insert sale into DB")
			return false, err
		}
	}
	return true, nil
}
