package service

import (
	"context"
	db_config "user-service/db"
	"user-service/models"

	"gorm.io/gorm"
)

type UserService struct{}

func (s *UserService) VerifyUserByToken(ctx context.Context, token *TokenRequest) (*BaseRespond, error) {
	var userRes = BaseRespond{
		Success: false,
	}

	var db *gorm.DB = db_config.SetupDatabaseConnectionGORM()
	defer db_config.CloseConnectionGORM(db)

	userRepo := models.NewUserRepo(db)

	_, err := userRepo.FindByToken(token.Token)
	if err != nil {
		return nil, err
	}

	userRes.Success = true
	return &userRes, nil

}
