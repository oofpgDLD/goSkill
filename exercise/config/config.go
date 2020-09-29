package config

type Config struct {
	HTTPServer HTTPServer
}

type HTTPServer struct {
	Env  string
	Addr string
}
