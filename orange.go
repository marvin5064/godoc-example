package fruit

import "time"

var (
	orangeExpiredTime = 2 * 24 * 30 * time.Hour // 60 days
)

type orange struct {
	taste []string
}

// NewOrange creates a orange object associate with taste
func NewOrange() Fruit {
	return &orange{
		taste: []string{"sour", "juicy"},
	}
}

// GetExpireTime returns expiring time for orange
func (o *orange) GetExpireTime() time.Duration {
	return orangeExpiredTime
}
func (o *orange) GetTaste() []string {
	return o.taste
}
