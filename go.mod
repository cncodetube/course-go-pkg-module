module github.com/cncodetube/course-go-pkg-module

go 1.16

require (
	github.com/cncodetube/faker v1.0.1
	github.com/google/uuid v1.2.0
	golang.org/x/text v0.3.6
)

replace (
	github.com/cncodetube/faker v1.0.1 => ../faker
	golang.org/x/text v0.3.6 => github.com/golang/text v0.3.6
)

//禁止被依赖示例见：https://github.com/kubernetes/kubernetes/blob/master/go.mod
