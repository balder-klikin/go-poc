package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoPocE2E(t *testing.T) {
	t.Parallel()

	var req *http.Request
	var resp *httptest.ResponseRecorder
	var DbSession *DbSession
	var server *Server

	Convey("GET /ping", t, func() {
		resp = httptest.NewRecorder()
		DbSession = NewDbSession("go-poc-e2e")
		server = NewServer(DbSession)
		req, _ = http.NewRequest("GET", "/ping", nil)

		server.ServeHTTP(resp, req)
		var ping Ping
		json.Unmarshal(resp.Body.Bytes(), &ping)

		Convey("should get 200", func() {
			So(resp.Code, ShouldEqual, 200)
		})
		Convey("should get pong", func() {
			So(ping.Value, ShouldEqual, "PONG!")
		})
	})

	Convey("GET /check", t, func() {
		resp = httptest.NewRecorder()
		DbSession = NewDbSession("go-poc-e2e")
		server = NewServer(DbSession)
		req, _ = http.NewRequest("GET", "/check", nil)

		server.ServeHTTP(resp, req)

		Convey("should get 200", func() {
			So(resp.Code, ShouldEqual, 200)
		})
		Convey("should get pong", func() {
			So(resp.Body.String(), ShouldEqual, "OK")
		})
	})

	Convey("Given a starting integer value", t, func() {
		x := 42

		Convey("When incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 43)
			})
			Convey("The value should NOT be what it used to be", func() {
				So(x, ShouldNotEqual, 42)
			})
		})
		Convey("When decremented", func() {
			x--

			Convey("The value should be lesser by one", func() {
				So(x, ShouldEqual, 41)
			})
			Convey("The value should NOT be what it used to be", func() {
				So(x, ShouldNotEqual, 42)
			})
		})
	})
}
