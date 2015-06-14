package gopoc

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGoPocE2E(t *testing.T) {
	t.Parallel()

	var req *http.Request               // the req
	var resp *httptest.ResponseRecorder // resp
	var mgoSession *MgoSession
	var server *gin.Engine              // the gin Engine


	Convey("GET /ping", t, func() {
		resp = httptest.NewRecorder()
		mgoSession = NewMgoSession("go-poc-e2e")
		server = NewServer(mgoSession)
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
