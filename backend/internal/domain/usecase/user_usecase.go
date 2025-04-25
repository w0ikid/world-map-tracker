package usecase

import (
	"context"
	"github.com/w0ikid/world-map-tracker/internal/domain/models"
	"github.com/w0ikid/world-map-tracker/internal/domain/services"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"
)

type UserUseCaseInterface interface {
	CreateUser(ctx context.Context, user *UserInput) (*models.User, error)
	LoginUser(ctx context.Context, email, password string) (*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserInput struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUseCase struct {
	service services.UserServiceInterface
}

func NewUserUseCase(service services.UserServiceInterface) *UserUseCase {
	return &UserUseCase{service: service}
}

func (u *UserUseCase) CreateUser(ctx context.Context, user *UserInput) (*models.User, error) {
	validator := validator.New()
	if err := validator.Struct(user); err != nil {
		return nil, err
	}

	userModel := &models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	createdUser, err := u.service.CreateUser(ctx, userModel)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (u *UserUseCase) LoginUser(ctx context.Context, email, password string) (*models.User, error) {
	user, err := u.service.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	user, err := u.service.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := u.service.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	user, err := u.service.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	updatedUser, err := u.service.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (u *UserUseCase) DeleteUser(ctx context.Context, id int) error {
	err := u.service.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}