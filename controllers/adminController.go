package controllers

import admincontrollers "devwithgo/controllers/adminControllers"

func AdminController() {
	admincontrollers.AdminLogin()

	admincontrollers.AdminDashboard()

	admincontrollers.AdminUsers()
	admincontrollers.AdminUserAdd()
	admincontrollers.AdminUserEdit()
}
