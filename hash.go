// A Fast, Minimal Memory, Consistent Hash Algorithm https://arxiv.org/pdf/1406.2294
package hash

// JumpConsistentHash returns the bucket index for the given key using Jump Consistent Hash algorithm.
// numBuckets must be > 0
// the bucket index in the range [0, buckets)
func JumpConsistentHash(key uint64, numBuckets int32) int32 {

	b, j := int64(-1), int64(0)
	for j < int64(numBuckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * float64(int64(1)<<31) / float64((key>>33)+1))
	}
	return int32(b)
}
