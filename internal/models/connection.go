package models

type Connection struct {
	OrderID   int    `json:"order_id"`
	CreatedAt string `json:"created_at"`
	UserID    string `json:"user_id"`
	DeviceID  string `json:"device_id"`
}
