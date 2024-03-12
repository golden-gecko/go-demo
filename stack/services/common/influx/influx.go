package influx

import (
	"context"
	"fmt"
	"os"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	log "github.com/sirupsen/logrus"

	"services/common"
)

var (
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
)

func Connect() error {
	host := os.Getenv("INFLUX_HOST")
	port := os.Getenv("INFLUX_PORT")
	secret := os.Getenv("INFLUX_SECRET")
	organization := os.Getenv("INFLUX_ORGANIZATION")
	database := os.Getenv("INFLUX_DATABASE")

	uri := fmt.Sprintf("http://%s:%s", host, port)

    log.Info(fmt.Sprintf("Connecting to InfluxDB (%s)...", uri))

	client := influxdb2.NewClientWithOptions(uri, secret, influxdb2.DefaultOptions().SetBatchSize(10))

	if _, err := client.Health(context.Background()); err != nil {
		return err
	}

    log.Info("Connected to InfluxDB")

	writeAPI = client.WriteAPIBlocking(organization, database)

	return nil
}

func Disconnect() {
    log.Info("Disonnecting from InfluxDB...")

	client.Close()

    log.Info("Disonnected from InfluxDB")
}

func WritePoint(point *write.Point) error {
	return writeAPI.WritePoint(context.Background(), point);
}

func WriteTemperature(location string, value float32, timestamp string) error {
	timestampParsed, err := common.StringToTimestamp(timestamp)

	if err != nil {
		return err
	}

    return WritePoint(influxdb2.NewPointWithMeasurement("temperatue").
		AddTag("location", location).
        AddTag("unit", "celsius").
        AddField("current", value).
        SetTime(timestampParsed),
	)
}
