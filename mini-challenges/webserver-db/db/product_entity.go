package db

import (
    "database/sql"
    "log"
    "time"
)

type Product struct {
    ID uint
    Name string
    CreatedAt  *time.Time
    UpdatedAt *time.Time
}

type ProductWithVariant struct {
    Product Product
    Variant []Variant
}

func ProductSchemaInit(db *sql.DB) error {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS product (ID INTEGER CONSTRAINT product_id_pk PRIMARY KEY AUTOINCREMENT, NAME VARCHAR(50) NOT NULL, CREATED_AT TIMESTAMP NOT NULL, UPDATED_AT TIMESTAMP)")
    if nil != err {
        log.Fatal(err)

        return err
    }

    return nil
}

/**
* save new product data to data source
* id, created_at, and updated_at value will be ignored
*/
func ProductSave(db *sql.DB, p Product) (uint, error) {
    row, err := db.Exec("INSERT INTO product (NAME, CREATED_AT) VALUES (?, DATETIME())", p.Name)

    if nil != err {
        log.Println(err)

        return 0, err
    }

    result, err := row.LastInsertId()

    if nil != err {
        log.Println(err)

        return 0, err
    }

    return uint(result), nil
}

/**
* get all existing product data from data source
*/
func ProductGet(db *sql.DB, page uint, limit uint) ([]Product, error) {
    row, err := db.Query("SELECT ID, NAME, CREATED_AT, UPDATED_AT FROM product LIMIT ? OFFSET ?", limit, (page-1) * limit)

    if nil != err {
        log.Println(err)

        return []Product{}, err
    }

    result := []Product{}

    for row.Next() {
        mProduct := Product{}
        row.Scan(&mProduct.ID, &mProduct.Name, &mProduct.CreatedAt, &mProduct.UpdatedAt)

        result = append(result, mProduct)
    }

    return result, nil
}

/**
* get existing product data by id from data source
*/
func ProductGetById(db *sql.DB, id uint) (Product, error) {
    row := db.QueryRow("SELECT ID, NAME, CREATED_AT, UPDATED_AT FROM product WHERE ID = ?", id)

    result := Product{}
    row.Scan(&result.ID, &result.Name, &result.CreatedAt, &result.UpdatedAt)

    return result, nil
}

/**
* update existing product data by id from data source with new data
* ID, CreatedAt, UpdatedAt will be ignored
*/
func ProductUpdateById(db *sql.DB, id uint, p Product) (uint, error) {
    row, err := db.Exec("UPDATE product SET NAME = ?, UPDATED_AT = DATETIME() WHERE ID = ?", p.Name, id)

    if nil != err {
        log.Println(err)

        return 0, err
    }

    result, err := row.RowsAffected()

    if nil != err {
        log.Println(err)

        return 0, err
    }

    return uint(result), nil
}

/**
* delete product data to data source
* will return either the data is exist or not
*/
func ProductDeleteById(db *sql.DB, id uint) (uint, error) {
    _, err := db.Exec("DELETE FROM product WHERE id = ?", id)

    if nil != err {
        log.Println(err)

        return 0, err
    }

    return id, nil
}

/**
* get all existing product data from data source
* will eager load variants
*/
func ProductWithVariantGet(db *sql.DB, page uint, limit uint) ([]ProductWithVariant, error) {
    row, err := db.Query("SELECT ID, NAME, CREATED_AT, UPDATED_AT FROM product LIMIT ? OFFSET ?", limit, (page-1) * limit)

    if nil != err {
        log.Println(err)

        return []ProductWithVariant{}, err
    }

    result := []ProductWithVariant{}

    for row.Next() {
        mProductWithVariant := ProductWithVariant{}

        mProduct := Product{}
        row.Scan(&mProduct.ID, &mProduct.Name, &mProduct.CreatedAt, &mProduct.UpdatedAt)

        mVariants, _ := VariantGetByProductId(_db, mProduct.ID)

        mProductWithVariant.Product = mProduct
        mProductWithVariant.Variant = mVariants

        result = append(result, mProductWithVariant)
    }

    return result, nil
}

