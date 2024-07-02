package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func sendMessageToTelegram(price float64, botToken string, chatID int64) error {
	// Create a new bot with the provided token
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}

	// Formulate the message
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Текущая цена BTC: $%.2f", price))

	// Send the message
	_, err = bot.Send(msg)
	return err
}

type CoinGeckoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func main() {
	// URL for fetching BTC price data
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

	// Execute HTTP request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Некорректный статус ответа:", resp.StatusCode)
		os.Exit(1)
	}

	// Decode JSON response
	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		os.Exit(1)
	}

	// Get current BTC price
	price := result.Bitcoin.USD
	fmt.Printf("Текущая цена BTC: $%.2f\n", price)

	// Replace with your actual bot token and chat ID
	botToken := "7175657723:AAGScSFnJ_AOrCrnFqXSwMKFv9wvpQUz8iI"
	chatID := int64(5114383390)

	// Call function to send message to Telegram
	err = sendMessageToTelegram(price, botToken, chatID)
	if err != nil {
		fmt.Println("Ошибка отправки сообщения в Телеграм:", err)
		os.Exit(1)
	}
}

