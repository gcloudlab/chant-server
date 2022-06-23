package main

import "chant/router"

func main() {
	e := router.Router()
	e.Run(":8088")
}
