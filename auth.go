package entities

import "time"

type (
	//Token represent token model
	Token struct {
		ID             string
		AccessToken    string
		RefreshToken   string
		RequestedCount int
		Blacklisted    bool
		ExpiredAt      time.Time
		CreatedAt      time.Time
		UpdatedAt      *time.Time
	}

	//LoginRequest represent the model of login request
	LoginRequest struct {
		Username string `validate:"required"`
		Password string `validate:"required"`
	}

	//RenewTokenRequest represent the model while requesting a new access token
	RenewTokenRequest struct {
		AccessToken  string `validate:"required"`
		RefreshToken string `validate:"required"`
	}
)
