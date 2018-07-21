package fruit

import "time"

var (
	orangeExpiredTime = 2 * 24 * 30 * time.Hour // 60 days
)

type orange struct {
	taste []string
}

func NewOrange() Fruit {
	return &orange{
		taste: []string{"sour", "juicy"},
	}
}

func (o *orange) GetExpireTime() time.Duration {
	return orangeExpiredTime
}
func (o *orange) GetTaste() []string {
	return o.taste
}
