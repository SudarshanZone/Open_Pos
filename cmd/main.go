package main

import (
	"context"
	"log"
	"time"

	"github.com/SudarshanZone/Open_Pos/config"
	positions "github.com/SudarshanZone/Open_Pos/internal/generated"
	"github.com/SudarshanZone/Open_Pos/internal/service"
	//pb "github.com/SudarshanZone/internal/proto"
	
	
)
func main() {
    serviceName := "main"
    fileName := "config/EnvConfig.ini"
    
    // Set up the config manager.
    cm := &config.ConfigManager{}
    
    // Load the PostgreSQL config.
    dbConfig, err := cm.LoadPostgreSQLConfig(serviceName, fileName)
    if err != nil {
        log.Fatalf("Failed to load database config: %v", err)
    }
    
    // Get the database connection.
    connectionStatus := cm.GetDatabaseConnection(serviceName, *dbConfig)
    if connectionStatus != 0 {
        log.Fatalf("Failed to connect to the database")
    }
    
    // Get the database handle.
    db := cm.GetDB(serviceName)
    
    // Set up the server with the database connection.
    server := &service.Server{ 
		Db: db,
	}
    
    // Create a request object with sample data.
    request := &positions.FnoPositionRequest{
        FcpClmMtchAccnt: "8989890000",
    }
    
    // Create a context with a timeout for the RPC call.
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    
    // Call the GetFNOPosition method.
    response, err := server.GetFNOPosition(ctx, request)
    if err != nil {
        log.Fatalf("Could not get FNO position: %v", err)
    }
    
    // Log the response.
    log.Printf("FNO Position Response: %v", response)
    
    // Log the database connection.
    log.Printf("Database connection established successfully: %v", db)
}
