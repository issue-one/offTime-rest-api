package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"

	"github.com/issue-one/offTime-rest-api/gen/restapi"
	"github.com/issue-one/offTime-rest-api/gen/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewOffTimeAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	if _, err := os.Stat(".env"); !os.IsNotExist(err) {
		err = godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "offTime"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	portString, ok := os.LookupEnv("PORT")
	if !ok {
		portString = "8080"
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		panic(fmt.Sprintf("Unable to parse port string: %v", portString))
	}
	server.Port = port

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
