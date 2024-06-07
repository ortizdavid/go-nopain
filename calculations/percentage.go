package calculations

// PercentageOfvalue calculates the percentage of a given value.
// It takes the original value and the percentage to calculate.
func PercentageOfValue(originalValue float32, percentage float32) float32 {
    return (percentage * originalValue) / 100
}

// alueFromPercentage calculates the original value based on a given percentage and a calculated value.
// It takes the percentage and the calculated value as parameters.
func ValueFromPercentage(percentage float32, calculatedValue float32) float32 {
    return (calculatedValue * 100) / percentage
}
