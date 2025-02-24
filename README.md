# Go Error Handling Example

This example demonstrates a common error in Go programming: neglecting to check the error returned by a function that might fail.  The `Calculate` function correctly returns an error if the denominator is zero, but the `main` function doesn't properly handle this error.

## Bug
The `main` function doesn't check for errors returned by `Calculate` and thus the program may crash or produce incorrect results.