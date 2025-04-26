package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/w0ikid/world-map-tracker/internal/domain/models"
)

type CountryStatusesRepositoryInterface interface {
	CreateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error)
	GetCountryStatuses(ctx context.Context, userID int) ([]*models.CountryStatus, error)
	UpdateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error)
	DeleteCountryStatus(ctx context.Context, userID int, countryISO string) error
	GetVisitedCount(ctx context.Context, userID int) (int, error)
}

type CountryStatusesRepository struct {
	db *pgxpool.Pool
}

func NewCountryStatusesRepository(db *pgxpool.Pool) *CountryStatusesRepository {
	return &CountryStatusesRepository{db: db}
}

func (r *CountryStatusesRepository) CreateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error) {
	query := `INSERT INTO country_statuses (user_id, country_iso, status) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(ctx, query, countryStatus.UserID, countryStatus.CountryISO, countryStatus.Status).Scan(&countryStatus.ID)
	if err != nil {
		return nil, err
	}
	return countryStatus, nil
}

func (r *CountryStatusesRepository) GetCountryStatuses(ctx context.Context, userID int) ([]*models.CountryStatus, error) {
	query := `SELECT id, user_id, country_iso, status FROM country_statuses WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var countryStatuses []*models.CountryStatus
	for rows.Next() {
		countryStatus := &models.CountryStatus{}
		err := rows.Scan(&countryStatus.ID, &countryStatus.UserID, &countryStatus.CountryISO, &countryStatus.Status)
		if err != nil {
			return nil, err
		}
		countryStatuses = append(countryStatuses, countryStatus)
	}

	return countryStatuses, nil
}

func (r *CountryStatusesRepository) UpdateCountryStatus(ctx context.Context, countryStatus *models.CountryStatus) (*models.CountryStatus, error) {
	query := `UPDATE country_statuses SET status = $1 WHERE user_id = $2 AND country_iso = $3 RETURNING id`
	err := r.db.QueryRow(ctx, query, countryStatus.Status, countryStatus.UserID, countryStatus.CountryISO).Scan(&countryStatus.ID)
	if err != nil {
		return nil, err
	}
	return countryStatus, nil
}

func (r *CountryStatusesRepository) DeleteCountryStatus(ctx context.Context, userID int, countryISO string) error {
	query := `DELETE FROM country_statuses WHERE user_id = $1 AND country_iso = $2`
	_, err := r.db.Exec(ctx, query, userID, countryISO)
	if err != nil {
		return err
	}
	return nil
}

func (r *CountryStatusesRepository) GetVisitedCount(ctx context.Context, userID int) (int, error) {
	query := `SELECT COUNT(*) FROM country_statuses WHERE user_id = $1 AND status = 'visited'`
	var count int
	err := r.db.QueryRow(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

