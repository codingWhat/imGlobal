package ws

type ConnectionManager struct {
	Buckets []*Bucket
}

func (cm *ConnectionManager) GetBucket(id uint64) *Bucket {
	return cm.Buckets[id%uint64(len(cm.Buckets))]
}
