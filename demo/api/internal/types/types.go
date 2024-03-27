// Code generated by goctl. DO NOT EDIT.
package types

type CreateOrderRequest struct {
	UserID int64   `json:"user_id"`
	Goods  []Goods `json:"goods"`
}

type CreateOrderResponse struct {
	Url string `json:"url"`
}

type Goods struct {
	GoodID int64 `json:"good_id"`
	Num    int64 `json:"num"`
}

type OrderNotifRequest struct {
}

type OrderNotifResponse struct {
	Message string `json:"message"`
}