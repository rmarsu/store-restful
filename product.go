package dronemarket

type Product struct {
	Id    string     `json:"id"`
	Image string  `json:"image"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
