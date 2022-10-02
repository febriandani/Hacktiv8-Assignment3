package main

import "Assignment3/router"

func main() {
	r := router.StartApp()
	r.Run(":9000")
}
