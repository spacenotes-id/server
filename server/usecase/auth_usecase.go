package usecase

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/tfkhdyt/SpaceNotes/server/database/postgres/sqlc"
	"github.com/tfkhdyt/SpaceNotes/server/dto"
	"github.com/tfkhdyt/SpaceNotes/server/helper/sql"
	"github.com/tfkhdyt/SpaceNotes/server/repository/postgres"
	"github.com/tfkhdyt/SpaceNotes/server/service"
)

type AuthUsecase struct {
	refreshTokenRepo *postgres.RefreshTokenRepoPostgres `di.inject:"refreshTokenRepo"`
	userRepo         *postgres.UserRepoPostgres         `di.inject:"userRepo"`
	bcryptService    *service.BcryptService             `di.inject:"bcryptService"`
	jwtService       *service.JwtService                `di.inject:"jwtService"`
}

func (a *AuthUsecase) verifyUsernameAvailability(
	ctx context.Context,
	username string,
) error {
	if _, err := a.userRepo.FindUserByUsername(ctx, username); err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Username has been used")
	}

	return nil
}

func (a *AuthUsecase) verifyEmailAvailability(
	ctx context.Context,
	email string,
) error {
	if _, err := a.userRepo.FindUserByEmail(ctx, email); err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Email has been used")
	}

	return nil
}

func (a *AuthUsecase) Register(
	newUser *dto.RegisterRequest,
) (*dto.RegisterResponse, error) {
	ctx := context.Background()

	if err := a.verifyUsernameAvailability(ctx, newUser.Username); err != nil {
		return nil, err
	}

	if err := a.verifyEmailAvailability(ctx, newUser.Email); err != nil {
		return nil, err
	}

	var errHash error
	newUser.Password, errHash = a.bcryptService.HashPassword(newUser.Password)
	if errHash != nil {
		return nil, errHash
	}

	registeredUser, errRegister := a.userRepo.CreateUser(ctx,
		sqlc.CreateUserParams{
			FullName: pgtype.Text(sql.NewNullString(newUser.FullName)),
			Username: newUser.Username,
			Email:    newUser.Email,
			Password: newUser.Password,
		},
	)
	if errRegister != nil {
		return nil, errRegister
	}

	response := &dto.RegisterResponse{
		Message: "Your account has been created successfully",
		Data: dto.RegisterResponseData{
			ID:        int(registeredUser.ID),
			FullName:  registeredUser.FullName.String,
			Username:  registeredUser.Username,
			Email:     registeredUser.Email,
			CreatedAt: registeredUser.CreatedAt.Time,
		},
	}

	return response, nil
}

func (a *AuthUsecase) Login(
	data *dto.LoginRequest,
) (*dto.LoginResponse, error) {
	ctx := context.Background()

	user, err := a.userRepo.FindUserByEmail(ctx, data.Email)
	if err != nil {
		return nil, err
	}

	if err := a.bcryptService.ComparePassword(
		user.Password,
		data.Password,
	); err != nil {
		return nil, err
	}

	accessToken, errAccess := a.jwtService.CreateAccessToken(int(user.ID))
	if errAccess != nil {
		return nil, errAccess
	}

	refreshToken, errRefresh := a.jwtService.CreateRefreshToken(int(user.ID))
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
