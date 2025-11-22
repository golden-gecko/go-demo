package models

// for JSON

type Car struct {
    Year     int    `bson:"Year"`
    Brand    string `bson:"Brand"`
    Model    string `bson:"Model"`
    Category string `bson:"Category"`
}

type Cars struct {
    Cars []Car `json:"cars"`
}

type Names struct {
    Names []string `json:"names"`
}

// for API

type Coordinate struct {
    Latitude  float32 `bson:"Latitude"`
    Longitude float32 `bson:"Longitude"`
}

type Temperature struct {
    Location  string  `bson:"Location"`
    Value     float32 `bson:"Value"`
    Timestamp string  `bson:"Timestamp"`
}

type Transit struct {
    Plate     string     `bson:"Plate"`
    Location  Coordinate `bson:"Location"`
    Timestamp string     `bson:"Timestamp"`
}

type User struct {
    Name     string `bson:"Name"`
    Password string `bson:"Password"`
}

type Vehicle struct {
    Plate    string `bson:"Plate"`
    Brand    string `bson:"Brand"`
    Model    string `bson:"Model"`
    Year     int    `bson:"Year"`
    Category string `bson:"Category"`
}
