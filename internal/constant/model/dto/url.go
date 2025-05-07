package dto

import (
	"fmt"
	"net/url"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
)

type URLRequest struct {
	OriginalURL string `json:"original_url"`
	ShortCode   string  `json:"short_code"`
}

func IsURL() validation.RuleFunc {
	return func(value interface{}) error {
		uri := fmt.Sprintf("%s", value)
		_, err := url.ParseRequestURI(uri)
		if err != nil {
			return fmt.Errorf("url must be a valid")
		}

		return nil
	}
}

func (u URLRequest) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.OriginalURL, validation.Required.Error("url is required"), validation.By(IsURL())),
	)
}

type URLResponse struct {
	ID          uuid.UUID `json:"id,omitempty,omitzero"`
	OriginalURL string    `json:"original_url,omitempty"`
	ShortURL    string    `json:"short_url,omitempty"`
	ShortCode   string    `json:"short_code,omitempty"`
	Count       int32     `json:"count,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitzero"`
	UpdatedAt   time.Time `json:"updated_at,omitzero"`
	DeletedAt   time.Time `json:"deleted_at,omitzero"`
}
