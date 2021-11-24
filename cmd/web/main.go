package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/lildannylin/bookings-app-golang/internal/config"
	"github.com/lildannylin/bookings-app-golang/internal/driver"
	"github.com/lildannylin/bookings-app-golang/internal/handlers"
	"github.com/lildannylin/bookings-app-golang/internal/helpers"
	"github.com/lildannylin/bookings-app-golang/internal/models"
	"github.com/lildannylin/bookings-app-golang/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var (
	app        config.AppConfig
	session    *scs.SessionManager
	infoLog    *log.Logger
	errorLog   *log.Logger
	DBhost     string
	DBport     string
	DBuser     string
	DBpassword string
	DBname     string
)

// main is the main function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(app.MailChan)

	fmt.Println("Starting mail listener...")
	listenForMail()

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	//Put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InforLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	DBhost = "localhost"
	DBport = "5432"
	DBuser = "danny"
	DBpassword = "1201"
	DBname = "bookings"

	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL(fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		DBhost, DBport, DBname, DBuser, DBpassword))
	if err != nil {
		log.Fatal("Cannot connect database")
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelper(&app)

	return db, nil
}
