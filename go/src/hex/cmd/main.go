//resposible to orchestrating the startup of application, dependency injection and etc

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EdFazli/StructuringGOApp/internal/adapters/app/api"
	"github.com/EdFazli/StructuringGOApp/internal/adapters/core/arithmetic"
	"github.com/EdFazli/StructuringGOApp/internal/adapters/framework/right/db"
	"github.com/EdFazli/StructuringGOApp/internal/ports"

	gRPC "github.com/EdFazli/StructuringGOApp/internal/adapters/framework/left/grpc"
)

func main() {
	fmt.Println("Main file")

	var err error

	//ports
	var dbaseAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)
	gRPCAdapter = gRPC.NewAdapter(appAdapter)
	gRPCAdapter.Run()
}
