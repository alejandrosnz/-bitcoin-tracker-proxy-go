package controllers

import (
	"fmt"
	"io"
	"strconv"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TickerResponse struct {
	CurrentPrice float64 `json:"currentPrice"`
}

type ClosingPriceResponse struct {
	ClosingPrice float64 `json:"closingPrice"`
}

func GetCurrentPriceBySymbol(c *gin.Context) {
	/**
	 * Returns current price of a given symbol
	 *  {
	 * 		"currentPrice":60950.01
	 *  }
	 */

	symbol := c.Param("symbol")
	logrus.Infof("Received request to GET currentPrice for symbol: %s", symbol)

	resp, err := http.Get(fmt.Sprintf("https://api.binance.com/api/v3/ticker/24hr?symbol=%sUSDT", symbol))
	if err != nil {
		logrus.Errorf("Error getting ticker price: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting ticker price",
		})
		return
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Error getting ticker price: %d - %s", resp.StatusCode, resp.Status)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting ticker price",
		})
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	logrus.Infof("bodyBytes: %s", bodyBytes)
	if err != nil {
		logrus.Errorf("Error reading server response body: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading server response body",
		})
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logrus.Errorf("Error parsing ticker response: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error parsing ticker response",
		})
		return
	}

	currentPrice, err := strconv.ParseFloat(data["lastPrice"].(string), 64)
	if err != nil {
		logrus.Errorf("Error reading symbol price: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading symbol price",
		})
		return
	}

	c.JSON(http.StatusOK, &TickerResponse{
		CurrentPrice: currentPrice,
	})

	logrus.Infof("Returned currentPrice for symbol: %s: %s", symbol, currentPrice)
}

func GetClosingPriceBySymbol(c *gin.Context) {
	/**
	 * Returns closing price of a given symbol
	 *  {
	 * 		"closingPrice":60950.01
	 *  }
	 */

	symbol := c.Param("symbol")
	logrus.Infof("Received request to GET closingPrice for symbol: %s", symbol)

	resp, err := http.Get(fmt.Sprintf("https://min-api.cryptocompare.com/data/pricemultifull?fsyms=%s&tsyms=USD", symbol))
	if err != nil {
		logrus.Errorf("Error getting closing price: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting closing price",
		})
		return
	}

	if resp.StatusCode != http.StatusOK {
		logrus.Errorf("Error getting closing price: %d - %s", resp.StatusCode, resp.Status)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error getting closing price",
		})
		return
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	logrus.Infof("bodyBytes: %s", bodyBytes)
	if err != nil {
		logrus.Errorf("Error reading server response body: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error reading server response body",
		})
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		logrus.Errorf("Error parsing closing price response: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error parsing closing price response",
		})
		return
	}

	closingPrice := data["RAW"].(map[string]interface{})[symbol].(map[string]interface{})["USD"].(map[string]interface{})["OPENDAY"].(float64)
	
	c.JSON(http.StatusOK, &ClosingPriceResponse{
		ClosingPrice: closingPrice,
	})

	logrus.Infof("Returned closingPrice for symbol: %s: %s", symbol, closingPrice)
}
