package formatter

import (
	"fmt"
	"testing"
)

func Test_Currency(t *testing.T) {
	fmt.Println(FormatCurrency(6598721324, "USD"))
}