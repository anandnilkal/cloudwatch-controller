module github.com/anandnilkal/cloudwatch-controller

go 1.16

require (
	github.com/aws/aws-sdk-go-v2/config v1.13.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.14.0
	github.com/juju/errors v0.0.0-20210818161939-5560c4c073ff // indirect
	github.com/onsi/ginkgo v1.16.4
	github.com/onsi/gomega v1.15.0
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	sigs.k8s.io/controller-runtime v0.10.0
)
