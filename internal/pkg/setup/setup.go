/* Package setup initializes everything necessary to start the server*/
package setup

import (
	"errors"
	"fmt"
	"os"

	"github.com/thealamu/bookfinder/internal/pkg/server"
	"github.com/thealamu/bookfinder/pkg/finder"
)

var NoGRAPIKey = errors.New("GoodReads API key not found in running environment")

// ServerEnv creates a new server env and inits it
func ServerEnv(port string) (*server.ServerEnv, error) {
	srvEnv := server.NewServerEnv()
	srvEnv.Port = port

	// Create finders
	//Google Books
	srvEnv.GoogleBooksFinder = finder.NewGoogleBooksFinder()

	//GoodReads
	apiKey := os.Getenv("GREADS_KEY")
	if apiKey == "" {
		fmt.Println("setup.ServerEnv", NoGRAPIKey)
		return srvEnv, NoGRAPIKey
	}
	srvEnv.GoodReadsFinder = finder.NewGoodReadsFinder(apiKey)

	return srvEnv, nil
}
