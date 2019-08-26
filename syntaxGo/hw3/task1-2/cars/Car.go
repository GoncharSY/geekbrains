package cars

import "time"

// Car описывает автомобиль.
type Car struct {
	Model   string
	Type    string
	Release time.Time
	Engine  Engine
	Trunk   Trunk
	Windows []Window
}
