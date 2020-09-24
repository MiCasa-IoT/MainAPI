package models

type Connection struct {
	OrderId   int    `json:"order_id"`
	CreatedAt string `json:"created_at"`
	UserId    string `json:"user_id"`
	DeviceId  string `json:"device_id"`
}
