package main

import (
	"devwithgo/controllers"
	"net/http"

	psh "github.com/platformsh/gohelper"
)

func main() {
	// PlatformSH
	platformSH, err := psh.NewPlatformInfo()
	if err != nil {
		panic("Not in a Platform.sh environment")
	}

	// Controllers
	controllers.Home()

	// Local env
	// http.ListenAndServe(":80", nil)

	// Platform SH env
	http.ListenAndServe(":"+platformSH.Port, nil)
}
