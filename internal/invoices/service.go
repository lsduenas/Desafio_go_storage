package invoices

import "github.com/bootcamp-go/desafio-cierre-db.git/internal/domain"

type Service interface {
	Create(invoices *domain.Invoices) error
	ReadAll() ([]*domain.Invoices, error)
	UpdateTotalField() (error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Create(invoices *domain.Invoices) error {
	_, err := s.r.Create(invoices)
	if err != nil {
		return err
	}
	return nil

}

func (s *service) ReadAll() ([]*domain.Invoices, error) {
	return s.r.ReadAll()
}

func (s *service) UpdateTotalField() (error) {
	err := s.r.UpdateTotalField()
	if err != nil {
		return err
	}
	return nil
}