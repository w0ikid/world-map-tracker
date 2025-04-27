package services

import (
	"context"
	"github.com/w0ikid/world-map-tracker/internal/domain/models"
	"github.com/w0ikid/world-map-tracker/internal/domain/repositories"
)

type CountryStatusesServiceInterface interface {
	CreateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error)
	GetCountryStatuses(ctx context.Context, userID int) ([]*models.CountryStatus, error)
	UpdateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error)
	DeleteCountryStatus(ctx context.Context, userID int, countryISO string) error
	GetVisitedCount(ctx context.Context, userID int) (int, error)
	FindUsersWithSimilarList(ctx context.Context, userID int) ([]*models.User, error)
	GetWishListCount(ctx context.Context, userID int) (int, error)
	GetTopFiveVisitedCountries(ctx context.Context) ([]*models.TopCountry, error)
	GetTopFiveWishlistCountries(ctx context.Context) ([]*models.TopCountry, error)
}

type CountryStatusesService struct {
	repo repositories.CountryStatusesRepositoryInterface
}

func NewCountryStatusesService(repo repositories.CountryStatusesRepositoryInterface) *CountryStatusesService {
	return &CountryStatusesService{repo: repo}
}
func (s *CountryStatusesService) CreateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error) {
	countryStatus, err := s.repo.CreateCountryStatus(ctx, countryStatus)
	if err != nil {
		return nil, err
	}
	return countryStatus, nil
}

func (s *CountryStatusesService) GetCountryStatuses(ctx context.Context, userID int) ([]*models.CountryStatus, error) {
	countryStatuses, err := s.repo.GetCountryStatuses(ctx, userID)
	if err != nil {
		return nil, err
	}
	return countryStatuses, nil
}

func (s *CountryStatusesService) UpdateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error) {
	countryStatus, err := s.repo.UpdateCountryStatus(ctx, countryStatus)
	if err != nil {
		return nil, err
	}
	return countryStatus, nil
}

func (s *CountryStatusesService) DeleteCountryStatus(ctx context.Context, userID int, countryISO string) error {
	err := s.repo.DeleteCountryStatus(ctx, userID, countryISO)
	if err != nil {
		return err
	}
	return nil
}

func (s *CountryStatusesService) GetVisitedCount(ctx context.Context, userID int) (int, error) {
	visitedCount, err := s.repo.GetVisitedCount(ctx, userID)
	if err != nil {
		return 0, err
	}
	return visitedCount, nil
}

func (s *CountryStatusesService) FindUsersWithSimilarList(ctx context.Context, userID int) ([]*models.User, error) {
	users, err := s.repo.FindUsersWithSimilarList(ctx, userID)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *CountryStatusesService) GetWishListCount(ctx context.Context, userID int) (int, error) {
	wishListCount, err := s.repo.GetWishListCount(ctx, userID)
	if err != nil {
		return 0, err
	}
	return wishListCount, nil
}

func (s *CountryStatusesService) GetTopFiveVisitedCountries(ctx context.Context) ([]*models.TopCountry, error) {
	topCountries, err := s.repo.GetTopFiveVisitedCountries(ctx)
	if err != nil {
		return nil, err
	}
	return topCountries, nil
}

func (s *CountryStatusesService) GetTopFiveWishlistCountries(ctx context.Context) ([]*models.TopCountry, error) {
	topCountries, err := s.repo.GetTopFiveWishlistCountries(ctx)
	if err != nil {
		return nil, err
	}
	return topCountries, nil
}