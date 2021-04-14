module github.com/cncodetube/course-go-pkg-module

go 1.16

require (
	github.com/cncodetube/course-go-pkg-module-exclude v1.0.0
	github.com/google/uuid v1.1.5 // indirect
	golang.org/x/text v0.3.6
)

replace golang.org/x/text v0.3.6 => github.com/golang/text v0.3.6

exclude github.com/google/uuid v1.2.0
