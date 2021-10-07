package user

import (
	_handler "miniproject/app/presenter"
	"miniproject/business/user"
	"miniproject/app/presenter/user/request"
	"miniproject/app/presenter/user/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (handler *UserHandler) Register(ctx echo.Context) error {
	req := request.User{}
	if err := ctx.Bind(&req); err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusBadRequest, err)
	}

	data, err := handler.UserService.Register(req.ToDomain())
	if err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return _handler.NewSuccessResponse(ctx, response.FromDomainRegister(data))
}

func (handler *UserHandler) Login(ctx echo.Context) error {
	req := request.UserLogin{}

	if err := ctx.Bind(&req); err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusBadRequest, err)
	}
	data, err := handler.UserService.Login(req.Username, req.Password)
	if err != nil {
		return _handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
	}

	return _handler.NewSuccessResponse(ctx, response.FromDomainLogin(data))
}