package gopoc

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoPocE2E(t *testing.T) {
	t.Parallel()

	var request 	*http.Request
	var recorder 	*httptest.ResponseRecorder
	var server 		*gin.Engine

	Convey("GET /ping", t, func() {
		recorder = httptest.NewRecorder()
		server = NewServer()
		request, _ = http.NewRequest("GET", "/ping", nil)

		server.ServeHTTP(recorder, request)
		var p Ping
		json.Unmarshal(recorder.Body.Bytes(), &p)

		Convey("should get 200", func() {
			So(recorder.Code, ShouldEqual, 200)
		})
		Convey("should get pong", func() {
			So(p.Value, ShouldEqual, "pong")
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
