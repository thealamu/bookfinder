/* Package setup initializes everything necessary to start the server*/
package setup

import (
	"github.com/thealamu/bookfinder/internal/pkg/server"
	"github.com/thealamu/bookfinder/pkg/finder"
)

// ServerEnv creates a new server env and inits it
func ServerEnv(port string) (*server.ServerEnv, error) {
	srvEnv := server.NewServerEnv()
	srvEnv.Port = port

	// Create finders
	//Google Books
	srvEnv.GoogleBooksFinder = finder.NewGoogleBooksFinder()

	return srvEnv, nil
}
