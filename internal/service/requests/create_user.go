package requests

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cifra-city/users-storage/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func newDecodeError(what string, err error) error {
	return validation.Errors{
		what: fmt.Errorf("decode request %s: %w", what, err),
	}
}

func NewCreateUse(r *http.Request) (req resources.UserCreate, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		err = newDecodeError("body", err)
		return
	}

	errs := validation.Errors{
		"data/type":       validation.Validate(req.Data.Type, validation.Required, validation.In("create_user")),
		"data/attributes": validation.Validate(req.Data.Attributes, validation.Required),
	}
	return req, errs.Filter()
}
