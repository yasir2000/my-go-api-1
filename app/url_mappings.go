package app

func mapUrls() {
	router.POST("/users", users.Create)
}
