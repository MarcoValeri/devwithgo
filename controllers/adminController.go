package controllers

import admincontrollers "devwithgo/controllers/adminControllers"

func AdminController() {
	admincontrollers.AdminLogin()
	admincontrollers.AdminDashboard()
}
