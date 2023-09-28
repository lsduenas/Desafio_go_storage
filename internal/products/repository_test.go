package products

import (
	"database/sql"
	"testing"

	txdb "github.com/DATA-DOG/go-txdb"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init() {
	cfg := mysql.Config{
		User:                 "user1",
		Passwd:               "secret_password",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "fantasy_products",
		ParseTime:            true,
	}
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}
// Tests for StorageProductsMySQL.GetAll
func TestProductRepository_Create(t *testing.T) {
	t.Run("success - product ID", func(t *testing.T) {
		// arrange
		// -> database connection
		db, err := sql.Open("txdb", "TestProductRepository_Create_Product_ID")
		assert.NoError(t, err)
		defer db.Close()

		// -> storage
		st := NewRepository(db)
		// -> product
		var product *domain.Product
		product = &domain.Product{
			Description: "Papitas fritas :P",
			Price: 200.56,
		}

		// act
		p, err := st.Create(product)

		// assert
		var expectedProductID int64
		expectedProductID = 103
		expectedErr := error(nil)
		assert.Equal(t, expectedProductID, p)
		assert.Equal(t, expectedErr, err)
	})
}