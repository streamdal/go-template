module github.com/batchcorp/go-template

go 1.13

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/DataDog/datadog-go v4.0.1+incompatible // indirect
	github.com/InVisionApp/go-health v2.1.0+incompatible
	github.com/InVisionApp/go-logger v1.0.1
	github.com/batchcorp/grpc-collector v0.0.0-20201114005607-9551461241df // indirect
	github.com/batchcorp/rabbit v0.1.9
	github.com/batchcorp/schemas v0.2.115
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/jackc/pgx/v4 v4.9.2
	github.com/jmoiron/sqlx v1.2.0
	github.com/joho/godotenv v1.3.0
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/maxbrunsfeld/counterfeiter/v6 v6.2.2 // indirect
	github.com/onsi/ginkgo v1.14.1
	github.com/onsi/gomega v1.10.2
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/relistan/go-director v0.0.0-20200406104025-dbbf5d95248d
	github.com/satori/go.uuid v1.2.0
	github.com/segmentio/kafka-go v0.3.7
	github.com/sirupsen/logrus v1.7.0
	github.com/streadway/amqp v1.0.0
	go.etcd.io/etcd v3.3.22+incompatible
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0 // indirect
	gopkg.in/DataDog/dd-trace-go.v1 v1.27.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)
