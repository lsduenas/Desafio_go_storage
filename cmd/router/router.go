package router

import (
	"database/sql"

	"github.com/bootcamp-go/desafio-cierre-db.git/cmd/handler"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/customers"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/invoices"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/products"
	"github.com/bootcamp-go/desafio-cierre-db.git/internal/sales"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r, r.Group("/api/v1"), db}
}

func (r *router) MapRoutes() {
	r.buildCustomersRoutes()
	r.buildInvoicesRoutes()
	r.buildProductsRoutes()
	r.buildSalesRoutes()
}

func (r *router) buildCustomersRoutes() {

	repo := customers.NewRepository(r.db)
	service := customers.NewService(repo)
	handler := handler.NewHandlerCustomers(service)

	// DB initialization - Customers table
	customerList, err := handler.ValidateDB()
	if err != nil {
		return
	}
	if len(customerList) == 0 {
		_, err := handler.InsertCustomers()
		if err != nil {
			return
		}
	}

	c := r.rg.Group("/customers")
	{
		c.GET("", handler.GetAll())
		c.POST("", handler.Post())
	}
}

func (r *router) buildInvoicesRoutes() {
	repo := invoices.NewRepository(r.db)
	service := invoices.NewService(repo)
	handler := handler.NewHandlerInvoices(service)

	// DB initialization - Invoices table
	invoiceList, err := handler.ValidateDB()
	if err != nil {
		return
	}
	if len(invoiceList) == 0 {
		_, err := handler.InsertInvoices()
		if err != nil {
			return
		}
	}

	i := r.rg.Group("/invoices")
	{
		i.GET("", handler.GetAll())
		i.POST("", handler.Post())
		i.PUT("/update_total", handler.UpdateTotalField())
	}
}

func (r *router) buildProductsRoutes() {
	repo := products.NewRepository(r.db)
	service := products.NewService(repo)
	handler := handler.NewHandlerProducts(service)
	// DB initialization - products table
	productList, err := handler.ValidateDB()
	if err != nil {
		return
	}
	if len(productList) == 0 {
		_, err := handler.InsertProducts()
		if err != nil {
			return
		}
	}
	p := r.rg.Group("/products")
	{
		p.GET("", handler.GetAll())
		p.POST("", handler.Post())
	}
}

func (r *router) buildSalesRoutes() {
	repo := sales.NewRepository(r.db)
	service := sales.NewService(repo)
	handler := handler.NewHandlerSales(service)

	// DB initialization - Sales table
	invoiceList, err := handler.ValidateDB()
	if err != nil {
		return
	}
	if len(invoiceList) == 0 {
		_, err := handler.InsertSales()
		if err != nil {
			return
		}
	}

	s := r.rg.Group("/sales")
	{
		s.GET("", handler.GetAll())
		s.POST("", handler.Post())

	}
}
