package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/blackironj/bookchive-server/da"
	"github.com/blackironj/bookchive-server/env"
	"github.com/blackironj/bookchive-server/middleware/jwt"
	"github.com/blackironj/bookchive-server/model"
)

//Signin with google
func Signin(signinData *model.Users) error {
	where := "WHERE email = ?"
	condVal := []interface{}{signinData.Email}

	user, _ := da.GetUsers(da.DB, where, condVal)
	if user != nil {
		txErr := da.DoInTransaction(func(tx *sqlx.Tx) error {
			setStmt := "signin_dt = ? WHERE uuid = ?"
			val := []interface{}{time.Now().Unix(), user[0].UUID}

			if err := da.UpdateUser(tx, setStmt, val); err != nil {
				return err
			}

			signinData.UUID = user[0].UUID
			return nil
		})

		if txErr != nil {
			return txErr
		}
		return nil
	}

	txErr := da.DoInTransaction(func(tx *sqlx.Tx) error {
		u, err := uuid.NewRandom()
		if err != nil {
			return err
		}

		newUser := &model.Users{
			UUID:  u.String(),
			Email: signinData.Email,
			Name:  signinData.Name,
		}
		if err := da.InsertUser(tx, newUser); err != nil {
			return err
		}

		signinData.UUID = u.String()
		return nil
	})

	if txErr != nil {
		return txErr
	}
	return nil
}

func GenJWT(user *model.Users) (string, error) {
	claims := &jwt.Claims{
		UUID:  user.UUID,
		Email: user.Email,
	}

	tokenExpireFrom := time.Hour * time.Duration(env.Conf.Auth.TokenExpireTimeHour)
	token, tokErr := jwt.GenerateJWT(claims, env.Conf.Auth.JWTKey, env.Conf.Auth.Issuer, tokenExpireFrom)
	if tokErr != nil {
		return "", errors.New("fail to generate a jwt")
	}
	return token, nil
}
