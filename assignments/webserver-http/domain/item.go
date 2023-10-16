package domain

import "time"

type Item struct {
    ID uint
    Name string
    Description string
    Quantity uint
    CreatedAt *time.Time
    UpdatedAt *time.Time
}
