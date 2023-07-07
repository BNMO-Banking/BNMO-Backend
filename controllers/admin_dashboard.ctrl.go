package controllers

import (
	"BNMO/database"
	"BNMO/enum"
	"BNMO/models"
	"BNMO/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func GetPendingLists(c *gin.Context) {
	var pendingAccounts []models.PendingAccount
	var pendingRequests []models.PendingRequest
	var totalPendingAccounts int64
	var totalPendingRequests int64

	database.DB.
		Table("customers").
		Select("accounts.first_name, accounts.last_name, customers.created_at").
		Joins("JOIN accounts ON customers.account_id = accounts.id").
		Where("status = ?", enum.ACCOUNT_PENDING).
		Find(&pendingAccounts).
		Limit(5)
	database.DB.Table("customers").Where("status = ?", enum.ACCOUNT_PENDING).Count(&totalPendingAccounts)

	database.DB.
		Table("requests").
		Select("accounts.first_name, accounts.last_name, requests.request_type, requests.created_at").
		Joins("JOIN customers ON requests.customer_id = customers.id").
		Joins("JOIN accounts ON customers.account_id = accounts.id").
		Where("requests.status = ?", enum.REQUEST_PENDING).
		Find(&pendingRequests).
		Limit(5)
	database.DB.Table("requests").Where("status = ?", enum.REQUEST_PENDING).Count(&totalPendingRequests)

	c.JSON(http.StatusOK, gin.H{
		"accounts": gin.H{
			"pending": pendingAccounts,
			"total":   totalPendingAccounts,
		},
		"requests": gin.H{
			"pending": pendingRequests,
			"total":   totalPendingRequests,
		}})
}

func GetNewAccountStatistics(c *gin.Context) {
	year := c.Query("year")

	var totalAccounts int64
	var yearlyAccounts int64
	monthlyAccounts := make([]int64, 12)

	database.DB.Table("customers").Count(&totalAccounts)

	startDate, _ := time.Parse("2006-01-02", fmt.Sprintf("%s-01-01", year))
	endDate := startDate.AddDate(1, 0, 0).Add(-time.Nanosecond)

	database.DB.
		Table("customers").
		Where("created_at >= ? AND created_at <= ?",
			startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Count(&yearlyAccounts)

	startDate, _ = time.Parse("2006-01-02", fmt.Sprintf("%s-01-01", year))

	for i := 0; i < 12; i++ {
		var tempAccount int64
		endDate = startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

		database.DB.
			Table("customers").
			Where("created_at >= ? AND created_at <= ?",
				startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
			Count(&tempAccount)

		monthlyAccounts[i] = tempAccount

		startDate = startDate.AddDate(0, 1, 0)
	}

	response := models.NewAccountStatistics{
		TotalAccounts:   totalAccounts,
		YearlyAccounts:  yearlyAccounts,
		MonthlyAccounts: monthlyAccounts,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetRequestTypeStatistics(c *gin.Context) {
	year := c.Query("year")
	monthQuery := c.DefaultQuery("month", "")

	var entireYear bool
	if len(monthQuery) == 0 {
		entireYear = true
	} else {
		entireYear = false
	}

	var addRequest int64
	var subtractRequest int64

	var startDate time.Time
	var endDate time.Time

	if entireYear {
		startDate, _ = time.Parse("2006-01-02", fmt.Sprintf("%s-01-01", year))
		endDate = startDate.AddDate(1, 0, 0).Add(-time.Nanosecond)
	} else {
		month, err := utils.ParseMonth(monthQuery)
		if err != nil {
			utils.HandleBadRequest(c, err.Error())
			return
		}
		startDate, _ = time.Parse("2006-01-02", fmt.Sprintf("%s-%s-01", year, month))
		endDate = startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)
	}

	database.DB.
		Table("requests").
		Where("request_type = ? AND created_at >= ? AND created_at <= ?",
			enum.ADD, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Count(&addRequest)
	database.DB.
		Table("requests").
		Where("request_type = ? AND created_at >= ? AND created_at <= ?",
			enum.SUBTRACT, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Count(&subtractRequest)

	c.JSON(http.StatusOK, gin.H{
		"add":      addRequest,
		"subtract": subtractRequest,
	})
}

func GetBankStatistics(c *gin.Context) {
	year := c.Query("year")

	var totalExpense decimal.Decimal
	var totalIncome decimal.Decimal

	monthlyExpense := make([]decimal.Decimal, 12)
	monthlyIncome := make([]decimal.Decimal, 12)

	startDate, _ := time.Parse("2006-01-02", fmt.Sprintf("%s-01-01", year))
	endDate := startDate.AddDate(1, 0, 0).Add(-time.Nanosecond)

	database.DB.
		Table("requests").
		Select("SUM(converted_amount)").
		Where("request_type = ? AND updated_at >= ? AND updated_at <= ?",
			enum.ADD, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Scan(&totalExpense)
	database.DB.
		Table("requests").
		Select("SUM(converted_amount)").
		Where("request_type = ? AND updated_at >= ? AND updated_at <= ?",
			enum.SUBTRACT, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Scan(&totalIncome)

	startDate, _ = time.Parse("2006-01-02", fmt.Sprintf("%s-01-01", year))

	for i := 0; i < 12; i++ {
		var tempExpense decimal.Decimal
		var tempIncome decimal.Decimal

		endDate := startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

		database.DB.
			Table("requests").
			Select("SUM(converted_amount)").
			Where("request_type = ? AND updated_at >= ? AND updated_at <= ?",
				enum.ADD, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
			Scan(&tempExpense)
		database.DB.
			Table("requests").
			Select("SUM(converted_amount)").
			Where("request_type = ? AND updated_at >= ? AND updated_at <= ?",
				enum.SUBTRACT, startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
			Scan(&tempIncome)

		monthlyExpense[i] = tempExpense
		monthlyIncome[i] = tempIncome

		startDate = startDate.AddDate(0, 1, 0)
	}

	response := models.BankStatistics{
		TotalExpense:   totalExpense,
		TotalIncome:    totalIncome,
		MonthlyExpense: monthlyExpense,
		MonthlyIncome:  monthlyIncome,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
