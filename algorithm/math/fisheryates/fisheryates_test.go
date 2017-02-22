package fisheryates_test

import (
	"fmt"
	"testing"

	"github.com/Akagi201/algorithm-go/algorithm/math/fisheryates"
	"github.com/stretchr/testify/assert"
)

func TestShuffle(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(fisheryates.Shuffle(arr))

	assert.Equal(t, len(arr), len(fisheryates.Shuffle(arr)), "length should equal")
}

func ExampleShuffle() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(fisheryates.Shuffle(arr))
}
