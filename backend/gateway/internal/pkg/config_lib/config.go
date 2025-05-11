package config_lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetEnvInt(valueEnv string) int {
	value := os.Getenv(valueEnv)
	if value == "" {
		panic(valueEnv + " is not set")
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		panic(valueEnv + " is not int")
	}
	return valueInt
}

func GetEnvDuration(valueEnv string) time.Duration {
	value := os.Getenv(valueEnv)
	if value == "" {
		panic(valueEnv + " is not set")
	}
	valueDuration, err := time.ParseDuration(value)
	if err == nil {
		return valueDuration
	}
	valueDuration, err = parseExtendedDuration(value)
	if err != nil {
		panic(err)
	}
	return valueDuration
}

func parseExtendedDuration(value string) (time.Duration, error) {
	var numStr string
	var unit string

	for i, r := range value {
		if r >= '0' && r <= '9' || r == '.' || r == '-' {
			numStr += string(r)
		} else {
			unit = strings.ToLower(value[i:])
			break
		}
	}

	if numStr == "" {
		return 0, fmt.Errorf("no number found in %q", value)
	}
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, fmt.Errorf("error in parseExtendedDuration: value=%s", value)
	}

	switch unit {
	case "d":
		return time.Duration(num * float64(24*time.Hour)), nil
	case "w":
		return time.Duration(num * float64(7*24*time.Hour)), nil
	case "mo":
		return time.Duration(num * float64(30*24*time.Hour)), nil
	case "y":
		return time.Duration(num * float64(365*24*time.Hour)), nil
	default:
		return 0, fmt.Errorf("error in parseExtendedDuration: value=%s. unit must "+
			"be in [d, w, mo, y], but unit=%s", value, unit)
	}
}

func GetEnvStr(valueEnv string) string {
	value := os.Getenv(valueEnv)
	if value == "" {
		panic(valueEnv + " is not set")
	}
	return value
}

func GetEnvAddress(hostEnv, portEnv string) string {
	port := GetEnvStr(portEnv)
	host := GetEnvStr(hostEnv)
	return host + ":" + port
}

func GetEnvBool(valueEnv string) bool {
	value := os.Getenv(valueEnv)
	if value == "true" {
		return true
	}
	return false
}
