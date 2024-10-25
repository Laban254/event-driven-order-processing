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
        DBUser:     getEnv("DB_USER", "event"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "event"),
        DBPort:     dbPort, 
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
