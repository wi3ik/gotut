package main

import ("fmt"
        "net/http"
        "time")



type car struct {

	gas_pedal uint16 // min - 0, max 65535
	brake_pedal uint16
	steering_wheel int16 // [-32k, +32k]
	top_speed_kmh float64
}

func (c car) foo() float64 {
	return float64(c.gas_pedal) * float64(c.top_speed_kmh)
}

// modify field method
func (c *car) new_top_speed(new_speed float64) {
	c.top_speed_kmh = new_speed
}

func newer_top_speed(c car, speed float64) car {

	c.top_speed_kmh = speed
	return c
}


func index_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whoa, Go is neat!\n")

  dt := time.Now()
  fmt.Fprintln(w, "Current date and time is: ", dt.String(), "\n")
}

func add_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whoa, ADD Go is neat!\n")

  dt := time.Now()
  fmt.Fprintln(w, "Current date and time is: ", dt.String(), "\n")
}


func main() {
  //var a_car = car{gas_pedal: 123, brake_pedal: 0, steering_wheel: 123, top_speed_kmh: 23.3}
/*
  a_car := car{gas_pedal: 123, brake_pedal: 0, steering_wheel: -123, top_speed_kmh: 45.56}
  fmt.Println(a_car.gas_pedal, a_car.top_speed_kmh, a_car.foo())
  a_car.new_top_speed(222)
  a_car = newer_top_speed(a_car, 222)
  fmt.Println(a_car.gas_pedal, a_car.top_speed_kmh, a_car.foo())
*/


  http.HandleFunc("/", index_handler)
  http.HandleFunc("/add", add_handler)  
  http.ListenAndServe(":8000", nil)



}

/*
 http://users-macbook-pro.local:8000/
 localhost:8000
 <IP>:8000
*/

