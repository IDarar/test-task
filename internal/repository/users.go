package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/IDarar/test-task/internal/domain"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewUsersRepo(db *pgxpool.Pool) *UsersPostgres {
	return &UsersPostgres{
		db: db,
	}
}

type UsersPostgres struct {
	db *pgxpool.Pool
}

func (u *UsersPostgres) Create(name string) (domain.User, error) {
	createdAt := time.Now()
	row := u.db.QueryRow(context.Background(), "insert into users (name, created_at) values ($1, $2) RETURNING id", name, createdAt)

	user := domain.User{}

	err := row.Scan(&user.ID)
	if err != nil {
		//check if violation of unique constraint
		if pgErr, ok := err.(*pgconn.PgError); ok && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return domain.User{}, domain.ErrUserAlreadyExists
		} else {
			return domain.User{}, fmt.Errorf("failed create user: %w", err)
		}
	}

	user.Name = name
	user.CreatedAt = createdAt

	return user, nil
}

func (p *UsersPostgres) Delete(id int) error {
	res, err := p.db.Exec(context.Background(), "delete from users where id = $1", id)
	if err != nil {
		return fmt.Errorf("failed delete user: %w", err)
	}

	if res.RowsAffected() == 0 {
		return domain.ErrUserDoesNotExist
	}

	return nil
}

func (p *UsersPostgres) List() ([]domain.User, error) {
	users := []domain.User{}

	rows, err := p.db.Query(context.Background(), "select id, name, created_at from users")
	if err != nil {
		return users, fmt.Errorf("failed get users list: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.CreatedAt); err != nil {
			return users, fmt.Errorf("failed scan rows into user struct: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}
