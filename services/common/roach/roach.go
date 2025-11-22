package roach

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	log "github.com/sirupsen/logrus"

	"services/common/model"
)

var (
	connnection *pgx.Conn
)

func Connect() error {
    log.Info("Connecting to database...")

    config, err := pgx.ParseConfig("postgres://go_user:go_password@cockroach_1:26257/go_demo")

    if err != nil {
        return err
    }

    conn, err := pgx.ConnectConfig(context.Background(), config)

    if err != nil {
        return err
    }

    log.Info("Connected to database")

    connnection = conn

    return nil
}

func Disconnect() error {
    log.Info("Disconnecting from database...")

    err := connnection.Close(context.Background())

    if err != nil {
        return err
    }

    log.Info("Disconnected from database")

    return nil
}

func CreateUser(user model.User) error {
	_, err := connnection.Exec(context.Background(), "INSERT INTO users (id, name, password) VALUES ($1, $2, $3)", uuid.New(), user.Name, user.Password)

	if err != nil {
		return err
	}

    return nil
}

func CreateVehicle(vehicle model.Vehicle) error {
	_, err := connnection.Exec(context.Background(), "INSERT INTO vehicles (plate, brand, model, year, category) VALUES ($1, $2, $3, $4, $5)", vehicle.Plate, vehicle.Brand, vehicle.Model, vehicle.Year, vehicle.Category)

	if err != nil {
		return err
	}

    return nil
}

func GetUser(userId uuid.UUID) (model.User, error) {
    var user model.User

	if err := connnection.QueryRow(context.Background(), "SELECT * FROM users WHERE id = $1", userId).Scan(&user); err != nil {
		return user, err
	}

    return user, nil
}

func GetVehicleCountByModel() ([]model.VehicleCountByModel, error) {
    rows, err := connnection.Query(context.Background(), "SELECT brand, model, COUNT(*) FROM vehicles GROUP BY brand, model")

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var models []model.VehicleCountByModel

    for rows.Next() {
        var model model.VehicleCountByModel

        if err := rows.Scan(&model.Brand, &model.Model, &model.Count); err != nil {
            return nil, err
        }

        models = append(models, model)
    }

    return models, nil
}

func GetVehicleCountByYear() ([]model.VehicleCountByYear, error) {
    rows, err := connnection.Query(context.Background(), "SELECT year, COUNT(*) FROM vehicles GROUP BY year")

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    var years []model.VehicleCountByYear

    for rows.Next() {
        var year model.VehicleCountByYear

        if err := rows.Scan(&year.Year, &year.Count); err != nil {
            return nil, err
        }

        years = append(years, year)
    }

    return years, nil
}
