package main

import "darksky-go/darksky"
import "fmt"

func main() {

  dSky := darksky.DarkSky{ "0bf574e74efefda422944953b78b5df8"}
  fmt.Println(dSky.HourlyForecast(42.7243,-73.6927))

  //fmt.Println(hourlyForecast("42.7243","-73.6927"))
  //fmt.Println(interestingStorms())
}
