package main

import (
	"context"
	"errors"
	"fmt"
	"gym-tracker/app/series"
	"gym-tracker/app/set"
	config "gym-tracker/infra"
	"gym-tracker/infra/database/cache"
	"gym-tracker/infra/database/cache/builtin"
	"gym-tracker/infra/database/postgresql"
	"log"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func server() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
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
	seriesHandler := series.NewSeriesHandler(seriesService)
	setsHandler := set.NewSetHandler(setService)

	router := gin.Default()
	seriesGroup := router.Group("/series")
	{
		seriesGroup.GET("/", seriesHandler.GetAllSeries)
		seriesGroup.POST("/", seriesHandler.CreateSeries)
		seriesGroup.POST("/finalize", seriesHandler.FinalizeSerie)

	}
	setsGroup := router.Group("/sets")
	{
		setsGroup.GET("/:serieID", setsHandler.GetAllSet)
		setsGroup.POST("/", setsHandler.CreateSet)
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%v", cfg.Port),
		Handler: router,
	}

	go func() {
		log.Printf("Starting server on http://%s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %s", err)
		}

		log.Println("Server stopped gracefully")
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
	srv.Shutdown(ctx)
	stop()
}
