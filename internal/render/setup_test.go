package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ferhatyegin/goBookings/internal/config"
	"github.com/ferhatyegin/goBookings/internal/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M){
	//What am I going to put in the session
	gob.Register(models.Reservation{})

	// Change this to true when in production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session
	app = &testApp

	os.Exit(m.Run())
}

// Implement our own http.writer struct 
type myWriter struct{}

func (tw *myWriter) Header() http.Header {
 	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int){

} 

func (tw *myWriter) Write (b []byte) (int, error) {
	length := len(b)
	return length, nil
}

