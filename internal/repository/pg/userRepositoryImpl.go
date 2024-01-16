package pg

import (
	"context"
	"fmt"
	"github.com/6a6ydoping/online-shop/internal/entity"
	"time"
)

func (p *Postgres) GetAllUsers(ctx context.Context) (*[]entity.User, error) {
	var users []entity.User

	query := fmt.Sprintf(`
	SELECT user_id, username, email, password  FROM %s
`, usersTable)

	c, cancel := context.WithTimeout(ctx, 5*time.Second) // TODO: replace to conf
	defer cancel()

	err := p.DB.SelectContext(c, &users, query)
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (p *Postgres) CreateUser(c context.Context, user entity.User) error {
	// TODO: check if email or username already in db (now query broken)

	query := fmt.Sprintf(`
	INSERT INTO %s (username, email, password)
	VALUES ($1, $2, $3)
	ON CONFLICT (username) DO NOTHING;
`, usersTable)
	result, err := p.DB.ExecContext(c, query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user \"%s\" already exists", user.Username)
	}

	return nil
}
