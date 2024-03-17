package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Futturi/vktest/internal/handler"
	"github.com/Futturi/vktest/internal/repository"
	"github.com/Futturi/vktest/internal/server"
	"github.com/Futturi/vktest/internal/service"
	"github.com/Futturi/vktest/pkg"
	"github.com/spf13/viper"
)

// @title Cinema App API
// @version 1.0
// @description API Server 4 Cinema Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logg := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logg)
	err := InitConfig()
	if err != nil {
		slog.Error("error with config", slog.Any("error", err))
	}
	pcfg := pkg.PConfig{
		Hostname: viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		NameDB:   viper.GetString("db.name"),
		SSLmode:  viper.GetString("db.sslmode"),
	}

	db, err := pkg.InitPostgres(pcfg)
	if err != nil {
		slog.Error("error with db", slog.Any("error", err))
	}

	repo := repository.NewRepostitory(db)
	service := service.NewService(repo)
	han := handler.NewHandl(service)
	server := new(server.Server)
	go func() {
		if err = server.InitServer(viper.GetString("port"), han.NewHan()); err != nil {
			slog.Error("error with server", slog.Any("error", err))
		}
	}()
	logg.Info("statring app in port: ", slog.String("port", viper.GetString("port")))
	if err := pkg.Migrat(viper.GetString("db.host")); err != nil {
		slog.Error("error with migratedb", slog.Any("error", err))
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logg.Info("shutdown server", slog.String("port", viper.GetString("port")))
	if err = server.ShutDown(context.Background()); err != nil {
		logg.Error("error with shutdown server", slog.String("port", viper.GetString("port")), slog.Any("error", err))
		os.Exit(1)
	}
	if err = pkg.ShutDown(db); err != nil {
		logg.Error("error with close db", slog.Any("error", err))
		os.Exit(1)
	}
}
func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}
