package config

import "os"

var ServerPort = ":" + os.Getenv("PORT")

var DataBaseUri = os.Getenv("MONGODB_URI")
