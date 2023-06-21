package controllers

import (
	"BNMO/models"
	"BNMO/redis"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


var (
	redisClient redis.RedisCache = redis.NewRedisCache("bnmo-redis:6379", 0, time.Hour * 24)
)

func requestSymbolsFromAPI() map[string]string {
	url := "https://api.apilayer.com/exchangerates_data/symbols"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("apikey", "5dFvi1dREEw8zy6ZR4CkPoIahE8oT9kN")

	if err != nil {
		panic(err)
	}

	// Start the request
	result, _ := client.Do(request)
	if result.Body != nil {
		defer result.Body.Close()
	}

	// Read the response
	body, _ := ioutil.ReadAll(result.Body)

	// Unmarshal the response to struct
	var output models.ExchangeSymbols
	json.Unmarshal(body, &output)

	redisClient.SetCache("symbols", body)
	return output.Symbols
}

func requestRatesFromAPI(requestedKey string) float64 {
	url := "https://api.apilayer.com/exchangerates_data/latest?symbols=&base=IDR"

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("apikey", "5dFvi1dREEw8zy6ZR4CkPoIahE8oT9kN")

	if err != nil {
		panic(err)
	}

	// Start the request
	result, _ := client.Do(request)
	if result.Body != nil {
		defer result.Body.Close()
	}

	// Read the response
	body, _ := ioutil.ReadAll(result.Body)

	// Unmarshal the response to struct
	var output models.ExchangeRates
	json.Unmarshal(body, &output)

	redisClient.SetCache("rates", body)
	return output.Rates[requestedKey]
}

func getSymbolsFromRedis() (string, map[string]string) {
	// Check cache availability
	cacheEntry := redisClient.GetCache("symbols", "")
	
	// Cache hit events
	if cacheEntry != -1 && cacheEntry != -2 {
		output := cacheEntry.(map[string]string)
		fmt.Println("Value symbols found within redis")
		return "Cache", output
	}

	// Cache miss events
	fmt.Println("Pulling symbols value from API")
	apiEntry := requestSymbolsFromAPI()
	
	return "API", apiEntry
}

func getRatesFromRedis(requestedKey string) (string, float64) {
	// Check cache availability
	cacheEntry := redisClient.GetCache("rates", requestedKey)
	
	// Cache hit events
	if cacheEntry != -1 && cacheEntry != -2 {
		fmt.Println("Value symbols found within redis")
		return "Cache", cacheEntry.(float64)
	}

	// Cache miss events
	fmt.Println("Pulling symbols value from API")
	apiEntry := requestRatesFromAPI(requestedKey)
	
	return "API", apiEntry
}


func GetSymbols(c * gin.Context) {
	_, symbols := getSymbolsFromRedis()
	c.JSON(http.StatusOK, gin.H{
		"symbols": symbols,
	})
}
