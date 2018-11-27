package service

import (
	"errors"

	"git.zam.io/microservices/customer-api/models"
	"git.zam.io/microservices/customer-api/service/repository"
)

type ICustomerAPIService interface {
	Health() bool
	Create(c *models.Customer) (*models.Customer, error)
	LoadByID(id uint64) (*models.Customer, error)
	LoadByPhone(phone string) (*models.Customer, error)
}

var _ ICustomerAPIService = (*CustomerAPIService)(nil)

type CustomerAPIService struct {
	r repository.Repository
}

func (i *CustomerAPIService) Create(c *models.Customer) (*models.Customer, error) {
	panic("implement me")
}

func (i *CustomerAPIService) Health() bool {
	return true
}
func (i *CustomerAPIService) LoadByID(id uint64) (*models.Customer, error) {
	res := &models.Customer{}
	err := i.r.LoadByID(id, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (i *CustomerAPIService) LoadByPhone(phone string) (*models.Customer, error) {
	panic(errors.New("not implemented"))
}
