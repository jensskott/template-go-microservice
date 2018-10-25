package collect

type Collect struct {
	Name string
	Age  int
}

func Json() Collect {
	d := Collect{
		Name: "my-name",
		Age:  2,
	}

	return d
}
