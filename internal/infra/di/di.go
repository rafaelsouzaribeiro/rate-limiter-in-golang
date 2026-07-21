package di

import (
	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/infra/database/redis/connection"
	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/infra/database/redis/repository"
	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/infra/web/handler"
	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/infra/web/server"
	"github.com/rafaelsouzaribeiro/rate-limiter-in-golang/internal/usecase"
	"github.com/spf13/viper"
)

func NewDI() {
	con := connection.NewConnection(viper.GetString("HOST_REDIS"), viper.GetString("PASSWORD_REDIS"))
	repo := repository.NewRepository(con)
	usecease := usecase.NewUsecase(repo)
	handler := handler.NewHandler(*usecease)
	server := server.NewServer(viper.GetString("SERVER_PORT"))
	server.RegisterHandler("GET /rate-limiter", handler.RateLimiter)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
