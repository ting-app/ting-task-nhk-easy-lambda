package ting

import (
	"os"
	"strconv"
)

type DbConfig struct {
	DriverName string
	UserName   string
	Password   string
	Host       string
	Port       int
}

func ParseDbConfig() (DbConfig, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		return DbConfig{}, err
	}

	return DbConfig{
		DriverName: os.Getenv("DB_DRIVER_NAME"),
		UserName:   os.Getenv("DB_USER_NAME"),
		Password:   os.Getenv("DB_PASSWORD"),
		Host:       os.Getenv("DB_HOST"),
		Port:       port,
	}, nil
}
