package http

import "time"

type VariantCreateRequest struct {
    Name string
    Quantity uint
    ProductID uint
}

type VariantCreateResponse struct {
    ID uint
}

type VariantGetResponse struct {
    ID uint
    Name string
    Quantity uint
    ProductID uint
    CreatedAt  *time.Time
    UpdatedAt *time.Time

}

type VariantUpdateRequest struct {
    Name string
    Quantity uint
    ProductID uint

}

type VariantUpdateResponse struct {
    ID uint
}

type VariantDeleteResponse struct {
    ID uint
}
