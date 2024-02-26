package meteorology

import (
	"fmt"
	"strconv"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	if t == Celsius {
		return "°C"
	}
	return "°F"
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t *Temperature) String() string {
	v := " °C"
	if t.unit == Fahrenheit {
		v = " °F"
	}
	return strconv.Itoa(t.degree) + v
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	if s == KmPerHour {
		return "km/h"
	}
	return "mph"
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s *Speed) String() string {
	v := " km/h"
	if s.unit == MilesPerHour {
		v = " mph"
	}
	return strconv.Itoa(s.magnitude) + v
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m *MeteorologyData) String() string {
	return m.location + ": " + m.temperature.String() + fmt.Sprintf(", Wind %v at %v, ", m.windDirection, m.windSpeed.String()) + strconv.Itoa(m.humidity) + "% " + "Humidity"
}
