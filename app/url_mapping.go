package app

import "bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
}
