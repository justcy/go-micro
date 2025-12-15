package template

var (
	Module = `module {{.Dir}}

go 1.18

require (
	go-micro.kanter.cn/v1 latest
	github.com/golang/protobuf latest
	google.golang.org/protobuf latest
)
`
)
