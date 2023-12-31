package domain

import "time"

type Order struct {
	ID           uint
	CustomerName string
	Items        *[]Item
	OrderedAt    *time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type OrderRepository interface {
	/**
	 * create new order data
	 *
	 * @param m a reference of Order
	 * @return  new order id or error
	 */
	Create(m *Order) (uint, error)

	/**
	 * get all order data. all order relations will be fetched eagerly
	 *
	 * @return  an array or orders or error
	 */
	Get() ([]Order, error)

	/**
	 * update if exist order data. will return error if the data if not exist.
	 * all the data referenced to this order will be deleted and re-inserted
	 *
	 * @return  updated order id or error
	 */
	Update(id uint, m *Order) (uint, error)

	/**
	 * delete if exist order data. will not error if the data if not exist.
	 * all the data referenced to this order will be deleted
	 *
	 * @return  deleted order id or error
	 */
	Delete(id uint) (uint, error)
}

type orderController struct {
	repository OrderRepository
}

func NewOrderController(r OrderRepository) *orderController {
	return &orderController{repository: r}
}

func (controller *orderController) Create(r *Order) Result {
	// validate
	if nil == r {
		return Result{Error: "request body must not be nil", Code: "400"}
	}

	if "" == r.CustomerName {
		return Result{Error: "customer name must not be nil", Code: "400"}
	}

	if nil == r.Items || 1 > len(*r.Items) {
		return Result{Error: "order must have at least 1 item", Code: "400"}
	}

	if nil == r.OrderedAt {
		return Result{Error: "order data must not be nil", Code: "400"}
	}
	// end of validate

	orderId, err := controller.repository.Create(r)
	if nil != err {
		return Result{Error: err.Error(), Code: "500"}
	}

	return Result{Error: "", Code: "200", Data: orderId}
}

func (controller *orderController) Update(id uint, r *Order) Result {
	// validate
	if nil == r {
		return Result{Error: "request body must not be nil", Code: "400"}
	}

	if "" == r.CustomerName {
		return Result{Error: "customer name must not be nil", Code: "400"}
	}

	if nil == r.Items || 1 > len(*r.Items) {
		return Result{Error: "order must have at least 1 item", Code: "400"}
	}

	if nil == r.OrderedAt {
		return Result{Error: "order data must not be nil", Code: "400"}
	}
	// end of validate

	orderId, err := controller.repository.Update(id, r)
	if nil != err {
		return Result{Error: err.Error(), Code: "500"}
	}

	return Result{Error: "", Code: "200", Data: orderId}
}

func (controller *orderController) Get() Result {
	order, err := controller.repository.Get()
	if nil != err {
		return Result{Error: err.Error(), Code: "500"}
	}

	return Result{Error: "", Code: "200", Data: order}
}

func (controller *orderController) Delete(id uint) Result {
	orderId, err := controller.repository.Delete(id)
	if nil != err {
		return Result{Error: err.Error(), Code: "500"}
	}

	return Result{Error: "", Code: "200", Data: orderId}
}
