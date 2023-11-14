package formatter

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(FormatCurrency(6598721324, "USD"))
}