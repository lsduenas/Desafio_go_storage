package invoices

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type Repository interface {
	Create(invoices *domain.Invoices) (int64, error)
	ReadAll() ([]*domain.Invoices, error)
	UpdateTotalField() error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(invoices *domain.Invoices) (int64, error) {
	query := `INSERT INTO invoices (customer_id, datetime, total) VALUES (?, ?, ?)`
	row, err := r.db.Exec(query, &invoices.CustomerId, &invoices.Datetime, &invoices.Total)
	if err != nil {
		return 0, err
	}
	id, err := row.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadAll() ([]*domain.Invoices, error) {
	query := `SELECT id, customer_id, datetime, total FROM invoices`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	invoices := make([]*domain.Invoices, 0)
	for rows.Next() {
		invoice := domain.Invoices{}
		err := rows.Scan(&invoice.Id, &invoice.CustomerId, &invoice.Datetime, &invoice.Total)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, &invoice)
	}
	return invoices, nil
}

func (r *repository) UpdateTotalField() error {
	stmt, err := r.db.Prepare(`
	UPDATE fantasy_products.invoices AS f1
	JOIN (
 	SELECT TRUNCATE(SUM(p.price*s.quantity),2) as nuevo_total, s.invoice_id as invoice FROM fantasy_products.sales s INNER JOIN
	fantasy_products.products p ON s.product_id=p.id
	GROUP BY invoice
	) AS f2 ON f1.id = f2.invoice
	SET f1.total = f2.nuevo_total`)

	if err != nil {
		log.Println("error ", err.Error())
	}
	fmt.Println("stament", stmt)
	defer stmt.Close()
	_, err = stmt.Query()
	fmt.Println(err)
	driverErr, ok := err.(*mysql.MySQLError)
	if ok {

		log.Println("Error in create product: ", driverErr.Number, driverErr.Message, driverErr.Error())
		err = errors.New("Internal")
		return err
	}
	if err != nil {
		log.Println("Error CREATE execute Query", err.Error())
		err = errors.New("Internal")
		return err
	}

	return err
}
