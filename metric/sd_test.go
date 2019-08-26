package metric

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

	err = c.Publish(ctx, "test1", "test-metric", float32(1.23))
	assert.Nil(t, err)

	err = c.Publish(ctx, "test1", "test-metric", int64(4))
	assert.Nil(t, err)

	err = c.Publish(ctx, "test1", "test-metric", int(1))
	assert.Nil(t, err)

}

func TestMetricWithSource(t *testing.T) {

	ctx := context.Background()
	c, err := NewClientWithSource(ctx, "test2")
	assert.Nil(t, err)
	assert.NotNil(t, c)

	err = c.PublishForSource(ctx, "test-metric", float64(1.23))
	assert.Nil(t, err)

	err = c.PublishForSource(ctx, "test-metric", float32(1.23))
	assert.Nil(t, err)

	err = c.PublishForSource(ctx, "test-metric", int64(4))
	assert.Nil(t, err)

	err = c.PublishForSource(ctx, "test-metric", int(1))
	assert.Nil(t, err)

}
