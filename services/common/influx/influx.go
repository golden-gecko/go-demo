package influx

import (
	"context"

	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"

	"services/common"
)

var (
	Client   influxdb2.Client
	WriteAPI api.WriteAPIBlocking
)

func Connect() error {
	Client = influxdb2.NewClientWithOptions("http://influx:8086", "my-super-secret-auth-token", influxdb2.DefaultOptions().SetBatchSize(10))

	if _, err := Client.Health(context.Background()); err != nil {
		return err
	}

	WriteAPI = Client.WriteAPIBlocking("my-org", "house")

	return nil
}

func Disconnect() {
	Client.Close()
}

func WritePoint(point *write.Point) error {
	return WriteAPI.WritePoint(context.Background(), point);
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
