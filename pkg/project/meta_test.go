package project

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProjectID(t *testing.T) {

	projectKeys = []string{
		"GCP_PROJECT",
	}

	p, err := GetID()
	assert.Nil(t, err)
	assert.NotEmpty(t, p)

	projectKeys = []string{}

	p, err = GetID()
	assert.NotNil(t, err)

}
