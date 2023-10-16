package db

import (
	"database/sql"
	"errors"
	"log"

	"webserver-http/domain"
)

type orderRepository struct {
	database *sql.DB
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &orderRepository{database: db}
}

func (r *orderRepository) Create(m *domain.Order) (uint, error) {
	tx, err := r.database.Begin()
	if nil != err {
		log.Println("[ORDER] failed to create new order", err)

		return 0, nil
	}

	defer tx.Rollback() // will be ignored if transaction is committed

	row, err := tx.Exec(`
        INSERT INTO "order" (CUSTOMER_NAME, ORDERED_AT) VALUES (?, ?)
    `, m.CustomerName, m.OrderedAt)
	if nil != err {
		log.Println("[ORDER] failed to create new order", err)

		return 0, nil
	}

	orderId, err := row.LastInsertId()
	if nil != err {
		log.Println("[ORDER] failed to create new order", err)

		return 0, nil
	}

	if nil == m.Items {
		tx.Commit()

		return uint(orderId), nil
	}

	for _, item := range *m.Items {
		_, err := tx.Exec(`
            INSERT INTO item (NAME, DESCRIPTION, QUANTITY, ORDER_ID) VALUES (?, ?, ?, ?) 
        `, item.Name, item.Description, item.Quantity, orderId)
		if nil != err {
			log.Println("[ORDER] failed to create new item", err)

			return 0, nil
		}
	}

	tx.Commit()

	return uint(orderId), nil
}
func (r *orderRepository) Get() ([]domain.Order, error) {
	result := []domain.Order{}

	row, err := r.database.Query(`SELECT ID, CUSTOMER_NAME, ORDERED_AT, CREATED_AT, UPDATED_AT FROM "order"`)
	if nil != err {
		log.Println("[ORDER] failed to get orders")

		return result, err
	}

	for row.Next() {
		order := domain.Order{}

		row.Scan(&order.ID, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)

		// eager fetch items
		itemRow, err := r.database.Query("SELECT ID, NAME, DESCRIPTION, QUANTITY, CREATED_AT, UPDATED_AT FROM item WHERE ORDER_ID = ?", order.ID)
		if nil != err {
			log.Println("[ORDER] failed to get items for order", order.ID)

			return result, err
		}

		items := []domain.Item{}

		for itemRow.Next() {
			item := domain.Item{}

			itemRow.Scan(&item.ID, &item.Name, &item.Description, &item.Quantity, &item.CreatedAt, &item.UpdatedAt)

			items = append(items, item)
		}

		order.Items = &items

		result = append(result, order)
	}

	return result, nil
}

func (r *orderRepository) Update(id uint, m *domain.Order) (uint, error) {
	tx, err := r.database.Begin()
	if nil != err {
		log.Println("[ORDER] failed to update order", err)

		return 0, err
	}

	defer tx.Rollback() // will be ignored if transaction is committed

	row, err := tx.Exec(`
		UPDATE "order" SET CUSTOMER_NAME = ?, ORDERED_AT = ? WHERE ID = ? 
    `, m.CustomerName, m.OrderedAt, id)
	if nil != err {
		log.Println("[ORDER] failed to update order", err)

		return 0, err
	}

	updatedRows, err := row.RowsAffected()
	if nil != err {
		log.Println("[ORDER] failed to update order", err)

		return 0, err
	}
	if 0 == updatedRows {
		err = errors.New("resource is not found")
		log.Println("[ORDER] failed to update order", err)

		return 0, err
	}

	_, err = tx.Exec(`
		DELETE FROM item WHERE ORDER_ID = ?
    `, id)
	if nil != err {
		log.Println("[ORDER] failed to create order", err)

		return 0, err
	}

	if nil == m.Items {
		tx.Commit()

		return id, nil
	}

	for _, item := range *m.Items {
		_, err := tx.Exec(`
            INSERT INTO item (NAME, DESCRIPTION, QUANTITY, ORDER_ID) VALUES (?, ?, ?, ?) 
        `, item.Name, item.Description, item.Quantity, id)
		if nil != err {
			log.Println("[ORDER] failed to create item", err)

			return 0, err
		}
	}

	tx.Commit()

	return id, nil
}

func (r *orderRepository) Delete(id uint) (uint, error) {
	_, err := r.database.Exec(`
		DELETE FROM "order" WHERE id = ?
	`, id)
	if nil != err {
		log.Println("[ORDER] failed to delete order with id", id)

		return id, err
	}

	return id, nil
}
