package main

import "fmt"
import "net/http"
import "io/ioutil"

const API_KEY = "0bf574e74efefda422944953b78b5df8"
const API_HOST = "api.darkskyapp.com"


func webRequest(url string) string{
  resp, err := http.Get(url)

  if err != nil {}

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  return string(body)
}

func hourlyForecast(lat ,long string) string {
  return webRequest("https://"+API_HOST+"/v1/forecast/"+API_KEY+"/"+lat+","+long+"")
}

func interestingStorms()string{
  return webRequest("https://"+API_HOST+"/v1/interesting/"+API_KEY+"")
}

func main() {
  fmt.Println(hourlyForecast("42.7243","-73.6927"))
  //fmt.Println(interestingStorms())
}
