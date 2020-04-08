package integration_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSmoke(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(5 * time.Minute)
	RunSpecs(t, "Smoke Tests Suite")
}
