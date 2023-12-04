package calculations

const hundred = 100

// CalculatePercentage calculates the percentage of a given value.
// It takes the original value and the percentage to calculate.
func CalculatePercentage(originalValue float32, percentage float32) float32 {
    return (percentage * originalValue) / hundred
}

// CalculateValueFromPercentage calculates the original value based on a given percentage and a calculated value.
// It takes the percentage and the calculated value as parameters.
func CalculateValueFromPercentage(percentage float32, calculatedValue float32) float32 {
    // Formula: originalValue = (calculatedValue * 100) / percentage
    return (calculatedValue * hundred) / percentage
}
