module L0

go 1.22.1

replace L0/dbpart => ./dbpart

require (
	L0/dbpart v0.0.0-00010101000000-000000000000
	L0/model v0.0.0-00010101000000-000000000000
	github.com/nats-io/stan.go v0.10.4
)

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/nats-io/nats.go v1.22.1 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.5.0 // indirect
)

replace L0/model => ./model

replace L0/common => ./common
