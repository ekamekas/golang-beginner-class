package http

import (
    "net/http"
    "encoding/json"
    "strings"
    "strconv"
    
    "webserver-db/db"
)

func VariantControllerInit() {
    http.HandleFunc("/api/variant/", func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/api/variant/"), 10, 32)
        if nil != err {
            http.Error(w, "Request is not valid", http.StatusBadRequest)

            return
        }

        switch r.Method {
            case "PUT":
                var dto VariantUpdateRequest

                err := json.NewDecoder(r.Body).Decode(&dto)

                if nil != err {
                    http.Error(w, "Request is not valid", http.StatusBadRequest)

                    return
                }

                result, err := variantUpdateById(uint(id), dto)
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                if 0 == result.ID {
                    http.Error(w, "Variant is not found", http.StatusNotFound)

                    return
                }

                json.NewEncoder(w).Encode(result)

                break
            case "DELETE":
                result, err := variantDeleteById(uint(id))
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

    http.HandleFunc("/api/variant", func(w http.ResponseWriter, r *http.Request) {
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

                result, err := variantGet(page, limit)
                if nil != err {
                    http.Error(w, "Server could not process request", http.StatusInternalServerError)

                    return
                }

                json.NewEncoder(w).Encode(result)

                break
            case "POST":
                var dto VariantCreateRequest

                err := json.NewDecoder(r.Body).Decode(&dto)

                if nil != err {
                    http.Error(w, "Request is not valid", http.StatusBadRequest)

                    return
                }

                result, err := variantCreate(dto)
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

func variantCreate(p VariantCreateRequest) (VariantCreateResponse, error) {
    _db, _ := db.GetOrInit()
    result := VariantCreateResponse{}

    mVariantID, err := db.VariantSave(_db, db.Variant{ Name: p.Name, Quantity: p.Quantity, ProductID: p.ProductID })

    result = VariantCreateResponse{ ID: mVariantID }

    return result, err
}

func variantGet(page uint, limit uint) ([]VariantGetResponse, error) {
    _db, _ := db.GetOrInit()

    mVariants, err := db.VariantGet(_db, page, limit)
    result := []VariantGetResponse{}

    for _, mVariant := range mVariants {
        result = append(result, VariantGetResponse{ ID: mVariant.ID, Name: mVariant.Name, Quantity: mVariant.Quantity, ProductID: mVariant.ProductID, CreatedAt: mVariant.CreatedAt, UpdatedAt: mVariant.UpdatedAt })
    }

    return result, err
}

func variantUpdateById(id uint, p VariantUpdateRequest) (VariantUpdateResponse, error) {
    _db, _ := db.GetOrInit()

    mVariantID, err := db.VariantUpdateById(_db, id, db.Variant{ Name: p.Name, Quantity: p.Quantity, ProductID: p.ProductID })

    result := VariantUpdateResponse{ ID: id}

    if 0 == mVariantID {
        result.ID = mVariantID
    }

    return result, err
}

func variantDeleteById(id uint) (VariantDeleteResponse, error) {
    _db, _ := db.GetOrInit()

    mVariantID, err := db.VariantDeleteById(_db, id)

    result := VariantDeleteResponse{ ID: mVariantID}

    return result, err
}
