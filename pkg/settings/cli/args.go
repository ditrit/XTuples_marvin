package cli

import (
	"flag"
	"os"
)

const (
	// Key for the environment variable. If it exists, the value of it is used as the path to the config file.
	ENV_CONF_PATH_KEY = "MY_PROJECT_CONF_PATH"
	// Path to the config file if the environment variable is not set or the cli args are not provided.
	DEFAULT_CONFIG_PATH = "./conf/.env.dev"
)

// Args is a struct that holds the command line arguments
// that might be passed to the application
type Args struct {
	Conf string // path to config file
}

// NewArgs returns the command line arguments. If no arguments are passed it
// will return the default values. If arguments to the function are passed,
// it will return those arguments and ignore the command line arguments.
func NewArgs(cliArgs ...Args) Args {
	if len(cliArgs) == 1 {
		return cliArgs[0]
	}

	envPath := setEnvPath()
	if envPath != "" {
		return Args{Conf: envPath}
	}

	var args Args
	flag.StringVar(&args.Conf, "conf", DEFAULT_CONFIG_PATH, "path to the config file")
	flag.Parse()
	return args
}

// setEnvPath returns the path to the config file if the environment variable exits.
// If the environment variable is not set, it will return an empty string.
func setEnvPath() string {
	envPath, ok := os.LookupEnv(ENV_CONF_PATH_KEY)
	if ok && envPath != "" {
		return envPath
	}
	return ""
}
