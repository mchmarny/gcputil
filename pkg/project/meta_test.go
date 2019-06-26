package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProjectID(t *testing.T) {

	projectKeys = []string{
		"GCP_PROJECT",
	}

	p, err := GetProjectID()
	assert.Nil(t, err)
	assert.NotEmpty(t, p)

	projectKeys = []string{}

	p, err = GetProjectID()
	assert.NotNil(t, err)

}
