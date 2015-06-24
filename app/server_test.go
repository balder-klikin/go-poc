package app_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/balder-klikin/go-poc/app"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server E2E", func() {
	var (
		req  *http.Request
		resp *httptest.ResponseRecorder
	)
	BeforeEach(func() {
		resp = httptest.NewRecorder()
	})

	Context("GET /ping", func() {
		var ping app.Ping

		BeforeEach(func() {
			req, _ = http.NewRequest("GET", "/ping", nil)

			App.ServeHTTP(resp, req)
			json.Unmarshal(resp.Body.Bytes(), &ping)
		})

		It("should get 200", func() {
			Expect(resp.Code).To(Equal(200))
		})
		It("should get pong", func() {
			Expect(ping.Value).To(Equal("PONG!"))
		})
	})

	Context("GET /check", func() {
		BeforeEach(func() {
			req, _ = http.NewRequest("GET", "/check", nil)

			App.ServeHTTP(resp, req)
		})

		It("should get 200", func() {
			Expect(resp.Code).To(Equal(200))
		})
		It("should get pong", func() {
			Expect(resp.Body.String()).To(Equal("OK"))
		})
	})
})
