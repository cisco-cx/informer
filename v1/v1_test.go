package v1_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cisco-cx/informer"
	"github.com/cisco-cx/informer/v1"
)

// Confirm that v1.InfoService implements the informer.InfoService interface.
func TestInfoService_Interface(t *testing.T) {
	var _ informer.InfoService = &v1.InfoService{}
	assert.Nil(t, nil) // If we get this far, the test passed.
}

func TestInfoService_NewAndString(t *testing.T) {
	info := v1.NewInfoService(
		"program",
		"license",
		"url",
		"user",
		"date",
		"goVersion",
		"version",
		"revision",
		"branch",
	)
	infoExpect := "(metadata=(program=program, license=license, url=url), versionInfo=(version=version, branch=branch, revision=revision), buildInfo=(go=goVersion, user=user, date=date))"
	assert.Equal(t, infoExpect, info.String())
}
