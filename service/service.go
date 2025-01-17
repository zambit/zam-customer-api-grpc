package service

import (
	"context"
	"time"

	"git.zam.io/microservices/customer-api/models"
	"golang.org/x/crypto/bcrypt"
)

type ICustomerAPIService interface {
	Health(ctx context.Context) bool
	Create(ctx context.Context, req *CreateRequest, c *models.Customer) error
	LoadByID(ctx context.Context, id uint64) (*models.Customer, error)
	LoadByPhone(ctx context.Context, phone string) (*models.Customer, error)
	Login(ctx context.Context, phone string, password string) (*models.Customer, error)
}

var _ ICustomerAPIService = (*CustomerAPIService)(nil)

type CustomerAPIService struct {
	r Repository
}

func (i *CustomerAPIService) Login(ctx context.Context, phone string, password string) (*models.Customer, error) {
	// obtain customer data
	c := models.Customer{}
	err := i.r.LoadByPhone(phone, &c)
	if err != nil {
		return nil, err
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (i *CustomerAPIService) Create(ctx context.Context, req *CreateRequest, c *models.Customer) error {
	{
		t := time.Now()
		c.CreatedAt = t
		c.RegisteredAt = t
	}
	// hash password
	{
		bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		c.Password = string(bytes)
	}
	{
		c.Phone = req.Phone
		c.ReferrerID = req.ReferrerID
		c.StatusID = req.StatusID
	}
	_, err := i.r.Create(c)
	if err != nil {
		return err
	}
	return nil
}

func (i *CustomerAPIService) Health(ctx context.Context) bool {
	return true
}

func (i *CustomerAPIService) LoadByID(ctx context.Context, id uint64) (*models.Customer, error) {
	res := &models.Customer{}
	err := i.r.LoadByID(id, res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (i *CustomerAPIService) LoadByPhone(ctx context.Context, phone string) (*models.Customer, error) {
	res := &models.Customer{}
	err := i.r.LoadByPhone(phone, res)
	if err != nil {
		return res, err
	}
	return res, nil
}
