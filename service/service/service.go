package service

import (
	"errors"

	"git.zam.io/microservices/customer-api/models"
	"git.zam.io/microservices/customer-api/service"
)

var _ service.ICustomerAPIService = (*CustomerAPIService)(nil)

type CustomerAPIService struct {
}

func (i *CustomerAPIService) Create(c *models.Customer) (*models.Customer, error) {
	panic("implement me")
}

func (i CustomerAPIService) Health() bool {
	return true
}
func (i *CustomerAPIService) LoadByID(id uint64) *models.Customer {
	panic(errors.New("not implemented"))
}
func (i *CustomerAPIService) LoadByPhone(phone string) *models.Customer {
	panic(errors.New("not implemented"))
}
