package presenters

import (
	"net/http"
	"reflect"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/jackc/pgerrcode"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type userPresenter struct {
}

func NewUserPresenter() *userPresenter {
	return &userPresenter{}
}

func (p *userPresenter) Create(err error) (int, *output.CreateUserResponse) {
	if err != nil {
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(echo.ErrBadRequest):
			return 400, nil
		case reflect.TypeOf(&pq.Error{Code: pgerrcode.UniqueViolation}):
			return 409, nil
		default:
			return 500, nil
		}
	}

	return http.StatusOK, &output.CreateUserResponse{Message: "create successful"} // TODO:
}

func (p *userPresenter) Read(error) (int, *output.ReadUserResponse) {
	return 0, nil
}

func (p *userPresenter) Delete(err error) (int, *output.DeleteUserResponse) {
	if err != nil {
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(echo.ErrBadRequest):
			return 400, nil
		case reflect.TypeOf(&pq.Error{Code: pgerrcode.NoDataFound}):
			return 400, nil
		default:
			return 500, nil
		}
	}

	return http.StatusOK, &output.DeleteUserResponse{Message: "delete successful"} // TODO:
}
