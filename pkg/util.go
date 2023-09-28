package pkg

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
)

// Un metodo util para cargar datos de json en bbdd
func FullfilDB() ([]domain.Customers, error) {
	// Customers
	data, err := os.Open("./datos/customers.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	slice := []domain.Customers{}
	json.Unmarshal(dataRead, &slice)

	return slice, nil
}
