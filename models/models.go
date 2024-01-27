package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id"`
	First_Name      *string
	Last_Name       *string
	Password        *string
	Email           *string
	Phone           *string
	Token           *string
	Refresh_Token   *string
	Created_At      time.Time
	Updated_At      time.Time
	User_ID         string
	UserCart        []ProductUser
	Address_Details []Address
	Order_Status    []Order
}

type Product struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        *uint64
	Rating       *uint8
	Image        *string
}

type ProductUser struct {
	Product_ID   primitive.ObjectID
	Product_Name *string
	Price        *uint64
	Rating       *uint8
	Image        *string
}

type Address struct {
	Address_ID primitive.ObjectID
	House      *string
	Street     *string
	City       *string
	Pincode    *string
}

type Order struct {
	Order_ID       primitive.ObjectID
	Order_Cart     []ProductUser
	Ordered_At     time.Time
	Price          *uint64
	Discount       *uint64
	Payment_Method Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
