package lasagna
import "fmt"

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,time int) int {
	if time == 0 {
        return len(layers) * 2
    }else{
        return len(layers) * time
    }
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    var noodlesTotal int = 0;
    var sauceTotal float64 = 0;
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodlesTotal++
        }else if layers[i] == "sauce" {
            sauceTotal++
        }
    }
    return noodlesTotal * 50, sauceTotal * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient (friendsList []string,myList []string) []string{
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
    return myList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe (quantities []float64,portion int) []float64{
    var perSlice float64 = 0
    var value float64 = 0
    var amount []float64
	perSlice = float64(portion) / 2
    for i := 0; i < len(quantities); i++ {
        fmt.Printf("slice = %f", perSlice)
		value = quantities[i] * perSlice
        amount = append(amount, value)
    }
    return amount
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
