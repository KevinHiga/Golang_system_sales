package usecase

import (
	"context"
	"fmt"
	"golang-project/config/dbiface"
	"golang-project/config/security"
	"golang-project/models"
	usersRepo "golang-project/repository/mongodb"
	"io"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func CreateUsersData(ctx context.Context, users []models.Users, collection dbiface.CollectionAPI) ([]interface{}, error) {
	return usersRepo.InsertUser(ctx, users, collection)
}

func FindUserAllData(ctx context.Context, collection dbiface.CollectionAPI) (*[]models.Users, error) {
	return usersRepo.FindUsers(ctx, collection)
}

func Login(ctx context.Context, inputs []models.Users, collection dbiface.CollectionAPI, c echo.Context) (*[]models.Users, error) {
	var user []models.Users
	for _, input := range inputs {
		uname, err := usersRepo.FindUserName(ctx, input.UserName, collection)
		if err != nil {
			err = fmt.Errorf("No se encuentra el usuario")
			return nil, err
		}
		err = security.VerifyPassword(uname.Password, input.Password)
		if err != nil {
			err = fmt.Errorf("Contraseña Equivocada")
			return nil, err
		}
		fmt.Println("paso")
		cookie := &http.Cookie{}
		claims := &models.Users{
			uname.ID,
			uname.UserName,
			uname.Mail,
			uname.Password,
			true,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(60 * time.Second).Unix(),
			},
		}
		fmt.Printf("linea 50 %v\n", claims)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return nil, err
		}
		cookie.Name = "sessionID"
		cookie.Value = t
		cookie.Expires = time.Now().Add(60 * time.Second)
		c.SetCookie(cookie)
		user = append(user, *uname)
	}
	return &user, nil
}

func ForgotPassword(ctx context.Context, inputs models.Users, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (*models.Users, error) {
	return usersRepo.ForgotPassword(ctx, inputs, reqBody, collection)
}

func Logout(c echo.Context) error {
	fmt.Println("paso")
	cookie := &http.Cookie{}
	cookie.Name = "sessionID"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Second)
	c.SetCookie(cookie)
	return nil
}

func User(ctx context.Context, collection dbiface.CollectionAPI, c echo.Context) (*models.Users, error) {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		return nil, err
	}
	tknStr := cookie.Value
	claims := &models.Users{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, err
	}
	uname, err := usersRepo.FindUserName(ctx, claims.UserName, collection)
	if err != nil {
		err = fmt.Errorf("No se encuentra el usuario")
		return nil, err
	}
	return uname, nil
}
