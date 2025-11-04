package structs

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type Restaurant struct {
	ID      int    `json:"id"`
	Name    string `json:"name" binding:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type MenuItem struct {
	ID           int    `json:"id"`
	RestaurantID int    `json:"restaurant_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Price        int    `json:"price" binding:"required"`
	Available    bool   `json:"available"`
}

type CreateOrderItem struct {
	MenuItemID int `json:"menu_item_id" binding:"required"`
	Quantity   int `json:"quantity" binding:"required,min=1"`
}

type CreateOrderRequest struct {
	RestaurantID int               `json:"restaurant_id" binding:"required"`
	Items        []CreateOrderItem `json:"items" binding:"required,min=1"`
}

type Order struct {
	ID           int              `json:"id"`
	UserID       int              `json:"user_id"`
	RestaurantID int              `json:"restaurant_id"`
	TotalPrice   int              `json:"total_price"`
	Status       string           `json:"status"`
	Items        []OrderItemPrint `json:"items,omitempty"`
}

type OrderItemPrint struct {
	MenuItemID int    `json:"menu_item_id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	PriceEach  int    `json:"price_each"`
	Subtotal   int    `json:"subtotal"`
}
