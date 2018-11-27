package repository

import (
	"git.zam.io/microservices/customer-api/models"
	"git.zam.io/microservices/customer-api/pkg/db"
)

type IRepository interface {
	Create(customer *models.Customer) (uint64, error)
	LoadByID(id uint64, buf *models.Customer) error
	LoadByPhone(email string, buf *models.Customer) error
}

var _ IRepository = (*Repository)(nil)

type Repository struct{}

func (Repository) Create(customer *models.Customer) (uint64, error) {
	return 0, nil
}

func (Repository) LoadByID(id uint64, buf *models.Customer) error {
	buf.ID = id
	err := db.DB().Model(&buf).Select()
	if err != nil {
		return err
	}
	return nil
}

func (Repository) LoadByPhone(phone string, buf *models.Customer) error {
	buf.Phone = phone
	err := db.DB().Model(&buf).Select()
	if err != nil {
		return err
	}
	return nil
}
