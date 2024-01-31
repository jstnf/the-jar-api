package model

import "github.com/google/uuid"

type Reason struct {
    Id        int       `json:"id"`
    Jar       uuid.UUID `json:"jar"`
    Name      string    `json:"name"`
    Amount    float64   `json:"amount" default:"1.0"`
    Timestamp int64     `json:"timestamp" default:"0"`
    Deleted   bool      `json:"deleted" default:"false"`
}
