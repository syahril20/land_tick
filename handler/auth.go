package handler

import (
	authdto "landtick/dto/auth"
	dto "landtick/dto/result"
	userdto "landtick/dto/user"
	"log"
	"net/http"
	"time"

	"landtick/models"
	"landtick/pkg/bcrypt"
	jwtToken "landtick/pkg/jwt"
	"landtick/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type handlerAuth struct {
	AuthRepository repositories.AuthRepository
}

func HandlerAuth(AuthRepository repositories.AuthRepository) *handlerAuth {
	return &handlerAuth{AuthRepository}
}

func (h *handlerAuth) Register(c echo.Context) error {
	request := new(userdto.CreateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	Password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		NamaLengkap:  request.NamaLengkap,
		Username:     request.Username,
		Email:        request.Email,
		Password:     Password,
		JenisKelamin: request.JenisKelamin,
		Telp:         request.Telp,
		Alamat:       request.Alamat,
		IdRole:       request.IdRole,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 2 hours expired
	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	RegResponse := authdto.RegResponse{
		Email: user.Email,
		Token: token,
	}

	_, err = h.AuthRepository.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: RegResponse})
}

func (h *handlerAuth) Login(c echo.Context) error {
	request := new(authdto.LoginRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	user := models.User{
		Email:    request.Email,
		Password: request.Password,
	}

	// Check email
	user, err := h.AuthRepository.Login(user.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	// Check password
	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"})
	}

	//generate token
	claims := jwt.MapClaims{}
	claims["id"] = user.Id
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 hours expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	loginResponse := authdto.LoginResponse{
		Email:  user.Email,
		Token:  token,
		RoleId: user.IdRole,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginResponse})
}
