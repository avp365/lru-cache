package cache

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddNullElement(t *testing.T) {

	_, err := NewLRUCache(0)

	require.Error(t, err, "error zero element")

}
func TestAddOneElement(t *testing.T) {

	c, _ := NewLRUCache(1)

	result := c.Add("key1", "value1")
	require.Equal(t, true, result, "fail add")

	result = c.Add("key1", "value1")
	require.Equal(t, false, result, "fail add if var exists")

}
func TestAdd(t *testing.T) {

	c, _ := NewLRUCache(5)

	result := c.Add("key1", "value1")
	require.Equal(t, true, result, "fail add")

	result = c.Add("key1", "value1")
	require.Equal(t, false, result, "fail add if var exists")

}
func TestGet(t *testing.T) {

	c, _ := NewLRUCache(5)

	result, ok := c.Get("key1")
	require.Equal(t, "", result, "fail get if no value for value")
	require.Equal(t, false, ok, "fail get if no value for ok")

	c.Add("key2", "value2")
	result, ok = c.Get("key2")
	require.Equal(t, "value2", result, "fail get")

}
func TestPriority(t *testing.T) {

	c, _ := NewLRUCache(3)

	c.Add("key1", "value1")
	c.Add("key2", "value2")
	c.Add("key3", "value3")
	c.Add("key4", "value4")

	result, ok := c.Get("key1")

	require.Equal(t, "", result, "fail get if no value for value")
	require.Equal(t, false, ok, "fail get if no value for ok")

}
func TestRemove(t *testing.T) {

	c, _ := NewLRUCache(5)

	ok := c.Remove("key1")
	require.Equal(t, false, ok, "fail,element exist")

}
func TestTime_1(t *testing.T) {

	timeStart := time.Now()

	c, _ := NewLRUCache(5)

	key := "someKey"
	c.Add(key, "value1")
	c.Get(key)
	c.Remove(key)

	assert.WithinDuration(t, timeStart, time.Now(), 2000*time.Nanosecond)

}
func TestTime_2(t *testing.T) {

	n := 50
	key := "someKey"
	c, _ := NewLRUCache(n)

	for i := 1; i <= n; i++ {
		c.Add("key"+strconv.Itoa(i), "value1")
	}

	timeStart := time.Now()
	c.Add(key, "value1")
	c.Get(key)
	c.Remove(key)

	assert.WithinDuration(t, timeStart, time.Now(), 2000*time.Nanosecond)

}
