package darksky

import "net/http"
import "io/ioutil"
import "strconv"

const API_HOST = "api.darkskyapp.com"

type DarkSky struct {
  ApiKey string
}

func webRequest(url string) string{
  resp, err := http.Get(url)

  if err != nil {}

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  return string(body)
}

func (d DarkSky) HourlyForecast(lat ,long float64) string {
  return webRequest("https://"+API_HOST+"/v1/forecast/"+d.ApiKey+"/"+strconv.FormatFloat(lat,'f',4,32)+
  ","+strconv.FormatFloat(long,'f',4,32)+"")
}

func (d DarkSky) BriefHourlyForecast(lat ,long float64) string {
  return webRequest("https://"+API_HOST+"/v1/brief_forecast/"+d.ApiKey+"/"+strconv.FormatFloat(lat,'f',4,32)+
  ","+strconv.FormatFloat(long,'f',4,32)+"")
}

func (d DarkSky) InterestingStorms()string{
  return webRequest("https://"+API_HOST+"/v1/interesting/"+d.ApiKey+"")
}

