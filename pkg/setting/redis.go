package setting

import "time"

/*
  |--------------------------------------------------------------------------
  | Redis Databases
  |--------------------------------------------------------------------------
  |
  | Redis is an open source, fast, and advanced key-value store that also
  | provides a richer set of commands than a typical key-value systems
  | such as APC or Memcached. Laravel makes it easy to dig right in.
  |
*/

type Redis struct {
	Host string
	Password string
	Port string
	Database int
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}
