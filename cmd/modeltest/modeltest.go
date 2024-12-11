package main

import (
	"fmt"

	"github.com/galeone/tfgo"
)

// import tf "github.com/galeone/tensorflow/tensorflow/go"

func main() {
	model := tfgo.LoadModel("./models/model_occupancy_fix.h5", []string{"serve"}, nil)
	fmt.Println(model)
}
