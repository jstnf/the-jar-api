package model

import "github.com/google/uuid"

type Jar struct {
    Id        uuid.UUID `json:"id"`
    Guild     string    `json:"guild"`
    Name      string    `json:"name"`
    Currency  string    `json:"currency" default:"USD"`
    Timestamp int64     `json:"timestamp" default:"0"`
    Deleted   bool      `json:"deleted" default:"false"`
}
