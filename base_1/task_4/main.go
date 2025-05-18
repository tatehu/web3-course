package main

import (
	"personal_blog/model"
	"personal_blog/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
