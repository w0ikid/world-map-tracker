package usecase

import (
	"context"
	"github.com/w0ikid/world-map-tracker/internal/domain/models"
	"github.com/w0ikid/world-map-tracker/internal/domain/services"
)

type CountryStatusesUseCaseInterface interface {
	CreateCountryStatus(ctx context.Context, input *CountryStatusInput) (*models.CountryStatus, error)
	GetCountryStatuses(ctx context.Context, userID int) ([]*models.CountryStatus, error)
	UpdateCountryStatus(ctx context.Context, input *CountryStatusInput) (*models.CountryStatus, error)
	DeleteCountryStatus(ctx context.Context, userID int, countryISO string) error
	GetVisitedPercentage(ctx context.Context, userID int) (int, error)
	GetVisitedCount(ctx context.Context, userID int) (int, error)
	FindUsersWithSimilarList(ctx context.Context, userID int) ([]*models.User, error)
	GetWishListCount(ctx context.Context, userID int) (int, error)
	GetTopFiveVisitedCountries(ctx context.Context) ([]*models.TopCountry, error)
	GetTopFiveWishlistCountries(ctx context.Context) ([]*models.TopCountry, error)
}

type CountryStatusesUseCase struct {
	service services.CountryStatusesServiceInterface
}

func NewCountryStatusesUseCase(service services.CountryStatusesServiceInterface) *CountryStatusesUseCase {
	return &CountryStatusesUseCase{service: service}
}

type CountryStatusInput struct {
	UserID     int    `json:"user_id" validate:"required"`
	CountryISO string    `json:"country_iso" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

func (u *CountryStatusesUseCase) CreateCountryStatus(ctx context.Context, input *CountryStatusInput) (*models.CountryStatus, error) {
	countryStatus := &models.CountryStatus{
		UserID:     input.UserID,
		CountryISO: input.CountryISO,
		Status:     input.Status,
	}

	countryStatus, err := u.service.CreateCountryStatus(ctx, countryStatus)
	if err != nil {
		return nil, err
	}
	return countryStatus, nil
}

func (u *CountryStatusesUseCase) GetCountryStatuses(ctx context.Context, userID int) ([]*models.CountryStatus, error) {
	countryStatuses, err := u.service.GetCountryStatuses(ctx, userID)
	if err != nil {
		return nil, err
	}
	return countryStatuses, nil
}

func (u *CountryStatusesUseCase) UpdateCountryStatus(ctx context.Context, input *CountryStatusInput) (*models.CountryStatus, error) {
	countryStatus := &models.CountryStatus{
		UserID:     input.UserID,
		CountryISO: input.CountryISO,
		Status:     input.Status,
	}

	countryStatus, err := u.service.UpdateCountryStatus(ctx, countryStatus)
	if err != nil {
		return nil, err
	}
	return countryStatus, nil
}

func (u *CountryStatusesUseCase) DeleteCountryStatus(ctx context.Context, userID int, countryISO string) error {
	err := u.service.DeleteCountryStatus(ctx, userID, countryISO)
	if err != nil {
		return err
	}
	return nil
}

func (u *CountryStatusesUseCase) GetVisitedPercentage(ctx context.Context, userID int) (int, error) {
	visitedCount, err := u.service.GetVisitedCount(ctx, userID)
	if err != nil {
		return 0, err
	}
	var totalCountries = 195
	percentage := (visitedCount * 100) / totalCountries
	return percentage, nil
}

func (u *CountryStatusesUseCase) GetVisitedCount(ctx context.Context, userID int) (int, error) {
	visitedCount, err := u.service.GetVisitedCount(ctx, userID)
	if err != nil {
		return 0, err
	}
	return visitedCount, nil
}

func (u *CountryStatusesUseCase) FindUsersWithSimilarList(ctx context.Context, userID int) ([]*models.User, error) {
	users, err := u.service.FindUsersWithSimilarList(ctx, userID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *CountryStatusesUseCase) GetWishListCount(ctx context.Context, userID int) (int, error) {
	wishListCount, err := u.service.GetWishListCount(ctx, userID)
	if err != nil {
		return 0, err
	}
	return wishListCount, nil
}

func (u *CountryStatusesUseCase) GetTopFiveVisitedCountries(ctx context.Context) ([]*models.TopCountry, error) {
	topCountries, err := u.service.GetTopFiveVisitedCountries(ctx)
	if err != nil {
		return nil, err
	}
	return topCountries, nil
}

func (u *CountryStatusesUseCase) GetTopFiveWishlistCountries(ctx context.Context) ([]*models.TopCountry, error) {
	topCountries, err := u.service.GetTopFiveWishlistCountries(ctx)
	if err != nil {
		return nil, err
	}
	return topCountries, nil
}