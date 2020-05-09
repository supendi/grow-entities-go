package main

import "time"

type (
	//Account represent account entity model
	Account struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		Phone     string     `json:"phone"`
		Password  string     `json:"-"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}

	//Registrant represent a registrant data model who wants to register as a new account
	Registrant struct {
		Name     string `validate:"required"`
		Email    string `validate:"required"`
		Phone    string `validate:"required"`
		Password string `validate:"required"`
	}

	//UpdateRequest represent account update request model
	UpdateRequest struct {
		AccountID string `validate:"required"`
		Name      string `validate:"required"`
	}

	//GetRequest represent account get request model
	GetRequest struct {
		AccountID string `validate:"required"`
	}
)
