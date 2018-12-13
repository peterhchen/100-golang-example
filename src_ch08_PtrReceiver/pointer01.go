package main
import "fmt"
const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934 // i miles = 1.6 KM.

type car struct {
	gas_pedal uint16 // min 0 max 65535
	brake_pedal uint16
	steering_wheel int16  // -32K to +32K
	top_speed_kmh float64
}

// cvalue receiver. Make a copy and pass
func (c car) kmh () float64 {
	return float64 (c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
} 

func (c car) mph () float64 {
	return float64 (c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
}

// pointer receiver. Share the memory
func (c *car) new_top_speed (newspeed float64) {
	c.top_speed_kmh = newspeed
}

func main () {
	a_car := car {gas_pedal: 65000, 
		brake_pedal: 0, 
		steering_wheel: 12561, 
		top_speed_kmh: 225.0 } 
	//a_car := car {22341, 0, 12562, 225.0} 
	fmt.Println (a_car.gas_pedal)
	fmt.Println (a_car.kmh()) // access method with type struct
	fmt.Println (a_car.mph())
	a_car.new_top_speed (500)
	fmt.Println (a_car.kmh())
	fmt.Println (a_car.mph())
}