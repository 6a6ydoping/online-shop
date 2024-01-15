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
