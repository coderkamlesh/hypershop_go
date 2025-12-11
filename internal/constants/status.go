package constants

type OrderStatus string

const (
	OrderPending    OrderStatus = "PENDING"
	OrderConfirmed  OrderStatus = "CONFIRMED"
	OrderPacked     OrderStatus = "PACKED"
	OrderDispatched OrderStatus = "DISPATCHED"
	OrderDelivered  OrderStatus = "DELIVERED"
	OrderCancelled  OrderStatus = "CANCELLED"
)

type ProductStatus string

const (
	ProductActive     ProductStatus = "ACTIVE"
	ProductInactive   ProductStatus = "INACTIVE"
	ProductOutOfStock ProductStatus = "OUT_OF_STOCK"
)

const (
	SUCCESS       = 1
	FAILED        = 0
	INVALID_TOKEN = 3
)

// Status messages
const (
	SuccessMessage      = "Success"
	FailedMessage       = "Failed"
	InvalidTokenMessage = "Invalid or expired token"
)
