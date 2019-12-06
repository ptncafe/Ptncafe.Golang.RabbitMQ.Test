package model

import "time"

type StoreDto struct {
	Id int
	Name string
	Code string
	ShopStatus int
	UpdatedDate time.Time
}