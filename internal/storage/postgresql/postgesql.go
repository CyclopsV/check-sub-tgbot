package postgresql

import (
	"context"
	"fmt"

	"github.com/CyclopsV/check-sub-tgbot/internal/configs"
	"github.com/jackc/pgx/v5"
)

type DB struct {
	url string
}

func New(c configs.DBConfig) (DB, error) {
	dbURL := fmt.Sprintf("postgres://%v:%v@localhost:5432/%v", c.Username, c.Password, c.Title)
	db := DB{url: dbURL}
	return db, db.checkDB()
}

func (db DB) checkDB() error {
	ctx := context.Background()
	defer ctx.Done()
	con, err := pgx.Connect(ctx, db.url)
	if err != nil {
		return err
	}
	defer con.Close(ctx)
	if err = con.Ping(ctx); err != nil {
		return err
	}
	return nil
}
