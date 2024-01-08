package main

import (
	"devwithgo/controllers"
	"net/http"
)

func main() {
	// Controllers
	controllers.Home()

	http.ListenAndServe(":80", nil)
}
