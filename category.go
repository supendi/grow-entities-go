package entities

import "time"

//Domain and or response model
type (
	//Category represent the product category model
	Category struct {
		AccountID string     `json:"-"`
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt"`
		DeletedAt *time.Time `json:"deletedAt"`
	}

	//CategoryListFilter represent the filter model for querying product categories
	CategoryListFilter struct {
		AccountID          string `json:"-"`
		Limit              int    `json:"limit"`
		PageIndex          int    `json:"pageIndex"`
		Keyword            string `json:"keyword"`
		OrderDir           string `json:"orderDir"`
		LastNameDisplayed  string `json:"lastNameDisplayed"`
		FirstNameDisplayed string `json:"firstNameDisplayed"`
	}

	//CategoryList represent the response model of product categories
	CategoryList struct {
		Records []*Category `json:"records"`
	}
)

//Request models
type (
	//CategoryCreateRequest represent the create request model of product category
	CategoryCreateRequest struct {
		AccountID string `validate:"required"`
		Name      string `validate:"required"`
	}

	//CategoryUpdateRequest represent the update request model of product category
	CategoryUpdateRequest struct {
		ID        string `validate:"required"`
		AccountID string `validate:"required"`
		Name      string `validate:"required"`
	}

	//CategoryDeleteRequest represent the delete request model of product category
	CategoryDeleteRequest struct {
		CategoryID string `validate:"required"`
		AccountID  string `validate:"required"`
	}

	//CategoryGetRequest represent the get request model for getting single product category
	CategoryGetRequest struct {
		CategoryID string `validate:"required"`
		AccountID  string `validate:"required"`
	}
)
