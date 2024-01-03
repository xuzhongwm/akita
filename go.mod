module github.com/sarchlab/akita/v4

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang/mock v1.6.0
	github.com/google/pprof v0.0.0-20231229205709-960ae82b1e42
	github.com/gorilla/mux v1.8.1
	github.com/mattn/go-sqlite3 v1.14.19
	github.com/onsi/ginkgo/v2 v2.13.2
	github.com/onsi/gomega v1.30.0
	github.com/rs/xid v1.5.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/syifan/goseth v0.1.1
	github.com/tebeka/atexit v0.3.0
)

require (
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/kr/pretty v0.1.0 // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/tools v0.16.1 // indirect
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace github.com/syifan/goseth => ../goseth

go 1.20
