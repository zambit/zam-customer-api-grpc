package endpoints

import "git.zam.io/microservices/customer-api/models"

type HealthRequest struct {
}

type HealthResponse struct {
	Ok bool
}

type LoadByIDRequest struct {
	Id uint64
}

type LoadByIDResponse struct {
	Customer models.Customer
}

type LoadByPhoneRequest struct {
	Phone string
}
type LoadByPhoneResponse struct {
	Customer models.Customer
}

type CreateRequest struct {
	Customer models.Customer
}

type CreateResponse struct {
	Customer models.Customer
	Err      error
}
