package usecase

import (
	"context"

	"github.com/tfkhdyt/SpaceNotes/server/internal/application/dto"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/entity"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/repository"
	"github.com/tfkhdyt/SpaceNotes/server/internal/domain/service"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/exception"
	"github.com/tfkhdyt/SpaceNotes/server/pkg/sql"
)

type AuthUsecase struct {
	refreshTokenRepo repository.RefreshTokenRepo `di.inject:"refreshTokenRepo"`
	userRepo         repository.UserRepo         `di.inject:"userRepo"`
	hashingService   service.HashingService      `di.inject:"hashingService"`
	authTokenService service.AuthTokenService    `di.inject:"authTokenService"`
}

func (a *AuthUsecase) Register(
	newUser *dto.RegisterRequest,
) (*dto.RegisterResponse, error) {
	ctx := context.Background()

	if _, err := a.userRepo.FindUserByUsername(
		ctx,
		newUser.Username,
	); err == nil {
		return nil, exception.NewHTTPError(400, "username has been used")
	}

	if _, err := a.userRepo.FindUserByEmail(ctx, newUser.Email); err == nil {
		return nil, exception.NewHTTPError(400, "email has been used")
	}

	var errHash error
	newUser.Password, errHash = a.hashingService.HashPassword(newUser.Password)
	if errHash != nil {
		return nil, errHash
	}

	registeredUser, errRegister := a.userRepo.CreateUser(ctx, &entity.NewUser{
		FullName: sql.NewNullString(newUser.FullName),
		Username: newUser.Username,
		Email:    newUser.Email,
		Password: newUser.Password,
	})
	if errRegister != nil {
		return nil, errRegister
	}

	response := &dto.RegisterResponse{
		Message: "Your account has been created successfully",
		Data: dto.RegisterResponseData{
			ID:        registeredUser.ID,
			FullName:  registeredUser.FullName.String,
			Username:  registeredUser.Username,
			Email:     registeredUser.Email,
			CreatedAt: registeredUser.CreatedAt,
		},
	}

	return response, nil
}

func (a *AuthUsecase) Login(data *dto.LoginRequest) (*dto.LoginResponse, error) {
	ctx := context.Background()

	user, err := a.userRepo.FindUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if err := a.hashingService.ComparePassword(user.Password, data.Password); err != nil {
		return nil, err
	}

	accessToken, errAccess := a.authTokenService.CreateAccessToken(user.ID)
	if errAccess != nil {
		return nil, errAccess
	}

	refreshToken, errRefresh := a.authTokenService.CreateRefreshToken(user.ID)
	if errRefresh != nil {
		return nil, errRefresh
	}

	response := &dto.LoginResponse{
		Message: "You've logged in successfully",
		Data: dto.LoginResponseData{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}

	return response, nil
}
