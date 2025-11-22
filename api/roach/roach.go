package roach

import (
	"context"
	"log"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"api/models"
)

var (
	Connnection *pgx.Conn
)

func Connect() error {
	config, err := pgx.ParseConfig("postgres://go_user:go_password@192.168.0.213:26257/go_demo")

	if err != nil {
		log.Fatal(err)
		return err
	}

	conn, err := pgx.ConnectConfig(context.Background(), config)

	if err != nil {
		log.Fatal(err)
		return err
	}

	Connnection = conn

	return nil
}

func Disconnect() error {
	err := Connnection.Close(context.Background())

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func CreateUser(user models.User) error {
	err := crdbpgx.ExecuteTx(context.Background(), Connnection, pgx.TxOptions{}, func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(), "INSERT INTO users (id, name, password) VALUES ($1, $2, $3)", uuid.New(), user.Name, user.Password)

		if err != nil {
			log.Fatalln(err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
