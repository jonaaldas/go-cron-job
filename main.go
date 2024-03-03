package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gopkg.in/robfig/cron.v2"
)

type RestaurantId struct {
	ID int `db:"id"`
}

func connectToDB() {
	// user := goDotEnvVariable("DB_USER")
	// dataBaseName := goDotEnvVariable("DATA_BASE_NAME")
	// password := goDotEnvVariable("PASSWORD")
	// host := goDotEnvVariable("HOST")
	user := os.Getenv("DB_USER")
	dataBaseName := os.Getenv("DATA_BASE_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable", user, dataBaseName, password, host)
	fmt.Print(database)

	db, err := sqlx.Connect("postgres", database)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	var restaurantId RestaurantId
	err = db.Get(&restaurantId, "SELECT id FROM restaurants LIMIT 1")
	if err != nil {
		log.Fatal("Failed to query restaurant ID:", err)
	}

	fmt.Printf("First Restaurant ID: %d\n", restaurantId.ID)
	fmt.Printf("I finish running")

}

func runCronJobs() {
	fmt.Print("I was running")
	s := cron.New()
	s.AddFunc("* * * * *", func() {
		connectToDB()
	})
	s.Start()
}

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	runCronJobs()
	select {}
}
