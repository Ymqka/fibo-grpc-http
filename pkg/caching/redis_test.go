package caching

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCache_SetBigInt(t *testing.T) {
	cache := NewCacheConnection(":6379")

	want := big.NewInt(10)
	var id uint32 = 500

	cache.SetBigInt(id, want)

	got, _ := cache.GetBigInt(id)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	return
}
