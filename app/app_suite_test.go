package app_test

import (
	"testing"

	. "github.com/balder-klikin/go-poc/app"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

var (
	dbSession *DbSession
	App       *Server
)

func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	rs := []Reporter{reporters.NewJUnitReporter("../report/gopoc-app-junit.xml")}
	RunSpecsWithDefaultAndCustomReporters(t, "go-poc/app Suite", rs)
}

var _ = BeforeSuite(func() {
	dbSession = NewDbSession("go-poc-e2e")
	App = NewServer(dbSession)
})
