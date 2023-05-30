package routes

import (
	"BNMO/controller"
	"BNMO/middleware"

	"github.com/gin-gonic/gin"
)

func MapUrls(Router *gin.Engine) {
	// Register account
	Router.POST("/register", controller.RegisterAccount)
	// Login account
	Router.POST("/login", controller.LoginAccount)
	// Logout accounts
	Router.POST("/logout", controller.LogoutAccount)
	// Server static folder (at /images)
	Router.Static("/images", "images")

	// Authorized side (both customer and admin)
	Router.Use(middleware.IsAuthenticate)
	// Poll endpoint for updating balance
	Router.GET("/update-balance", controller.UpdateBalance)
	// Get exchange symbols
	Router.GET("/get-symbols", controller.GetSymbols)

	/* CUSTOMER SIDE */
	/** REQUEST PAGE (ADD / SUBTRACT) **/
	// Request money (add / subtract)
	Router.POST("/request-money", controller.RequestMoney)
	
	/** TRANSFER PAGE **/
	// Get transfer destination based on id
	Router.GET("/get-destination", controller.GetDestination)
	// Check transfer destination in the database
	Router.POST("/check-destination", controller.CheckDestination)
	// Add transfer destination to user
	Router.POST("/add-destination", controller.AddDestination)
	// Transfer from source account to destination account
	Router.POST("/transfer", controller.Transfer)

	/** HISTORY PAGE **/
	// Get request history
	Router.GET("/request-history", controller.RequestHistory)
	// Get transfer history
	Router.GET("/transfer-history", controller.TransferHistory)

	/** PROFILE PAGE **/
	// Change image
	Router.PUT("/change-image", controller.ChangeImage)
	// Check PIN
	Router.GET("/check-pin", controller.CheckPin)
	// Change PIN (only once)
	Router.PUT("/change-pin", controller.ChangePin)
	// Change Password
	Router.PUT("/change-password", controller.ChangePassword)
	
	/* ADMIN SIDE */
	/** VALIDATE REQUEST PAGE **/
	// Display requests for admin
	Router.GET("/admin/get-pending-request", controller.DisplayRequests)
	// Validate selected requests for admin
	Router.POST("/admin/validate-request", controller.ValidateRequests)

	/** VALIDATE ACCOUNT PAGE **/
	// Display pending accounts for admin
	Router.GET("/admin/get-pending-account", controller.DisplayPendingAccount)
	// Validate selected accounts for admin
	Router.POST("/admin/validate-account", controller.ValidateAccount)
	
	/** CUSTOMER DATA PAGE **/
	// Display all customer data
	Router.GET("/admin/get-customer-data", controller.SendAllCustomerData)
	// Edit data
	Router.PUT("/admin/edit-data", controller.EditData)
	// Reset PIN
	Router.PUT("/admin/reset-pin", controller.ResetPIN)
	// Delete account
	Router.DELETE("admin/delete-account", controller.DeleteAccount)
}