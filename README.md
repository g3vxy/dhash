## Introduction

A dHash implementation with Golang. It can calculate hash value of given image or Hamming Distance between two 128-bit dHash values.

## Usage

```go
package main

import (
	"fmt"

	"github.com/g3vxy/dhash"
)

func main() {
	similarity := dhash.CalculateHammingDistance(dhash.CalculateHash("1.jpg"), dhash.CalculateHash("2.jpg"))
	fmt.Printf("Similarity is %%%f", similarity*100)
}
```