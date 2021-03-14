module github.com/iot-for-tillgenglighet/iot-device-registry

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/iot-for-tillgenglighet/messaging-golang v0.0.0-20201230002037-e79e8e927ae9
	github.com/iot-for-tillgenglighet/ngsi-ld-golang v0.0.0-20210314221235-1ccf2d63e861
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.8.1
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.3
)
