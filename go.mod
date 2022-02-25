module github.com/cloudwan/ztna-sdk

go 1.17

require (
	github.com/cloudwan/edgelq-sdk v0.4.17
	github.com/cloudwan/goten-sdk v0.4.18
	github.com/golang/protobuf v1.5.2
	github.com/google/cel-go v0.5.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/iancoleman/strcase v0.2.0
	github.com/spf13/pflag v1.0.5
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

require (
	cloud.google.com/go v0.81.0 // indirect
	github.com/alecthomas/participle v0.5.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.1 // indirect
	github.com/sirupsen/logrus v1.4.2 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/appengine v1.6.7 // indirect
)

replace google.golang.org/protobuf => github.com/cloudwan/goten-protobuf v1.26.0
