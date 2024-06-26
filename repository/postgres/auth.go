package postgres

import (
	"context"
	"fmt"
	"github.com/dlankinl/bmstu-ppo-bl/domain"
	repoDomain "github.com/dlankinl/bmstu-ppo-da/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) repoDomain.IAuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) Register(ctx context.Context, authInfo *domain.UserAuth) (err error) {
	query := `insert into ppo.users (username, password, role) values ($1, $2, 'user')`

	_, err = r.db.Exec(
		ctx,
		query,
		authInfo.Username,
		authInfo.HashedPass,
	)
	if err != nil {
		return fmt.Errorf("регистрация пользователя: %w", err)
	}

	return nil
}

func (r *AuthRepository) GetByUsername(ctx context.Context, username string) (data *domain.UserAuth, err error) {
	query := `select password, role from ppo.users where username = $1`

	tmp := new(UserAuth)
	err = r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&tmp.HashedPass,
		&tmp.Role,
	)
	if err != nil {
		return nil, fmt.Errorf("получение пользователя по username: %w", err)
	}

	return UserAuthDbToUserAuth(tmp), nil
}
