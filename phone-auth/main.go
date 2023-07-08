package main

import (
	"fmt"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/labstack/echo/v4"
	"gitlab.com/apito.io/engine/models"
)

type Authentication struct {
}

// Init system Function Implementation
func (g Authentication) Init(cred *protobuff.ThirdPartyCredential) error {
	fmt.Println("Running Init")
	return nil
}

// Migration system Function Implementation
func (g Authentication) Migration() error {
	fmt.Println("Running Migration")
	return nil
}

// Login system Function Implementation
func (g Authentication) SchemaRegister() (*models.ThirdPartyGraphQLSchemas, error) {
	fmt.Println("Registering Schema")

	return &models.ThirdPartyGraphQLSchemas{
		Queries:   nil,
		Mutations: nil,
	}, nil
}

// Register system Function Implementation
func (g Authentication) RESTApiRegister() ([]*models.ThirdPartyRESTApi, error) {
	return []*models.ThirdPartyRESTApi{
		{
			Path:       "/test2",
			Method:     "GET",
			Controller: g.ProviderRegister,
		},
	}, nil
}

// Register system Function Implementation
func (g Authentication) ProviderRegister(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"message": "provider registger",
		"code":    200,
	})
}

// EmailAuth because plugin Name is email-auth exported
// This exported Name is case-sensitive for the extension to load
var PhoneAuth Authentication
