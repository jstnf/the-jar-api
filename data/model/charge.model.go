package model

import (
    "github.com/google/uuid"
)

type Charge struct {
    Id           int       `json:"id"`
    Reason       uuid.UUID `json:"reason"`
    Victim       string    `json:"victim"`
    Accuser      string    `json:"accuser"`
    CustomAmount bool      `json:"amount_override" default:"false"`
    Amount       float64   `json:"amount" default:"0.0"`
    Timestamp    int64     `json:"timestamp" default:"0"`
    Deleted      bool      `json:"deleted" default:"false"`
}
