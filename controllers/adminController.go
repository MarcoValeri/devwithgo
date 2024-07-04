package controllers

import (
	admincontrollers "devwithgo/controllers/adminControllers"
)

func AdminController() {
	admincontrollers.AdminLogin()

	admincontrollers.AdminDashboard()

	admincontrollers.AdminUsers()
	admincontrollers.AdminUserAdd()
	admincontrollers.AdminUserDelete()
	admincontrollers.AdminUserEdit()

	admincontrollers.AdminGuides()
	admincontrollers.AdminGuideAdd()
	admincontrollers.AdminGuideDelete()
	admincontrollers.AdminGuideEdit()

	admincontrollers.AdminImages()
	admincontrollers.AdminImageAdd()
	admincontrollers.AdminImageEdit()
	admincontrollers.AdminImageDelete()
}
