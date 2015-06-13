package gopoc

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestPingValidation(t *testing.T) {
	t.Parallel()

	Convey("Given a Ping", t, func() {
		var ping Ping

		Convey("When all fields are OK", func() {
			ping = Ping{"pong"}

			Convey("Then it should be valid", func() {
				So(ping.valid(), ShouldBeTrue)
			})
		})

		Convey("When the value is blank", func() {
			ping = Ping{""}

			Convey("Then it should not be valid", func() {
				So(ping.valid(), ShouldBeFalse)
			})
		})
	})
}
