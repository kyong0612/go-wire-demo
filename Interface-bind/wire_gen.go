// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package interfacebind

// Injectors from wire.go:

func initializeBar() (Bar, error) {
	myFooer := providerMyFooer()
	bar := providerBar(myFooer)
	return bar, nil
}
