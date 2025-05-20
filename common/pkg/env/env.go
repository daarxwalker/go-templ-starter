package env

import "os"

func App() string {
	return os.Getenv("APP_ENV")
}

func Name() string {
	return os.Getenv("APP_NAME")
}

func Port() string {
	return os.Getenv("APP_PORT")
}

func Empty() bool {
	return len(App()) == 0
}

func Development() bool {
	return App() == "development"
}

func Production() bool {
	return App() == "production"
}
