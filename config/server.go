package config

import "time"

// Server is the server configuration. Used with "github.com/kelseyhightower/envconfig".
// It needs at least the port.
// Timeout and GracefullStop are set to 5 seconds by default (can be overwritten).
type Server struct {
	Port          string        `required:"true"`
	Timeout       time.Duration `default:"5s"`
	GracefullStop time.Duration `default:"5s"`
}
