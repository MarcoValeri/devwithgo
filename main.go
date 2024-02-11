package main

import (
	"devwithgo/controllers"
	"devwithgo/util"
	"net/http"

	psh "github.com/platformsh/gohelper"
)

func main() {
	// PlatformSH
	platformSH, err := psh.NewPlatformInfo()
	if err != nil {
		panic("Not in a Platform.sh environment")
	}

	// Static files
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	// Controllers
	controllers.Home()
	controllers.AdminController()

	// Database Local
	// util.DatabaseLocalConnection()

	// Database Platform sh
	util.DatabaseConnectionPlatformSh()

	// Local env
	// http.ListenAndServe(":80", nil)

	// Platform SH env
	http.ListenAndServe(":"+platformSH.Port, nil)
}
