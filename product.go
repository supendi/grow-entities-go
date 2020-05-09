package entities

import "time"

//Domain and or response model
type (
	//Product TODO: Add price
	//Product represent the product model
	Product struct {
		AccountID  string     `json:"-"`
		ID         string     `json:"id"`
		CategoryID string     `json:"categoryId"`
		Name       string     `json:"name"`
		Price      float64    `json:"price"`
		PictureURL string     `json:"pictureUrl"`
		CreatedAt  time.Time  `json:"createdAt"`
		UpdatedAt  *time.Time `json:"updatedAt"`
		DeletedAt  *time.Time `json:"deletedAt"`
	}

	//ProductListFilter represent the filter model for querying product
	ProductListFilter struct {
		AccountID          string `json:"-"`
		Limit              int    `json:"limit"`
		PageIndex          int    `json:"pageIndex"`
		Keyword            string `json:"keyword"`
		OrderDir           string `json:"orderDir"`
		LastNameDisplayed  string `json:"lastNameDisplayed"`
		FirstNameDisplayed string `json:"firstNameDisplayed"`
	}

	//ProductList represent the response model of product
	ProductList struct {
		Records []*Product `json:"records"`
	}
)

//Request models
type (
	//ProductCreateRequest represent the create request model of product
	ProductCreateRequest struct {
		AccountID  string  `validate:"required"`
		CategoryID string  `validate:"required"`
		Name       string  `validate:"required"`
		Price      float64 `validate:"min=0,max=1000000000"`
		PictureURL string  ``
	}

	//ProductUpdateRequest represent the update request model of product
	ProductUpdateRequest struct {
		AccountID  string  `validate:"required"`
		ID         string  `validate:"required"`
		CategoryID string  `validate:"required"`
		Name       string  `validate:"required"`
		Price      float64 `validate:"min=0,max=1000000000"`
		PictureURL string  ``
	}

	//ProductDeleteRequest represent the delete request model of product
	ProductDeleteRequest struct {
		ProductID string `validate:"required"`
		AccountID string `validate:"required"`
	}

	//ProductGetRequest represent the get request model for getting single product
	ProductGetRequest struct {
		ProductID string `validate:"required"`
		AccountID string `validate:"required"`
	}

	Meki struct {
	}
)
