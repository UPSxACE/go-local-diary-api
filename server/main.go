// @title Swagger Example API
// @version 1.0
// @description Is this even working?
// @host localhost:1323

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/UPSxACE/go-local-diary-api/server/controllers"
	"github.com/UPSxACE/go-local-diary-api/server/views"
	"github.com/boltdb/bolt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/UPSxACE/go-local-diary-api/server/docs"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

func main() {
	// Define a flag named "swag" with a default value of false
	swagFlag := flag.Bool("swag", false, "Run swagger route")
	storybookFlag := flag.Bool("storybook", false, "Run storybook preview route")
	flag.Parse()


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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323", "http://localhost:6006"},
		AllowHeaders: []string{"*"},
		// echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept
	  }))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Serve static files
	e.Static("/public", "/server/public")

	e.Renderer = t

	// Routes
	controllers.SetIndexRoutes(e)

	if(*swagFlag){
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	} 
	if(*storybookFlag){
		controllers.SetStorybookPreviewRoutes(e)
	} 

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}