package main

//import "darksky/darksky"

import "fmt"
import "net/http"
import "io/ioutil"
import "strconv"

const API_HOST = "api.darkskyapp.com"

type DarkSky struct {
  apiKey string
}

func webRequest(url string) string{
  resp, err := http.Get(url)

  if err != nil {}

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  return string(body)
}

func (d DarkSky) hourlyForecast(lat ,long float64) string {
  return webRequest("https://"+API_HOST+"/v1/forecast/"+d.apiKey+"/"+strconv.FormatFloat(lat,'f',4,32)+","+strconv.FormatFloat(long,'f',4,32)+"")
}

func (d DarkSky) briefHourlyForecast(lat ,long float64) string {
  return webRequest("https://"+API_HOST+"/v1/brief_forecast/"+d.apiKey+"/"+strconv.FormatFloat(lat,'f',4,32)+","+strconv.FormatFloat(long,'f',4,32)+"")
}

func (d DarkSky) interestingStorms()string{
  return webRequest("https://"+API_HOST+"/v1/interesting/"+d.apiKey+"")
}

func main() {

  dSky := DarkSky{ "0bf574e74efefda422944953b78b5df8"}
  fmt.Println(dSky.hourlyForecast(42.7243,-73.6927))


  //fmt.Println(hourlyForecast("42.7243","-73.6927"))
  //fmt.Println(interestingStorms())
}
