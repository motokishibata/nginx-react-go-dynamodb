package main

import (
    "fmt"
    "net/http"
    "context"
    "log"
    "github.com/aws/aws-sdk-go-v2/aws"	
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
    // controller
    http.HandleFunc("/", echoHello)
    // port
    http.ListenAndServe(":7000", nil)
}

func echoHello(w http.ResponseWriter, r *http.Request) {
    cfg, err := config.LoadDefaultConfig(context.TODO(),
    config.WithEndpointResolver(aws.EndpointResolverFunc(
        func(service, region string) (aws.Endpoint, error) {
            return aws.Endpoint{URL: "http://dynamodb-local:8000"}, nil
        })),
    )

    if err != nil {
        log.Fatal(err)
    }

    svc := dynamodb.NewFromConfig(cfg)

    resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{
        Limit: aws.Int32(5),
    })
    if err != nil {
        log.Fatalf("failed to list tables, %v", err)
    }

    data := "<h1>DynamoDb</h1>"
    for _, tableName := range resp.TableNames {
        data += fmt.Sprintf("<p>%s</p>", tableName)
    }

    fmt.Fprint(w, data)
}