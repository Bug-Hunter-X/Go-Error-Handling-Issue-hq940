func Calculate(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("division by zero")
  }
  return a / b, nil
}

func main() {
  result, err := Calculate(10, 0)
  if err != nil {
    fmt.Printf("Error: %%v\n", err) //Improved error handling
    return //Exit early after error
  }
  fmt.Printf("Result: %%d\n", result)
} 