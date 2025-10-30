package handlers

import (
	"cashierease/internal/models"
	"cashierease/internal/repositories"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateOrderInput struct {
	OrderItems []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"order_items"`
	Coupon struct {
		CouponID uint `json:"couponId"`
	} `json:"coupon"`
	PaymentMethod string `json:"payment_method"`
}

func CreateOrder(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var totalPrice float64 = 0
	var orderItemsDetail []models.OrderItem
	const tax float64 = 0.1

	for _, item := range input.OrderItems {
		produk, err := repositories.GetProdukById(item.ProductID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Produk not found: " + produk.NamaProduk})
			return
		}

		totalPrice += produk.HargaProduk * float64(item.Quantity)

		orderItemsDetail = append(orderItemsDetail, models.OrderItem{
			ProductID:   item.ProductID,
			ProductName: produk.NamaProduk,
			Quantity:    item.Quantity,
		})
	}

	var couponDetails models.CouponDetails
	var totalPriceDiscount float64 = 0
	var discountAmount float64 = 0

	if input.Coupon.CouponID != 0 {
		coupon, err := repositories.GetCouponByID(input.Coupon.CouponID)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Coupon not found"})
			return
		}

		couponDetails = models.CouponDetails{
			CouponID:      coupon.ID,
			KodeCoupon:    coupon.KodeCoupon,
			BesarDiscount: coupon.BesarDiscount,
		}
		discountAmount = coupon.BesarDiscount
		totalPriceDiscount = totalPrice - (totalPrice * discountAmount)
	}

	var finalPriceWithTax float64
	if totalPriceDiscount > 0 {
		finalPriceWithTax = totalPriceDiscount + (totalPriceDiscount * tax)
	} else {
		finalPriceWithTax = totalPrice + (totalPrice * tax)
	}

	order := models.Order{
		OrderDate:            time.Now(),
		TotalPrice:           totalPrice,
		TotalPriceWithDiscount: totalPriceDiscount,
		TotalPriceWithTax:    finalPriceWithTax,
		PaymentMethod:        input.PaymentMethod,
		Tax:                  tax,
		DiscountAmount:       discountAmount,
		Coupon:               couponDetails,
		OrderItems:           orderItemsDetail,
	}

	if err := repositories.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetAllOrders(c *gin.Context) {
	orders, err := repositories.GetAllOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func filterOrdersByDuration(orders []models.Order, duration time.Duration) []models.Order {
	filtered := []models.Order{}
	timeLimit := time.Now().Add(-duration)
	for _, order := range orders {
		if order.CreatedAt.After(timeLimit) {
			filtered = append(filtered, order)
		}
	}
	return filtered
}

func GetAllStatistics(c *gin.Context) {
	now := time.Now()
	last30Days := now.AddDate(0, 0, -30)
	filteredOrders, err := repositories.GetOrdersByDateRange(last30Days, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	var totalIncome float64 = 0
	var totalQuantity int = 0
	var totalCouponUse int = 0

	for _, item := range filteredOrders {
		totalIncome += item.TotalPriceWithTax

		for _, orderItem := range item.OrderItems {
			totalQuantity += orderItem.Quantity
		}

		if item.Coupon.CouponID != 0 {
			totalCouponUse++
		}
	}

	ordersOneHour := filterOrdersByDuration(filteredOrders, 1*time.Hour)
	ordersSixHour := filterOrdersByDuration(filteredOrders, 6*time.Hour)
	ordersTwelveHour := filterOrdersByDuration(filteredOrders, 12*time.Hour)
	ordersOneDay := filterOrdersByDuration(filteredOrders, 24*time.Hour)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"totalIncome":      totalIncome,
			"totalQuantity":    totalQuantity,
			"totalOrders":      len(filteredOrders),
			"totalCouponUse":   totalCouponUse,
			"ordersOneHour":    ordersOneHour,
			"ordersSixHour":    ordersSixHour,
			"ordersTwelveHour": ordersTwelveHour,
			"ordersOneDay":     ordersOneDay,
			"ordersOneMonth":   filteredOrders,
		},
	})
}

type WeeklyIncome struct {
	Week   string  `json:"week"`
	Income float64 `json:"income"`
}

func generateWeeklyChartData(orders []models.Order, weeksRange int) []WeeklyIncome {
	now := time.Now()
	chartData := []WeeklyIncome{}

	for week := 0; week < weeksRange; week++ {
		startOfWeek := now.AddDate(0, 0, -(week+1)*7)
		endOfWeek := now.AddDate(0, 0, -week*7)

		formattedWeek := "Minggu ke-" + strconv.Itoa(weeksRange-week)
		var totalIncome float64 = 0

		for _, order := range orders {
			if order.OrderDate.After(startOfWeek) && order.OrderDate.Before(endOfWeek) {
				totalIncome += order.TotalPriceWithTax
			}
		}

		chartData = append(chartData, WeeklyIncome{
			Week:   formattedWeek,
			Income: totalIncome,
		})
	}

	for i, j := 0, len(chartData)-1; i < j; i, j = i+1, j-1 {
		chartData[i], chartData[j] = chartData[j], chartData[i]
	}
	return chartData
}

func GetAllPendapatan(c *gin.Context) {
	now := time.Now()
	lastThreeMonths := now.AddDate(0, 0, -93)
	orders, err := repositories.GetOrdersByDateRange(lastThreeMonths, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	weeksInMonth := 5
	weeksInTwoMonths := 9
	weeksInThreeMonths := 14

	weeklyDataOneMonth := generateWeeklyChartData(orders, weeksInMonth)
	weeklyDataTwoMonth := generateWeeklyChartData(orders, weeksInTwoMonths)
	weeklyDataThreeMonth := generateWeeklyChartData(orders, weeksInThreeMonths)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"oneMonth":   weeklyDataOneMonth,
			"twoMonth":   weeklyDataTwoMonth,
			"threeMonth": weeklyDataThreeMonth,
		},
	})
}

type DailyOrderCount struct {
	Date   string `json:"date"`
	Orders int    `json:"orders"`
}

func generateOrderCountData(orders []models.Order, daysRange int) []DailyOrderCount {
	now := time.Now()
	chartData := []DailyOrderCount{}
	
	ordersByDate := make(map[string]int)
	for _, order := range orders {
		formattedDate := order.OrderDate.Format("2 January")
		ordersByDate[formattedDate]++
	}

	for i := 0; i < daysRange; i++ {
		date := now.AddDate(0, 0, -i)
		formattedDate := date.Format("2 January")
		
		count, exists := ordersByDate[formattedDate]
		if !exists {
			count = 0
		}
		
		chartData = append(chartData, DailyOrderCount{
			Date:   formattedDate,
			Orders: count,
		})
	}
	
	for i, j := 0, len(chartData)-1; i < j; i, j = i+1, j-1 {
		chartData[i], chartData[j] = chartData[j], chartData[i]
	}
	return chartData
}

func GetAllPelanggan(c *gin.Context) {
	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)
	orders, err := repositories.GetOrdersByDateRange(lastYear, now)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	last7DaysOrders := filterOrdersByDuration(orders, 7*24*time.Hour)
	last31DaysOrders := filterOrdersByDuration(orders, 31*24*time.Hour)
	
	chartData7Days := generateOrderCountData(last7DaysOrders, 7)
	chartData31Days := generateOrderCountData(last31DaysOrders, 31)
	chartData1Year := generateOrderCountData(orders, 365)

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"oneweek":    chartData7Days,
			"onemonth":   chartData31Days,
			"threemonth": chartData1Year,
		},
	})
}

type PopularMenu struct {
	ProductName string `json:"product_name"`
	Count       int    `json:"count"`
}

func GetPopularMenu(c *gin.Context) {
	monParam := c.Param("mon")
	monthOffset, _ := strconv.Atoi(monParam)

	now := time.Now()
	currentYear := now.Year()
	currentMonth := now.Month()

	var firstDate time.Time
	var lastDate time.Time

	if monthOffset == 0 {
		firstDate = time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, now.Location())
		lastDate = now
	} else {
		targetMonth := currentMonth - time.Month(monthOffset)
		firstDate = time.Date(currentYear, targetMonth, 1, 0, 0, 0, 0, now.Location())
		lastDate = time.Date(currentYear, targetMonth+1, 0, 23, 59, 59, 0, now.Location())
	}
	
	orders, err := repositories.GetOrdersByDateRange(firstDate, lastDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get orders"})
		return
	}

	productCount := make(map[string]int)
	for _, order := range orders {
		for _, item := range order.OrderItems {
			productCount[item.ProductName] += item.Quantity
		}
	}

	popularList := []PopularMenu{}
	for name, count := range productCount {
		popularList = append(popularList, PopularMenu{ProductName: name, Count: count})
	}

	sort.Slice(popularList, func(i, j int) bool {
		return popularList[i].Count > popularList[j].Count
	})

	topThree := popularList
	if len(popularList) > 3 {
		topThree = popularList[:3]
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   topThree,
	})
}