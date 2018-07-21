package fruit

import "time"

type Fruit interface {
	GetExpireTime() time.Duration
	GetTaste() []string
}
