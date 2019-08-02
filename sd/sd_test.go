package sd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetric(t *testing.T) {

	ctx := context.Background()
	c, err := NewClient(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, c)

	err = c.Publish(ctx, "test1", "test-metric", float64(1.23))
	assert.Nil(t, err)

}
