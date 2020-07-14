package config

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	var (
		cfg *Config
	)

	BeforeEach(func() {
		cfg = New()
	})

	Describe("New", func() {
		Context("when instantiating a new config", func() {
			It("should return new config", func() {
				Expect(cfg).ToNot(BeNil())
			})
		})
	})

	Describe("LoadEnvVars", func() {
		Context("when TLS is enabled but TLS vars are empty", func() {
			It("should return an error", func() {
				os.Setenv("ETCD_USE_TLS", "true")

				err := cfg.LoadEnvVars()

				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("must be set when ETCD_USE_TLS set to true"))
			})
		})
	})
})
