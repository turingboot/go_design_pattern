package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Errorf("instance is not equal,test fail")
			}
		}
	})
}

func TestGetLazyInstance(t *testing.T) {
	assert.Equal(t, GetLazyInstance(), GetLazyInstance())
}

func BenchmarkGetLazyInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazyInstance() != GetLazyInstance() {
				b.Errorf("instance is not equal,test fail")
			}
		}
	})
}
