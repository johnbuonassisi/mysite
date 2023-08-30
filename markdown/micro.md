# Skip Tests

Do you need to comment out a test to fix it later? You've come back to fix your test and it no longer compiles
because of other code changes. To avoid this consider using `testing.Skip`

```
func TestSomething(t *testing.T) {
	t.Skip("Comeback to fix this later, but keep the test compiling!")
}
```

# Table Test

Table tests are great for when you need to test a number of different scenarios and you find you are repeating a lot
of test code.

# Multiple test cases in a single tst function
