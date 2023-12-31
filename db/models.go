// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type CouponScope string

const (
	CouponScopeGlobal CouponScope = "global"
	CouponScopeShop   CouponScope = "shop"
)

func (e *CouponScope) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CouponScope(s)
	case string:
		*e = CouponScope(s)
	default:
		return fmt.Errorf("unsupported scan type for CouponScope: %T", src)
	}
	return nil
}

type NullCouponScope struct {
	CouponScope CouponScope `json:"coupon_scope"`
	Valid       bool        `json:"valid"` // Valid is true if CouponScope is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCouponScope) Scan(value interface{}) error {
	if value == nil {
		ns.CouponScope, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CouponScope.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCouponScope) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CouponScope), nil
}

type CouponType string

const (
	CouponTypePercentage CouponType = "percentage"
	CouponTypeFixed      CouponType = "fixed"
	CouponTypeShipping   CouponType = "shipping"
)

func (e *CouponType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CouponType(s)
	case string:
		*e = CouponType(s)
	default:
		return fmt.Errorf("unsupported scan type for CouponType: %T", src)
	}
	return nil
}

type NullCouponType struct {
	CouponType CouponType `json:"coupon_type"`
	Valid      bool       `json:"valid"` // Valid is true if CouponType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCouponType) Scan(value interface{}) error {
	if value == nil {
		ns.CouponType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CouponType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCouponType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CouponType), nil
}

type OrderStatus string

const (
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusFinished  OrderStatus = "finished"
)

func (e *OrderStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = OrderStatus(s)
	case string:
		*e = OrderStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for OrderStatus: %T", src)
	}
	return nil
}

type NullOrderStatus struct {
	OrderStatus OrderStatus `json:"order_status"`
	Valid       bool        `json:"valid"` // Valid is true if OrderStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullOrderStatus) Scan(value interface{}) error {
	if value == nil {
		ns.OrderStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.OrderStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullOrderStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.OrderStatus), nil
}

type RoleType string

const (
	RoleTypeAdmin    RoleType = "admin"
	RoleTypeCustomer RoleType = "customer"
)

func (e *RoleType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = RoleType(s)
	case string:
		*e = RoleType(s)
	default:
		return fmt.Errorf("unsupported scan type for RoleType: %T", src)
	}
	return nil
}

type NullRoleType struct {
	RoleType RoleType `json:"role_type"`
	Valid    bool     `json:"valid"` // Valid is true if RoleType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoleType) Scan(value interface{}) error {
	if value == nil {
		ns.RoleType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.RoleType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoleType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.RoleType), nil
}

type Cart struct {
	ID     int32 `json:"id" param:"cart_id"`
	UserID int32 `json:"user_id"`
	ShopID int32 `json:"shop_id"`
}

type CartCoupon struct {
	CartID   int32 `json:"cart_id"`
	CouponID int32 `json:"coupon_id"`
}

type CartProduct struct {
	CartID    int32 `json:"cart_id"`
	ProductID int32 `json:"product_id" param:"id"`
	Quantity  int32 `json:"quantity"`
}

type Coupon struct {
	ID          int32              `json:"id" param:"id"`
	Type        CouponType         `json:"type"`
	Scope       CouponScope        `json:"scope"`
	ShopID      pgtype.Int4        `json:"shop_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Discount    pgtype.Numeric     `json:"discount" swaggertype:"number"`
	StartDate   pgtype.Timestamptz `json:"start_date" swaggertype:"string"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
}

type CouponTag struct {
	CouponID int32 `json:"coupon_id" param:"id"`
	TagID    int32 `json:"tag_id"`
}

type OrderDetail struct {
	OrderID        int32 `json:"order_id" param:"id"`
	ProductID      int32 `json:"product_id"`
	ProductVersion int32 `json:"product_version"`
	Quantity       int32 `json:"quantity"`
}

type OrderHistory struct {
	ID         int32              `json:"id" param:"id"`
	UserID     int32              `json:"user_id"`
	ShopID     int32              `json:"shop_id"`
	Shipment   int32              `json:"shipment"`
	TotalPrice int32              `json:"total_price"`
	Status     OrderStatus        `json:"status"`
	CreatedAt  pgtype.Timestamptz `json:"created_at" swaggertype:"string"`
}

type Product struct {
	ID          int32              `json:"id" param:"id"`
	Version     int32              `json:"version"`
	ShopID      int32              `json:"shop_id"`
	Name        string             `form:"name" json:"name"`
	Description string             `form:"description" json:"description"`
	Price       pgtype.Numeric     `json:"price" swaggertype:"number"`
	ImageID     string             `json:"image_id"`
	ExpireDate  pgtype.Timestamptz `json:"expire_date" swaggertype:"string"`
	EditDate    pgtype.Timestamptz `json:"edit_date" swaggertype:"string"`
	Stock       int32              `form:"stock" json:"stock"`
	Sales       int32              `json:"sales"`
	Enabled     bool               `form:"enabled" json:"enabled"`
}

type ProductArchive struct {
	ID          int32          `json:"id"`
	Version     int32          `json:"version"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       pgtype.Numeric `json:"price" swaggertype:"number"`
	ImageID     string         `json:"image_id"`
}

type ProductTag struct {
	TagID     int32 `json:"tag_id"`
	ProductID int32 `json:"product_id" param:"id"`
}

type Shop struct {
	ID          int32  `json:"id"`
	SellerName  string `json:"seller_name" param:"seller_name"`
	ImageID     string `json:"image_id" swaggertype:"string"`
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
	Enabled     bool   `form:"enabled" json:"enabled"`
}

type Tag struct {
	ID     int32  `json:"id"`
	ShopID int32  `json:"shop_id"`
	Name   string `json:"name"`
}

type User struct {
	ID                     int32              `json:"id" param:"id"`
	Username               string             `json:"username"`
	Password               string             `json:"password"`
	Name                   string             `form:"name" json:"name"`
	Email                  string             `form:"email" json:"email"`
	Address                string             `form:"address" json:"address"`
	ImageID                string             `json:"image_id" swaggertype:"string"`
	Role                   RoleType           `json:"role"`
	CreditCard             json.RawMessage    `json:"credit_card"`
	RefreshToken           string             `json:"refresh_token"`
	Enabled                bool               `json:"enabled"`
	RefreshTokenExpireDate pgtype.Timestamptz `json:"refresh_token_expire_date"`
}
