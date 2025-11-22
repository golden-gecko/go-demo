package roach

import (
	"context"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	log "github.com/sirupsen/logrus"

	"api/models"
)

var (
    Connnection *pgx.Conn
)

func Connect() error {
    log.Println("Connecting to database...")

    config, err := pgx.ParseConfig("postgres://go_user:go_password@haproxy:26257/go_demo")

    if err != nil {
        log.Error(err)
        return err
    }

    conn, err := pgx.ConnectConfig(context.Background(), config)

    if err != nil {
        log.Error(err)
        return err
    }

    Connnection = conn

    return nil
}

func Disconnect() error {
    log.Println("Disconnecting to database...")

    err := Connnection.Close(context.Background())

    if err != nil {
        log.Error(err)
        return err
    }

    return nil
}

func CreateUser(user models.User) error {
    err := crdbpgx.ExecuteTx(context.Background(), Connnection, pgx.TxOptions{}, func(tx pgx.Tx) error {
        _, err := tx.Exec(context.Background(), "INSERT INTO users (id, name, password) VALUES ($1, $2, $3)", uuid.New(), user.Name, user.Password)

        if err != nil {
            log.Error(err)
            return err
        }

        return nil
    })

    if err != nil {
        log.Error(err)
        return err
    }

    return nil
}

func CreateVehicle(vehicle models.Vehicle) error {
    err := crdbpgx.ExecuteTx(context.Background(), Connnection, pgx.TxOptions{}, func(tx pgx.Tx) error {
        _, err := tx.Exec(context.Background(), "INSERT INTO vehicles (plate, brand, model, year, category) VALUES ($1, $2, $3, $4, $5)", vehicle.Plate, vehicle.Brand, vehicle.Model, vehicle.Year, vehicle.Category)

        if err != nil {
            log.Error(err)
            return err
        }

        return nil
    })

    if err != nil {
        log.Error(err)
        return err
    }

    return nil
}

func GetUser(userId uuid.UUID) (models.User, error) {
    var user models.User

    err := crdbpgx.ExecuteTx(context.Background(), Connnection, pgx.TxOptions{}, func(tx pgx.Tx) error {
        err := tx.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", userId).Scan(&user)

        if err != nil {
            log.Error(err)
            return err
        }

        return nil
    })

    if err != nil {
        log.Error(err)
        return user, err
    }

    return user, nil
}
