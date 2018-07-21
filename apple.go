package fruit

import "time"

var (
	appleExpiredTime = 1 * 24 * 30 * time.Hour // 30 days
)

type apple struct {
	taste []string
}

func NewApple() Fruit {
	return &apple{
		taste: []string{"sweet", "juicy"},
	}
}

func (a *apple) GetExpireTime() time.Duration {
	return appleExpiredTime
}
func (a *apple) GetTaste() []string {
	return a.taste
}
