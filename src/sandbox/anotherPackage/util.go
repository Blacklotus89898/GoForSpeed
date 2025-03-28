package util

// Function to double a number
func Double(x int) int {
    return x * 2
}

// Unexported (private to the package)
func doubleInternal(x int) int { 
	return x * 2
 }
