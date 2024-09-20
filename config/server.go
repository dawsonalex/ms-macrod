package config

type Server struct {
	Host string `ini:"host"`
	Port string `ini:"port"`
}
