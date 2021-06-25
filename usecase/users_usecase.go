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
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Login(ctx context.Context, users []models.Users, reqBody io.ReadCloser, collection dbiface.CollectionAPI, c echo.Context) (*[]models.Users, error) {
	var user []models.Users
	var session models.Session
	for _, input := range users {
		uname, err := usersRepo.FindUserName(context.Background(), input.UserName, collection)
		if err != nil {
			err = fmt.Errorf("User not found, try again")
			return nil, err
		}
		err = security.VerifyPassword(uname.Password, input.Password)
		if err != nil {
			err = fmt.Errorf("Wrong password, try again")
			return nil, err
		}
		cookie := &http.Cookie{}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    uname.UserName,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(), //1 day
		})
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return nil, err
		}
		session = input.Sesion
		session.Ssid = t
		session.Username = input.UserName
		session.Enabled = true
		fmt.Printf("linea 42 user %v\n\n", session)
		sessionsCollection := usersRepo.GetSessionsCollection()
		err = sessionsCollection.AddSession(ctx, &session)
		if err != nil {
			err = fmt.Errorf("Internal Database error.")
			return nil, err
		}
		cookie.Name = "sessionID"
		cookie.Value = t
		cookie.Expires = time.Now().Add(60 * time.Minute)
		c.SetCookie(cookie)
		user = append(user, *uname)
	}
	return &user, nil
}

func LoginGoogle(ctx context.Context, users []models.Users, reqBody io.ReadCloser, collection dbiface.CollectionAPI, c echo.Context) (*[]models.Users, error) {
	var user []models.Users
	for _, input := range users {
		mail, err := usersRepo.FindMail(context.Background(), input.Mail, collection)
		if err != nil {
			IDs, err := usersRepo.InsertUser(context.Background(), users, collection)
			if err != nil {
				return nil, err
			}
			str := fmt.Sprintf("%v", IDs)
			str2 := strings.TrimLeft(str, "[")
			str2 = strings.TrimRight(str2, "]")
			u, err2 := usersRepo.FindUser(context.Background(), str2, collection)
			if err2 != nil {
				return nil, err2
			}
			cookie := &http.Cookie{}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
				Issuer:    u.UserName,
				ExpiresAt: time.Now().Add(60 * time.Minute).Unix(), //1 day
			})
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return nil, err
			}
			cookie.Name = "sessionID"
			cookie.Value = t
			cookie.Expires = time.Now().Add(60 * time.Minute)
			c.SetCookie(cookie)
			user = append(user, u)
		}
		if mail != nil {
			err = security.VerifyPassword(mail.Password, input.Password)
			if err != nil {
				err = fmt.Errorf("Contraseña Equivocada")
				return nil, err
			}
			cookie := &http.Cookie{}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
				Issuer:    mail.UserName,
				ExpiresAt: time.Now().Add(60 * time.Minute).Unix(), //1 day
			})
			t, err := token.SignedString([]byte("secret"))
			if err != nil {
				return nil, err
			}
			cookie.Name = "sessionID"
			cookie.Value = t
			cookie.Expires = time.Now().Add(60 * time.Minute)
			c.SetCookie(cookie)
			user = append(user, *mail)
		}
	}
	return &user, nil
}

func CreateUsersData(ctx context.Context, users []models.Users, collection dbiface.CollectionAPI) ([]interface{}, error) {
	return usersRepo.InsertUser(ctx, users, collection)
}

func FindUserAllData(ctx context.Context, collection dbiface.CollectionAPI) (*[]models.Users, error) {
	return usersRepo.FindUsers(ctx, collection)
}

func FindUserOneData(ctx context.Context, id string, collection dbiface.CollectionAPI) (models.Users, error) {
	return usersRepo.FindUser(ctx, id, collection)
}

func ForgotPassword(ctx context.Context, inputs []models.Users, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (*[]models.Users, error) {
	return usersRepo.ForgotPassword(ctx, inputs, reqBody, collection)
}

func Logout(c echo.Context) error {
	fmt.Println("paso\n")
	cookie := new(http.Cookie)
	//cookie := &http.Cookie{}
	fmt.Printf("paso %v\n", cookie)
	cookie.Name = "sessionID"
	cookie.Value = ""
	//cookie.Expires = time.Now().Add(-time.Duration(time.Now().Year()))
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "")
}

func User(ctx context.Context, collection dbiface.CollectionAPI, c echo.Context) (*models.Users, error) {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		return nil, err
	}
	tkn, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	claims := tkn.Claims.(*jwt.StandardClaims)
	uname, err := usersRepo.FindUserName(ctx, claims.Issuer, collection)
	if err != nil {
		err = fmt.Errorf("No se encuentra el usuario")
		return nil, err
	}
	uname.Sesion.Ssid = cookie.Value
	return uname, nil
}
