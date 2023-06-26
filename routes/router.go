package routes

// func MapUrls(Router *gin.Engine) {
// 	// Register account
// 	Router.POST("/register", controllers.RegisterAccount)
// 	// Login account
// 	Router.POST("/login", controllers.LoginAccount)
// 	// Logout accounts
// 	Router.POST("/logout", controllers.LogoutAccount)
// 	// Server static folder (at /images)
// 	Router.Static("/images", "images")

// 	// Authorized side (both customer and admin)
// 	Router.Use(middleware.JWTAuthMiddleware())
// 	// Poll endpoint for updating balance
// 	Router.GET("/update-balance", controllers.UpdateBalance)
// 	// Get exchange symbols
// 	Router.GET("/get-symbols", controllers.GetSymbols)

// 	/* CUSTOMER SIDE */
// 	/** REQUEST PAGE (ADD / SUBTRACT) **/
// 	// Request money (add / subtract)
// 	Router.POST("/request-money", controllers.AddRequest)

// 	/** TRANSFER PAGE **/
// 	// Get transfer destination based on id
// 	Router.GET("/get-destination", controllers.GetDestination)
// 	// Check transfer destination in the database
// 	Router.POST("/check-destination", controllers.CheckDestination)
// 	// Add transfer destination to user
// 	Router.POST("/add-destination", controllers.AddDestination)
// 	// Transfer from source account to destination account
// 	Router.POST("/transfer", controllers.Transfer)

// 	/** HISTORY PAGE **/
// 	// Get request history
// 	Router.GET("/request-history", controllers.RequestHistory)
// 	// Get transfer history
// 	Router.GET("/transfer-history", controllers.TransferHistory)

// 	/** PROFILE PAGE **/
// 	// Change image
// 	Router.PUT("/change-image", controllers.ChangeImage)
// 	// Check PIN
// 	Router.GET("/check-pin", controllers.CheckPin)
// 	// Change PIN (only once)
// 	Router.PUT("/change-pin", controllers.ChangePin)
// 	// Change Password
// 	Router.PUT("/change-password", controllers.ChangePassword)

// 	/* ADMIN SIDE */
// 	/** VALIDATE REQUEST PAGE **/
// 	// Display requests for admin
// 	Router.GET("/admin/get-pending-request", controllers.GetPendingRequests)
// 	// Validate selected requests for admin
// 	Router.POST("/admin/validate-request", controllers.ValidateRequests)

// 	/** VALIDATE ACCOUNT PAGE **/
// 	// Display pending accounts for admin
// 	Router.GET("/admin/get-pending-account", controllers.DisplayPendingAccount)
// 	// Validate selected accounts for admin
// 	Router.POST("/admin/validate-account", controllers.ValidateAccount)

// 	/** CUSTOMER DATA PAGE **/
// 	// Display all customer data
// 	Router.GET("/admin/get-customer-data", controllers.SendAllCustomerData)
// 	// Edit data
// 	Router.PUT("/admin/edit-data", controllers.EditData)
// 	// Reset PIN
// 	Router.PUT("/admin/reset-pin", controllers.ResetPIN)
// 	// Delete account
// 	Router.DELETE("admin/delete-account", controllers.DeleteAccount)
// }
