package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Futturi/vktest/internal/handler"
	"github.com/Futturi/vktest/internal/repository"
	"github.com/Futturi/vktest/internal/server"
	"github.com/Futturi/vktest/internal/service"
	"github.com/Futturi/vktest/pkg"
	"github.com/spf13/viper"
)

// @title Cinema App API
// @description API Server 4 Cinema Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	err := InitConfig()
	if err != nil {
		fmt.Println(err)
	}

	logg := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logg)
	logg.Info("statring app in port: ", slog.String("port", viper.GetString("port")))
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
		fmt.Println(err)
	}
	repo := repository.NewRepostitory(db)
	service := service.NewService(repo)
	han := handler.NewHandl(service)
	server := new(server.Server)
	if err = server.InitServer(viper.GetString("port"), han.NewHan()); err != nil {
		fmt.Println(err)
	}
}
func InitConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("internal/config")
	return viper.ReadInConfig()
}
