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
	query := fmt.Sprintf(`
	INSERT INTO %s (
	                username, --1 
	                email, 	  --2
	                password  --3
	                ) 
	VALUES ($1, $2, $3)
`, usersTable)
	_, err := p.DB.ExecContext(c, query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
