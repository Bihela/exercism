//Package weather this is package your using.
package weather

var (
    // CurrentCondition string variable.
	CurrentCondition string
    // CurrentLocation string variable.
	CurrentLocation  string
)
// Forecast your are using city and condition has parameters for the function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
