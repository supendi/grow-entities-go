package entities

import "time"

//Domain and or response model
type (
	//Company represent the company model
	Company struct {
		AccountID string     `json:"-"`
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Phone     *string    `json:"phone"`
		Address   *Address   `json:"address"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}

	//CompanyListFilter represent the filter model for querying companies
	CompanyListFilter struct {
		AccountID string `json:"-"`
		Limit     int    `json:"limit"`
		PageIndex int    `json:"pageIndex"`
		Keyword   string `json:"keyword"`
	}

	//CompanyList represent the response model of companies
	CompanyList struct {
		Filter *CompanyListFilter `json:"filter"`
		Data   []*Company         `json:"data"`
	}
)

//Request models
type (
	//CompanyCreateRequest represent the create request model of company
	CompanyCreateRequest struct {
		AccountID string  `validate:"required"`
		Name      string  `validate:"required"`
		Phone     *string `json:"phone"`
		Address   *Address
	}

	//CompanyUpdateRequest represent the update request model of company
	CompanyUpdateRequest struct {
		ID        string `validate:"required"`
		AccountID string `validate:"required"`
		Name      string `validate:"required"`
		Phone     string `json:"phone"`
		Address   *Address
	}

	//CompanyDeleteRequest represent the delete request model of company
	CompanyDeleteRequest struct {
		CompanyID string `validate:"required"`
		AccountID string `validate:"required"`
	}

	//CompanyGetRequest represent the get request model for getting single company
	CompanyGetRequest struct {
		CompanyID string `validate:"required"`
		AccountID string `validate:"required"`
	}
)
