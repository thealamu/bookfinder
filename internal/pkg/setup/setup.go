/* Package setup initializes everything necessary to start the server*/
package setup

import (
	"fmt"
	"os"

	"github.com/thealamu/bookfinder/internal/pkg/server"
	"github.com/thealamu/bookfinder/pkg/finder"
)

// ServerEnv creates a new server env and inits it
func ServerEnv(port string) (*server.ServerEnv, error) {
	srvEnv := server.NewServerEnv()
	srvEnv.Port = port

	// Create finders
	//Google Books
	gbKey := os.Getenv("GB_API_KEY")
	gbFinder, err := finder.NewGoogleBooksFinder(gbKey)
	if err != nil {
		fmt.Println("setup.ServerEnv", "Could not create a google books finder")
		return nil, err
	}
	srvEnv.GoogleBooksFinder = gbFinder

	return srvEnv, nil
}
