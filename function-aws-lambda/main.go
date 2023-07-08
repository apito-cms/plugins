package main

import (
	"fmt"
	"github.com/apito-cms/buffers/protobuff"
)

type AWSLambda struct {
	Config *protobuff.FunctionProviderConfig
}

var Config *protobuff.FunctionProviderConfig

// Init system Function Implementation
func (g AWSLambda) Init(config *protobuff.FunctionProviderConfig) error {
	fmt.Println(fmt.Sprintf("Running Function Init উইথ %+v", config))
	Config = config
	return nil
}

// Execute system Function Implementation
func (g AWSLambda) Execute(request interface{}) (interface{}, error) {
	fmt.Println(fmt.Sprintf("Running Function Execute উইথ %+v", Config))
	if g.Config != nil {
		for _, v := range g.Config.EnvVars {
			fmt.Println(v.Key)
		}
	}

	return nil, nil
}

// LambdaProvider because plugin Name is email-auth exported
// This exported Name is case-sensitive for the extension to load
var LambdaProvider AWSLambda
