package main

import "dotcount/router"

func main()  {
	engine := router.NewRouter()
	engine.Run("/9090")
}