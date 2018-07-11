package main

import (
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "errors"
  "log"
  "os"
  "net/http"
)

const (
  AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
)

func main() {
  _, ok := os.LookupEnv(AWSLambdaFunctionVersion)
  if ok {
    log.Printf("Running in AWS lambda environment, starting lambda handler.")
    lambda.Start(AWSHandler)
    os.Exit(0)
  }

  log.Printf("Not running in AWS lambda environment, starting mock handler.")
  MockHandler()
  os.Exit(0)
}

func AWSHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  log.Printf("")
  log.Printf("Request:")

  address, present := request.Headers["x-forwarded-for"]
  if !present {
    address = request.Headers["X-Forwarded-For"]
  }
  log.Printf("remote address: %s", address)

  referer, present := request.Headers["referer"]
  if !present {
    referer = request.Headers["Referer"]
  }
  log.Printf("referer: %s", referer)

  for k, v := range request.Headers {
    log.Printf("request.Headers[\"%s\"]=\"%s\"", k, v)
  }

  return events.APIGatewayProxyResponse{
    Body:       "not implemented yet",
    StatusCode: 501,
  }, errors.New("not implemented yet")
}

func MockHandler() {
  http.HandleFunc("/analytics.css", func(w http.ResponseWriter, r *http.Request) {
    log.Printf("")
    log.Printf("Request:")
    log.Printf("remote address: %s", r.RemoteAddr)
    log.Printf("referer: %s", r.Referer())

    http.Error(w, "not implemented yet", http.StatusNotImplemented)
  })

  http.Handle("/", http.FileServer(http.Dir("build")))
  log.Printf("Listening at http://127.0.0.1:4000")
  http.ListenAndServe(":4000", nil)
}
