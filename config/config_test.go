package config

import (
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
		Context("when no ENV vars are properly set", func() {
			It("should return an error about the unset vars", func() {
				err := cfg.LoadEnvVars()

				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(ContainSubstring("Unable to fetch env vars: required key LISTEN_ADDRESS missing value"))
			})
		})
	})
})
