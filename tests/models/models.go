package goar_test_models

type Vehicle struct {
	Year  int
	Make  string
	Model string
}

type Automobile struct {
	Vehicle
}

type Motorcycle struct {
	Vehicle
}
