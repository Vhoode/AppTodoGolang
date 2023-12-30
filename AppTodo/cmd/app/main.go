package main

import (
	"os"
	reapeatTodo "repeat/Todo"
	"repeat/Todo/pkg/handler"
	"repeat/Todo/pkg/repository"
	"repeat/Todo/pkg/server"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewMySqlDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetUint32("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed initialization db:%s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := server.NewService(repo)
	handlers := handler.NewHandler(services)

	svr := new(reapeatTodo.Server)
	if err := svr.Run(viper.GetString("port"), handlers.InitRouter()); err != nil {
		logrus.Fatalf("error http sever doesent Run:%s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
