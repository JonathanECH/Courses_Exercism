package purchase

import (
	"fmt"
	"strings"
)
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var required bool
	if strings.ToLower(kind) == "car" || strings.ToLower(kind) == "truck"{
		required = true
	}else{
		required = false
	}
	return required
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var msg string
	if option1 < option2 {
		msg = fmt.Sprintf("%s is clearly the better choice.", option1)
	} else {
		msg = fmt.Sprintf("%s is clearly the better choice.", option2)
	}
	return msg
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var percentage float64
	if age < 3 {
		percentage = originalPrice * 0.8
	} else if age < 10 {
		percentage = originalPrice * 0.7
	} else {
		percentage = originalPrice * 0.5
	}
	return percentage
}
