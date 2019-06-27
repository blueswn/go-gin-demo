package setting

import "time"

type Server struct {
	DebugMode    string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
