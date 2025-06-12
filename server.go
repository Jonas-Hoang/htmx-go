package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Get("/search", func(c *fiber.Ctx) error {
		ticker := c.Query("ticker")
		results := SearchTicker(ticker)

		return c.Render("results", fiber.Map{
			"Results": results,
		})
	})

	app.Get("/values/:ticker", func(c *fiber.Ctx) error {
		ticker := c.Params("ticker")
		values := GetDailyValues(ticker)

		return c.Render("values", fiber.Map{
			"Ticker": ticker,
			"Values": values,
		})
	})

	app.Get("/values/:ticker/:date", func(c *fiber.Ctx) error {
		ticker := c.Params("ticker")
		date := c.Params("date")
		values := GetDailyValues(ticker)
		dailyValue := GetDailyValueByDate(values, date)
		if dailyValue == nil {
			return c.Status(fiber.StatusNotFound).SendString("Value not found for the specified date")
		}
		return c.Render("daily_value", fiber.Map{
			"Ticker": ticker,
			"Date":   date,
			"Value":  dailyValue,
		})
	})

	app.Listen(":3000")
}
