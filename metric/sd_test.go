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

	labels := make(map[string]string, 0)
	labels["test_metric"] = "test1"
	labels["other_metric"] = "test2"

	err = c.Publish(ctx, "test-metric", float64(1.23), labels)
	assert.Nil(t, err)

	err = c.Publish(ctx, "test-metric", float32(1.23), labels)
	assert.Nil(t, err)

	err = c.Publish(ctx, "test-metric", int64(4), labels)
	assert.Nil(t, err)

	err = c.Publish(ctx, "test-metric", int(1), labels)
	assert.Nil(t, err)

}
