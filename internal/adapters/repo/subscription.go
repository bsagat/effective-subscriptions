package repo

import (
	"context"
	"errors"
	"fmt"
	"submanager/internal/domain"
	"time"

	"github.com/jackc/pgx/v5"
)

type SubsRepo struct {
	db *pgx.Conn
}

func NewSubsRepo(db *pgx.Conn) *SubsRepo {
	return &SubsRepo{
		db: db,
	}
}

func (repo *SubsRepo) IsUnique(ctx context.Context, serviceName, userID string) (bool, error) {
	const op = "SubsRepo.IsUnique"
	query := `
		SELECT COUNT(*) = 0 FROM Subscriptions
		WHERE Service_name = $1 AND User_ID = $2`

	var unique bool
	if err := repo.db.QueryRow(ctx, query, serviceName, userID).Scan(&unique); err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return unique, nil
}

// Creates a new subscription in the database
func (repo *SubsRepo) Create(ctx context.Context, subs domain.Subscription) error {
	const op = "SubsRepo.Create"
	query := `
		INSERT INTO Subscriptions(Service_name, Price, User_ID, Start_date, Exp_date)
		VALUES($1, $2, $3, $4, $5)
		RETURNING User_ID;
	`

	_, err := repo.db.Exec(ctx, query, subs.ServiceName, subs.Price, subs.UserID, subs.StartDate, subs.EndDate)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *SubsRepo) Get(ctx context.Context, serviceName, userID string) (domain.Subscription, error) {
	const op = "SubsRepo.Get"
	query := `
		SELECT Service_name, Price, User_ID, Start_date, Exp_date 
		FROM Subscriptions
		WHERE Service_name = $1 AND User_ID = $2;
	`
	var subs domain.Subscription
	if err := repo.db.QueryRow(ctx, query, serviceName, userID).Scan(&subs.ServiceName, &subs.Price, &subs.UserID, &subs.StartDate, &subs.EndDate); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Subscription{}, domain.ErrSubsNotFound
		}
		return domain.Subscription{}, fmt.Errorf("%s: %w", op, err)
	}
	return subs, nil
}

func (repo *SubsRepo) List(ctx context.Context, userID string) ([]domain.Subscription, error) {
	const op = "SubsRepo.List"
	query := `
		SELECT Service_name, Price, User_ID, Start_date, Exp_date 
		FROM Subscriptions
		WHERE User_ID = $1;
	`
	rows, err := repo.db.Query(ctx, query, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrSubsNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	subsList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (domain.Subscription, error) {
		var subs domain.Subscription
		if err := row.Scan(&subs.ServiceName, &subs.Price, &subs.UserID, &subs.StartDate, &subs.EndDate); err != nil {
			return domain.Subscription{}, fmt.Errorf("%s: %w", op, err)
		}
		return subs, nil
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, rows.Err())
	}

	return subsList, nil
}

func (repo *SubsRepo) Update(ctx context.Context, subs domain.Subscription) error {
	const op = "SubsRepo.Update"
	query := `
		UPDATE Subscriptions
		SET Price = $1, Start_date = $2, Exp_date = $3
		WHERE Service_name = $4 AND User_ID = $5;`

	res, err := repo.db.Exec(ctx, query, subs.Price, subs.StartDate, subs.EndDate, subs.ServiceName, subs.UserID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if res.RowsAffected() == 0 {
		return domain.ErrSubsNotFound
	}
	return nil
}

func (repo *SubsRepo) Delete(ctx context.Context, serviceName, userID string) error {
	const op = "SubsRepo.Delete"
	query := `
		DELETE FROM Subscriptions
		WHERE Service_name = $1 AND User_ID = $2;`

	res, err := repo.db.Exec(ctx, query, serviceName, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if res.RowsAffected() == 0 {
		return domain.ErrSubsNotFound
	}
	return nil
}

func (repo *SubsRepo) DeleteList(ctx context.Context, userID string) error {
	const op = "SubsRepo.DeleteList"
	query := `
		DELETE FROM Subscriptions
		WHERE User_ID = $1;`

	res, err := repo.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if res.RowsAffected() == 0 {
		return domain.ErrSubsNotFound
	}
	return nil
}

func (repo *SubsRepo) SubsListByFilter(ctx context.Context, start, end time.Time, serviceName, userID string, pageNum, pageSize int) ([]domain.Subscription, error) {
	const op = "SubsRepo.SubsListByFilter"
	query := `
		SELECT Service_name, Price, User_ID, Start_date, Exp_date FROM Subscriptions
		WHERE Start_date BETWEEN $1 and $2 `

	// Add filters and args dynamically
	args := []any{start, end}
	switch {
	case len(serviceName) != 0:
		query += `AND Service_name = $3 `
		args = append(args, serviceName)
		if len(userID) != 0 {
			query += `AND User_ID = $4 `
			args = append(args, userID)
		}
	case len(userID) != 0:
		query += `AND User_ID = $3 `
		args = append(args, userID)
	}

	offset := (pageNum - 1) * pageSize
	query += fmt.Sprintf(`
	ORDER BY start_date DESC
    LIMIT %d OFFSET %d;`, pageSize, offset)

	rows, err := repo.db.Query(ctx, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrSubsNotFound
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	subsList, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (domain.Subscription, error) {
		var subs domain.Subscription
		if err := row.Scan(&subs.ServiceName, &subs.Price, &subs.UserID, &subs.StartDate, &subs.EndDate); err != nil {
			return domain.Subscription{}, fmt.Errorf("%s: %w", op, err)
		}
		return subs, nil
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return subsList, nil
}
