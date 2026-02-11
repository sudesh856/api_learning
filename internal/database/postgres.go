package database

import (
	"context"
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool ,error) {

	var ctx context.Context = context.Background()
	var config *pgxpool.Config
	var err error	


	config, err = pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("we got an error here and right now!!!", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("We again got an error here!!!",err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {

		log.Printf("some error has come now!!!", err)
		return nil, err

	}

	 log.Println("The database is connected!!!")
	 return pool, nil
	 
}