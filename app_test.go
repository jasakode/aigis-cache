package aigiscache

import (
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	// aigis-cache
	New()
	by, _ := os.ReadFile("aigiscache.bin")
	fmt.Printf("%b", by)
	fmt.Println("HAllo")
}
