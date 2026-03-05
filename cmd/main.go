package main

import (
	"gym-tracker/app/series"
	"gym-tracker/app/set"
	"gym-tracker/app/user"
	config "gym-tracker/infra"
	"gym-tracker/infra/bot"
	"gym-tracker/infra/database/cache"
	"gym-tracker/infra/database/cache/builtin"
	"gym-tracker/infra/database/postgresql"
)

func main() {
	cfg := config.Parse()
	database, errDatabase := postgresql.New(cfg.Postgresql)
	if errDatabase != nil {
		defer errDatabase()
	}
	builtin := builtin.NewBuiltin()
	cache := cache.NewCache(builtin)
	seriesRepository := series.NewSeriesRepository(database)
	seriesService := series.NewService(seriesRepository)
	setRepository := set.NewSetsRepository(database)
	setService := set.NewService(cache, setRepository)
	userRepository := user.NewUserRepository(database)
	userService := user.NewService(userRepository)
	telegram := bot.NewTelegram(cfg.Telegram, seriesService, setService, userService)
	telegram.Chat()
}
