package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanakize/skill-api-kafka/api/database"
	"github.com/thanakize/skill-api-kafka/api/router"
	"github.com/thanakize/skill-api-kafka/api/skill"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	database := database.InitDatabase()
	defer database.CloseDatabase()

	skillRepo := skill.InitSkillRepo(database.DB)
	skillProducer := skill.CreateProducer("skill")
	defer skillProducer.CloseProducer()
	skillHandler := skill.NewHandler(skillRepo, skillProducer)

	r := gin.Default()
	router.InitRoute(r, skillHandler)

	srv := http.Server{
		Addr:    ":" + os.Getenv("PRODUCER_PORT"),
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down...")
		ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
		defer cancle()

		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println(err)
			}
		}

	}()
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
