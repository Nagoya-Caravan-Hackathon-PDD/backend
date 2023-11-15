package presenters

import (
	"database/sql"
	"net/http"
	"reflect"

	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/output"
	"github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/datastructure/types"
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

func (p *userPresenter) Read(user *types.User, err error) (int, *output.ReadUserResponse) {
	if err != nil {
		switch reflect.TypeOf(err) {
		case reflect.TypeOf(echo.ErrBadRequest):
			return 400, nil
		case reflect.TypeOf(sql.ErrNoRows):
			return 404, nil
		case reflect.TypeOf(&pq.Error{Code: pgerrcode.NoDataFound}):
			return 400, nil
		default:
			return 500, nil
		}
	}
	return http.StatusOK, &output.ReadUserResponse{UserID: user.UserID, GitHubID: user.GitHubID} // TODO:
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
