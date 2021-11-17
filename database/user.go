package database

import (
	"Toodoo/model"
	"context"
	"database/sql"

	"github.com/google/uuid"
)

func (c *Client) RegisterUser(ctx context.Context, user *model.User) (*model.User, error) {
	// buildSqlStatment("user", user.UserId, user.DeviceId, user.Token, user.Create_at)
	sqlStatment := `
	SELECT * FROM "user"
	WHERE userid = $1
	`
	var u *model.User

	//Query from db
	err := c.db.QueryRowContext(ctx, sqlStatment, user.UserId).Scan(&u.UserId, &u.DeviceId, &u.Token, &u.Create_at, &u.Os, &u.Os_ver)

	if err == sql.ErrNoRows {
		c.logger.Errorw("RegisterUser", "error", err)
		token, _ := c.AddUser(ctx, user)
		c.logger.Infow("RegisterUser", "register token", token)
	}

	return u, nil
}

func (c *Client) AddUser(ctx context.Context, user *model.User) (string, error) {
	sqlStatment := `
	INSERT INTO "user"(uid, userid, deviceid, token, create_at, os, os_ver)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING token
	`

	var token string

	err := c.db.QueryRowContext(ctx, sqlStatment, uuid.New().String(), user.UserId, user.DeviceId, user.Token, user.Create_at, user.Os, user.Os_ver).Scan(&token)
	checkError(err)

	return token, nil
}