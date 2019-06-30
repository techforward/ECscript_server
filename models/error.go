package models

import "errors"

// Common Validation
var (
	ErrorAmountInvalid = errors.New("Amount is bad number")
	ErrorPriceInvalid  = errors.New("Price is bad number")
)

// Address Validation
var (
	ErrorAddressUlidInvalid = errors.New("Address Ulid is empty")
	ErrorPostcodeInvalid    = errors.New("Postcode validation error")
	ErrorAddressInvalid     = errors.New("Address validation error")
)

// CartItem Validation
var (
	ErrorCartItemUlidInvalid = errors.New("CartItem Ulid is empty")
	ErrorIsDeletedInvalid    = errors.New("IsDeleted validation error")
	ErrorIsKeepInvalid       = errors.New("IsKeep validation error")
)

// Cart Validation
var (
	ErrorCartUlidInvalid = errors.New("Cart Ulid is empty")
	ErrorBoughtAtInvalid = errors.New("BoughtAt validation error")
)

// ItemImage Validation
var (
	ErrorItemImageUlidInvalid = errors.New("ItemImage Ulid is empty")
	ErrorPathInvalid          = errors.New("Path validation error")
)

// Item Validation
var (
	ErrorItemUlidInvalid = errors.New("Item Ulid is empty")
	ErrorNameInvalid     = errors.New("Name is empty")
	ErrorContextInvalid  = errors.New("Context is empty")
	ErrorPriorityInvalid = errors.New("Priority is bad number")
)

// OrderItem Validation
var (
	ErrorOrderItemUlidInvalid = errors.New("OrderItem Ulid is empty")
)

// Order Validation
var (
	ErrorOrderUlidInvalid  = errors.New("Order Ulid is empty")
	ErrorIsCanceledInvalid = errors.New("IsCanceled validation error")
	ErrorStatusInvalid     = errors.New("Status validation error")
)

// Price Validation
var (
	ErrorPriceUlidInvalid = errors.New("Price Ulid is empty")
)

// UserAddress Validation
var (
	ErrorUserAddressUlidInvalid = errors.New("UserAddress Ulid is empty")
)

// User Validation
var (
	ErrorUserUlidInvalid = errors.New("User Ulid is empty")
	ErrorEmailInvalid    = errors.New("Email validation error")
)
