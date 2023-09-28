package customers

import "github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"

type Service interface {
	Create(customers domain.Customers) error
	ReadAll() ([]*domain.Customers, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(customers domain.Customers) (err error) {
	_, err = s.r.Create(customers)
	if err != nil {
		return 
	}
	return 
}

func (s *service) ReadAll() ([]*domain.Customers, error) {
	return s.r.ReadAll()
}
