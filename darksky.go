package main

import "fmt"
import "net/http"
import "io/ioutil"

const API_KEY = "0bf574e74efefda422944953b78b5df8"
const API_HOST = "api.darkskyapp.com"

func hourlyForecast(lat ,long string) string {
  resp, err := http.Get("https://"+API_HOST+"/v1/forecast/"+API_KEY+"/"+lat+","+long+"")

  if err != nil {}

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  return string(body)
}

func interestingStorms()string{
  resp, err := http.Get("https://"+API_HOST+"/v1/interesting/"+API_KEY+"")

  if err != nil {}

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  return string(body)
}

func main() {
  //fmt.Println(hourlyForecast("42.7243","-73.6927"))
  fmt.Println(interestingStorms())
  fmt.Println("Hello, 世界")
}
