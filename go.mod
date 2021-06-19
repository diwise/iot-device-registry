module github.com/diwise/iot-device-registry

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/diwise/api-temperature v0.0.0-20210519153202-9fa4ca24d641
	github.com/diwise/messaging-golang v0.0.0-20210519125901-747dbe4d4b42
	github.com/diwise/ngsi-ld-golang v0.0.0-20210619150605-2c79d62512e8
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/rs/cors v1.7.0
	github.com/sirupsen/logrus v1.8.1
	github.com/vektah/gqlparser/v2 v2.1.0
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	gorm.io/driver/postgres v1.0.8
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.9
)
