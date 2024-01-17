package roach

import (
	"context"

	"github.com/jackc/pgx/v4"

	log "github.com/sirupsen/logrus"
)

var (
    Connnection *pgx.Conn
)

type VehicleCountByModel struct {
    Brand string `bson:"Brand"`
    Model string `bson:"Model"`
    Count int    `bson:"Count"`
}

type VehicleCountByYear struct {
    Year  int `bson:"Year"`
    Count int `bson:"Count"`
}

func Connect() error {
    log.Println("Connecting to database...")

    config, err := pgx.ParseConfig("postgres://go_user:go_password@cockroach_1:26257/go_demo")

    if err != nil {
        log.Error(err)
        return err
    }

    conn, err := pgx.ConnectConfig(context.Background(), config)

    if err != nil {
        log.Error(err)
        return err
    }

    log.Println("Connected to database")

    Connnection = conn

    return nil
}

func Disconnect() error {
    log.Println("Disconnecting from database...")

    err := Connnection.Close(context.Background())

    if err != nil {
        log.Error(err)
        return err
    }

    log.Println("Disconnected from database")

    return nil
}

func GetVehicleCountByModel() ([]VehicleCountByModel, error) {
    rows, err := Connnection.Query(context.Background(), "SELECT brand, model, COUNT(*) FROM vehicles GROUP BY brand, model")

    if err != nil {
        log.Error(err)
        return nil, err
    }

    defer rows.Close()

    var models []VehicleCountByModel

    for rows.Next() {
        var model VehicleCountByModel

        if err := rows.Scan(&model.Brand, &model.Model, &model.Count); err != nil {
            return nil, err
        }

        models = append(models, model)
    }

    return models, nil
}

func GetVehicleCountByYear() ([]VehicleCountByYear, error) {
    rows, err := Connnection.Query(context.Background(), "SELECT year, COUNT(*) FROM vehicles GROUP BY year")

    if err != nil {
        log.Error(err)
        return nil, err
    }

    defer rows.Close()

    var years []VehicleCountByYear

    for rows.Next() {
        var year VehicleCountByYear

        if err := rows.Scan(&year.Year, &year.Count); err != nil {
            return nil, err
        }

        years = append(years, year)
    }

    return years, nil
}
