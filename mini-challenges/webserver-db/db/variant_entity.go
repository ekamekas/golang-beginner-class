package db

import (
    "database/sql"
    "log"
    "time"
)

type Variant struct {
    ID uint
    Name string
    Quantity uint
    ProductID uint
    CreatedAt  *time.Time
    UpdatedAt *time.Time
}

func VariantSchemaInit(db *sql.DB) error {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS variant (ID INTEGER CONSTRAINT variant_id_pk PRIMARY KEY AUTOINCREMENT, NAME VARCHAR(50) NOT NULL, QUANTITY INTEGER NOT NULL DEFAULT 0, PRODUCT_ID INTEGER REFERENCES product (ID) ON DELETE CASCADE, CREATED_AT TIMESTAMP NOT NULL, UPDATED_AT TIMESTAMP)")
    if nil != err {
        log.Fatal(err)

        return err
    }

    return nil
}

/**
* save new variant data to data source
* will try to save new product if id is not empty
* id, created_at, and updated_at value will be ignored
*/
func VariantSave(db *sql.DB, p Variant) (uint, error) {
    var row sql.Result
    if 0 == p.ProductID {
        row, err = db.Exec("INSERT INTO variant (NAME, QUANTITY, CREATED_AT) VALUES (?, ?, DATETIME())", p.Name, p.Quantity)
        if nil != err {
            log.Println(err)

            return 0, err
        }

    } else {
        row, err = db.Exec("INSERT INTO variant (NAME, QUANTITY, PRODUCT_ID, CREATED_AT) VALUES (?, ?, ?, DATETIME())", p.Name, p.Quantity, p.ProductID)
        if nil != err {
            log.Println(err)

            return 0, err
        }
    }

    result, err := row.LastInsertId()
    if nil != err {
        log.Println(err)

        return 0, err
    }


    return uint(result), nil
}

/**
* get all existing variant data from data source
*/
func VariantGet(db *sql.DB, page uint, limit uint) ([]Variant, error) {
    row, err := db.Query("SELECT ID, NAME, QUANTITY, COALESCE(PRODUCT_ID, 0), CREATED_AT, UPDATED_AT FROM variant LIMIT ? OFFSET ?", limit, (page-1) * limit)

    if nil != err {
        log.Println(err)

        return []Variant{}, err
    }

    result := []Variant{}

    for row.Next() {
        mVariant := Variant{}
        row.Scan(&mVariant.ID, &mVariant.Name, &mVariant.Quantity, &mVariant.ProductID, &mVariant.CreatedAt, &mVariant.UpdatedAt)

        result = append(result, mVariant)
    }

    return result, nil
}

/**
* get all existing variant data by proudct id from data source
*/
func VariantGetByProductId(db *sql.DB, productID uint) ([]Variant, error) {
    row, err := db.Query("SELECT ID, NAME, QUANTITY, PRODUCT_ID, CREATED_AT, UPDATED_AT FROM variant WHERE PRODUCT_ID = ?", productID)

    if nil != err {
        log.Println(err)

        return []Variant{}, err
    }

    result := []Variant{}

    for row.Next() {
        mVariant := Variant{}
        row.Scan(&mVariant.ID, &mVariant.Name, &mVariant.Quantity, &mVariant.ProductID, &mVariant.CreatedAt, &mVariant.UpdatedAt)

        result = append(result, mVariant)
    }

    return result, nil
}


/**
* update existing variant data by id from data source with new data
* ID, CreatedAt, UpdatedAt will be ignored
*/
func VariantUpdateById(db *sql.DB, id uint, p Variant) (uint, error) {
    var row sql.Result
    if 0 == p.ProductID {
        row, err = db.Exec("UPDATE variant SET NAME = ?, QUANTITY = ?, PRODUCT_ID = NULl, UPDATED_AT = DATETIME() WHERE ID = ?", p.Name, p.Quantity, id)
        if nil != err {
            log.Println(err)

            return 0, err
        }

    } else {
        row, err = db.Exec("UPDATE variant SET NAME = ?, QUANTITY = ?, PRODUCT_ID = ?, UPDATED_AT = DATETIME() WHERE ID = ?", p.Name, p.Quantity, p.ProductID, id)
        if nil != err {
            log.Println(err)

            return 0, err
        }
    }

    result, err := row.RowsAffected()

    if nil != err {
        log.Println(err)

        return 0, err
    }

    return uint(result), nil
}

/**
* delete variant data to data source
* will return either the data is exist or not
*/
func VariantDeleteById(db *sql.DB, id uint) (uint, error) {
    _, err := db.Exec("DELETE FROM variant WHERE id = ?", id)

    if nil != err {
        log.Println(err)

        return 0, err
    }

    return id, nil
}
