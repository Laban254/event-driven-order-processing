package config

import (
    "os"
    "strconv"
)


type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    DBPort     int 
}


func LoadConfig() Config {
    dbPort, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
    if err != nil {
        dbPort = 5432 
    }
    return Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "postgres"),
        DBName:     getEnv("DB_NAME", "order_service"),
        DBPort:     dbPort, 
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
