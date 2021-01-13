package middleware

import (
	"github.com/labstack/echo/v4/middleware"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

// var DefaultLoggerConfig = LoggerConfig{
// 	Skipper: DefaultSkipper,
// 	Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
// 		`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
// 		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
// 		`"bytes_out":${bytes_out}}` + "\n",
// 	Output: os.Stdout,
// }
