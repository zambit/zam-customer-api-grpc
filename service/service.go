package service

import "git.zam.io/microservices/customer-api/models"

type ICustomerAPIService interface {
	Health() bool
	Create(c *models.Customer) (*models.Customer, error)
	LoadByID(id uint64) *models.Customer
	LoadByPhone(phone string) *models.Customer
}
