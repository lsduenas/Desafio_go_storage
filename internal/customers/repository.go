package customers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type Repository interface {
	Create(customers domain.Customers) (int64, error)
	ReadAll() ([]*domain.Customers, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(customers domain.Customers) (int64, error) {
	condition := customers.Condition
	var condition2 int
	if condition {
		condition2 = 1
	} else {
		condition2 = 0
	}
	query := `INSERT INTO customers (last_name, first_name, customers.condition) VALUES (?, ?, ?);`
	row, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println("olaaaaa",err)
		return 0, err
	}
	defer row.Close()
	result, err := row.Exec(&customers.LastName, &customers.FirstName, &condition2)
	drivererr, ok := err.(*mysql.MySQLError)
	if ok {
		//atrapamos los errores del driver
		log.Println("Error in create product: ", " ", drivererr.Number, drivererr.Message, drivererr.Error())
		err = errors.New("Internal")
		return 0, err
	}
	if err != nil {
		log.Println("Error CREATE execute Query", err.Error())
		err = errors.New("Internal")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Customers, error) {
	query := `SELECT id, first_name, last_name, customers.condition FROM customers`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := make([]*domain.Customers, 0)
	for rows.Next() {
		customer := domain.Customers{}
		err := rows.Scan(&customer.Id, &customer.FirstName, &customer.LastName, &customer.Condition)
		if err != nil {
			return nil, err
		}
		customers = append(customers, &customer)
	}
	return customers, nil
}
