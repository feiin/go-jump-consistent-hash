package hash

import (
	"fmt"
	"testing"
)

func BenchmarkHash(b *testing.B) {
	bucketNum := int32(16)
	for i := 0; i < b.N; i++ {
		if bucketIndex := JumpConsistentHash(uint64(i), bucketNum); bucketIndex > bucketNum {
			b.Fatal("invalid JumpConsistentHash:", bucketIndex)
		}
	}
}

func BenchmarkHashString(b *testing.B) {
	bucketNum := int32(16)
	for i := 0; i < b.N; i++ {
		key := fmt.Sprintf("key-%d", i)
		if bucketIndex := JumpHashString(key, bucketNum); bucketIndex > bucketNum {
			b.Fatal("invalid JumpHashString:", bucketIndex)
		}
	}
}
