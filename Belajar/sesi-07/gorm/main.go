package main

import "learning-gorm/database"

func main() {
	database.StartDB()

	//create user
	//getUserById(1)
	// updateUserById(1, "qwerty@gmail.com")
	// createProduct(1, "qwe", "YYY")
	// getUserWithProducts()
	deleteProductById(1)
}
