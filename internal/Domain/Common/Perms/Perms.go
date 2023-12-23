package Perms

type Perm byte

const (
	// user
	LIST_USERS      Perm = 1
	GET_USER_DETAIL Perm = 2
	UPDATE_USER     Perm = 3
	DELETE_USER     Perm = 4
	ADD_USER        Perm = 5

	// categories
	LIST_CATEGORIES           Perm = 6
	GET_CATEGORY_DETAIL       Perm = 7
	UPDATE_CATEGORY           Perm = 8
	ADD_ITEM_TO_CATEGORY      Perm = 9
	REMOVE_ITEM_FROM_CATEGORY Perm = 10

	// Products, Items
	LIST_PRODUCTS      Perm = 11
	GET_PRODUCT_DETAIL Perm = 12
	ADD_PRODUCT        Perm = 13
	UPDATE_PRODUCT     Perm = 14
	DELETE_PRODUCT     Perm = 15
)
