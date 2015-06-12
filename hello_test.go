package main

import (
	"testing"
//	"net/http"
//	"net/http/httptest"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIntegerManipulation(t *testing.T) {
	t.Parallel()

//	var request *http.Request
//	var recorder *httptest.ResponseRecorder
//
//	Convey("GET /ping", t, func() {
//		recorder = httptest.NewRecorder()
//
//		Convey("should get pong", func() {
//			request, _ = http.NewRequest("GET", "/ping", nil)
//
//		})
//
//	})

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
