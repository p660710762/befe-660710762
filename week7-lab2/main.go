package main

import (
	"fmt"
	"os"
)

func getEnv(key, dafaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return dafaultValue
}

func main() {
	//value := os.Getenv("DB_HOST")
	host := getEnv("DB_HOST", "")
	name := getEnv("DB_NAME", "")
	user := getEnv("DB_USER", "")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "")

	conSt := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s" , host, port , user, password, name)
	fmt.Println(conSt)

}

