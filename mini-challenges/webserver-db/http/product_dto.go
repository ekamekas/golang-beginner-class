package http

import "time"

type ProductCreateRequest struct {
    Name string
}

type ProductCreateResponse struct {
    ID uint
}

type ProductGetResponse struct {
    ID uint
    Name string
    CreatedAt  *time.Time
    UpdatedAt *time.Time

}

type ProductUpdateRequest struct {
    Name string
}

type ProductUpdateResponse struct {
    ID uint
}

type ProductDeleteResponse struct {
    ID uint
}

type ProductWithVariantGetResponse struct {
    ID uint
    Name string
    Variant []VariantGetResponse
    CreatedAt  *time.Time
    UpdatedAt *time.Time

}

