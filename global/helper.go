package global

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/nawafilhusnul/me-dashboard-api/src/domain"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Println(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func IsRequestValid(m *domain.Project) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
