// Package weather le indica a Go a qué paquete pertenece el archivo.
package weather
var (
    // CurrentCondition se usa para saber la temperatura actual.
	CurrentCondition string
	// CurrentLocation se usa para saber la ciudad actual.
    CurrentLocation  string
)
// Forecast lo que hace es mostrar el estado actual del pais en base a su temperatura.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
