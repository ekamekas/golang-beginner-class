package http

import (
    "net/http"
    "encoding/json"
    "strings"
    "strconv"

    "webserver-db/db"
)

func ProductControllerInit() {
    http.HandleFunc("/api/product/with-variants", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case "GET":
                page := uint(1)
                pageQuery := r.URL.Query().Get("page")
                if 0 < len(pageQuery) {
                    mPage, err := strconv.ParseUint(pageQuery, 10, 32)
                    if nil != err {
                        http.Error(w, "Request is not valid", http.StatusBadRequest)

                        return
                    }

                    page = uint(mPage)
                }

                limit := uint(10)
                limitQuery := r.URL.Query().Get("limit")
                if 0 < len(limitQuery) {
                    mLimit, err := strconv.ParseUint(limitQuery, 10, 32)
                    if nil != err {
                        http.Error(w, "Request is not valid", http.StatusBadRequest)

                        return
                    }

                    limit = uint(mLimit)
                }

                result, err := productWithVariantGet(page, limit)
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                json.NewEncoder(w).Encode(result)
                break
            default:
                 http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)       
        }
    });

    http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/api/product/"), 10, 32)
        if nil != err {
            http.Error(w, "Request is not valid", http.StatusBadRequest)

            return
        }

        switch r.Method {
            case "GET":
                result, err := productGetById(uint(id))
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                if 0 == result.ID {
                    http.Error(w, "Product is not found", http.StatusNotFound)

                    return
                }

                json.NewEncoder(w).Encode(result)

                break
            case "PUT":
                var dto ProductUpdateRequest

                err := json.NewDecoder(r.Body).Decode(&dto)

                if nil != err {
                    http.Error(w, "Request is not valid", http.StatusBadRequest)

                    return
                }

                result, err := productUpdateById(uint(id), dto)
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                if 0 == result.ID {
                    http.Error(w, "Product is not found", http.StatusNotFound)

                    return
                }

                json.NewEncoder(w).Encode(result)

                break
            case "DELETE":
                result, err := productDeleteById(uint(id))
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                json.NewEncoder(w).Encode(result)

                break

            default:
                 http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)       
        }
    });

    http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case "GET":
                page := uint(1)
                pageQuery := r.URL.Query().Get("page")
                if 0 < len(pageQuery) {
                    mPage, err := strconv.ParseUint(pageQuery, 10, 32)
                    if nil != err {
                        http.Error(w, "Request is not valid", http.StatusBadRequest)

                        return
                    }

                    page = uint(mPage)
                }

                limit := uint(10)
                limitQuery := r.URL.Query().Get("limit")
                if 0 < len(limitQuery) {
                    mLimit, err := strconv.ParseUint(limitQuery, 10, 32)
                    if nil != err {
                        http.Error(w, "Request is not valid", http.StatusBadRequest)

                        return
                    }

                    limit = uint(mLimit)
                }

                result, err := productGet(page, limit)
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                json.NewEncoder(w).Encode(result)

                break
            case "POST":
                var dto ProductCreateRequest

                err := json.NewDecoder(r.Body).Decode(&dto)

                if nil != err {
                    http.Error(w, "Request is not valid", http.StatusBadRequest)

                    return
                }

                result, err := productCreate(dto)
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                json.NewEncoder(w).Encode(result)

                break
            default:
                 http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)       
        }
    });

}

func productCreate(p ProductCreateRequest) (ProductCreateResponse, error) {
    _db, _ := db.GetOrInit()

    mProductID, err := db.ProductSave(_db, db.Product{ Name: p.Name })

    result := ProductCreateResponse{ ID: mProductID}

    return result, err
}

func productGet(page uint, limit uint) ([]ProductGetResponse, error) {
    _db, _ := db.GetOrInit()

    mProducts, err := db.ProductGet(_db, page, limit)

    result := []ProductGetResponse{}

    for _, mProduct := range mProducts {
        result = append(result, ProductGetResponse{ ID: mProduct.ID, Name: mProduct.Name, CreatedAt: mProduct.CreatedAt, UpdatedAt: mProduct.UpdatedAt })
    }

    return result, err
}

func productGetById(id uint) (ProductGetResponse, error) {
    _db, _ := db.GetOrInit()

    mProduct, err := db.ProductGetById(_db, id)

    result := ProductGetResponse{ ID: mProduct.ID, Name: mProduct.Name, CreatedAt: mProduct.CreatedAt, UpdatedAt: mProduct.UpdatedAt }

    return result, err
}

func productUpdateById(id uint, p ProductUpdateRequest) (ProductUpdateResponse, error) {
    _db, _ := db.GetOrInit()

    mProductID, err := db.ProductUpdateById(_db, id, db.Product{ Name: p.Name })

    result := ProductUpdateResponse{ ID: id}

    if 0 == mProductID {
        result.ID = mProductID
    }

    return result, err
}

func productDeleteById(id uint) (ProductDeleteResponse, error) {
    _db, _ := db.GetOrInit()

    mProductID, err := db.ProductDeleteById(_db, id)

    result := ProductDeleteResponse{ ID: mProductID}

    return result, err
}

func productWithVariantGet(page uint, limit uint) ([]ProductWithVariantGetResponse, error) {
    _db, _ := db.GetOrInit()

    mProducts, err := db.ProductWithVariantGet(_db, page, limit)

    result := []ProductWithVariantGetResponse{}

    for _, mProduct := range mProducts {
        mVariants := []VariantGetResponse{}

        for _, mVariant := range mProduct.Variant {
            mVariants = append(mVariants, VariantGetResponse{ ID: mVariant.ID, Name: mVariant.Name, Quantity: mVariant.Quantity, ProductID: mVariant.ProductID, CreatedAt: mVariant.CreatedAt, UpdatedAt: mVariant.UpdatedAt })
        }

        result = append(result, ProductWithVariantGetResponse{ ID: mProduct.Product.ID, Name: mProduct.Product.Name, Variant: mVariants, CreatedAt: mProduct.Product.CreatedAt, UpdatedAt: mProduct.Product.UpdatedAt })
    }

    return result, err
}

