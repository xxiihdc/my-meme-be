package env

import "os"

var (
	ClientSecretFile = os.Getenv("CLIENT_SECRET_FILE")
)
