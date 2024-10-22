package configs

import (
	. "api/pkg/configs"
	_ "api/pkg/dotenv"
)

var X_RAPIDAPI_KEY = Get("X_RAPIDAPI_KEY", "")
var X_RAPIDAPI_HOST = Get("X_RAPIDAPI_HOST", "")
