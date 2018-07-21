package fruit

import "time"

// Fruit object commonly used by apple and orange
type Fruit interface {
	// GetExpireTime returns expiring time for corresponding fruit
	GetExpireTime() time.Duration
	// GetTaste returns taste for corresponding fruit
	GetTaste() []string
}
