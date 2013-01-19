darksky-go
==========

##What is this?

Darksky Go is a wrapper around the [DarkskyApp api](https://developer.darkskyapp.com/), I am for it to be
feature complete at some point and idiomatic.

##Why was it made?

I wanted a worth while project to start playing with Go. After looking
at someone's implementation of a wrapper of the [DarkSkyApp API in
Clojure](https://github.com/jdhollis/darksky-clojure
) I figured that this would be a fun place to start.

##So what can I do with it?

Really anything your mind can come up with. An idea could be to use this
with a reverse geocoding api (ala google) to see what the weather stats
for a given address

##So how do I use this?

1. Create a DarkSky object

dSky := DarkSky{<API key>}

2. With the DarkSky object, you should now be able to make calls against
   the api.

##What methods exist on the DarkSky object?

1. [hourlyForecast](https://developer.darkskyapp.com/docs/forecast) 
2. [briefHourlyForecast](https://developer.darkskyapp.com/docs/forecast)
3. [interestingStorms](https://developer.darkskyapp.com/docs/interesting)

