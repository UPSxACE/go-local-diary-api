package main

import (
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/UPSxACE/go-local-diary-api/server/controllers"
	"github.com/UPSxACE/go-local-diary-api/server/views"
	"github.com/boltdb/bolt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		fmt.Println("ServerConfig: ", tx.Bucket([]byte("ServerConfig")))
		serverConfigBucket, err := tx.CreateBucketIfNotExists([]byte("ServerConfig"))
		if(err != nil){
			return err
		}

		initialized := serverConfigBucket.Get([]byte("initialized"))
		if(initialized == nil){
			fmt.Println("First time initializing the server")
			serverConfigBucket.Put([]byte("initialized"), []byte("true"))
			fmt.Println("Now initialized!")
		}

		initialized = serverConfigBucket.Get([]byte("initialized"))
		fmt.Printf("Server initialized: %s\n", initialized)

		//if err := initBucket.Put(true)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	// Pre-compile templates in views subdirectories, and subdirectories of those subdirectories
	tBuilder :=  template.Must(template.ParseGlob("server/views/*/*.html"));
	tBuilder = template.Must(tBuilder.ParseGlob("server/views/*/*/*.html"))
	t := &views.Template{
		Templates: tBuilder,
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/public", "/server/public")

	e.Renderer = t

	// Routes
	controllers.SetIndexRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
