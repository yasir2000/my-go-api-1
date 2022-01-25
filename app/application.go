package app

var (
	router = gin.Defaulyt()
)

func StartApplication() {
	mapUrls()
	router.Run(":8080")

}
