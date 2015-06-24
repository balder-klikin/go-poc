package app_test

import (
	. "github.com/balder-klikin/go-poc/app"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Given a Ping", func() {
	var ping Ping

	Context("When all fields are OK", func() {
		BeforeEach(func() {
			ping = Ping{"pong"}
		})

		It("Then it should be valid", func() {
			Expect(ping.Valid()).To(BeTrue())
		})
	})

	Context("When the value is blank", func() {
		BeforeEach(func() {
			ping = Ping{""}
		})

		It("Then it should not be valid", func() {
			Expect(ping.Valid()).To(BeFalse())
		})
	})
})
