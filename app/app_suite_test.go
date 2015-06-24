package app_test

import (
	"testing"

	. "github.com/balder-klikin/go-poc/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	dbSession *DbSession
	App       *Server
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "go-poc/app Suite")
}

var _ = BeforeSuite(func() {
	dbSession = NewDbSession("go-poc-e2e")
	App = NewServer(dbSession)
})
