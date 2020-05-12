package main

import (
	"fmt"
	"log"
	"os"

	_podcastHttpHandler "github.com/andrewjapar/andrew-cv-app/app/podcast/handler"
	_podcastRepo "github.com/andrewjapar/andrew-cv-app/app/podcast/repository"
	_weddingHttpHandler "github.com/andrewjapar/andrew-cv-app/app/wedding/handler"
	_weddingRepo "github.com/andrewjapar/andrew-cv-app/app/wedding/repository"
	"github.com/andrewjapar/andrew-cv-app/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		fmt.Print(envErr)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	defer conn.Close()

	conn.AutoMigrate(&domain.Profile{}, &domain.Podcast{}, &domain.Wedding{}, &domain.User{}, &domain.WeddingOrganizer{})

	e := echo.New()

	podcastRepo := _podcastRepo.NewPodcastRepository(conn)
	_podcastHttpHandler.NewPodcastHandler(e, podcastRepo)

	weddingRepo := _weddingRepo.NewWeddingRepository(conn)
	_weddingHttpHandler.NewWeddingHandler(e, weddingRepo)

	log.Fatal(e.Start(os.Getenv("server_address")))
}
