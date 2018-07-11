package main

import (
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/go-pg/pg"
  "github.com/go-pg/pg/orm"
  "fmt"
  "log"
  "os"
  "net/http"
  "errors"
)

const (
  AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
  PostgresUrl = "POSTGRES_URL"
)

type PageView struct {
    Id         int64
    Referer    string
    RemoteAddr string
}

func (u PageView) String() string {
    return fmt.Sprintf("PageView<%d %s %s>", u.Id, u.Referer, u.RemoteAddr)
}

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

func SetupDB() {
  url, present := os.LookupEnv(PostgresUrl)
  if !present {
    panic(errors.New(fmt.Sprintf("missing env %s", PostgresUrl)))
  }

  options, err := pg.ParseURL(url)
  if err != nil {
    panic(err)
  }

  db := pg.Connect(options)
  defer db.Close()

  for _, model := range []interface{}{&PageView{}} {
    err = db.CreateTable(model, &orm.CreateTableOptions{})

    if err != nil {
      panic(err)
    }
  }
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
    Body:       "",
    StatusCode: 200,
  }, nil
}

func MockHandler() {
  http.HandleFunc("/analytics.css", func(w http.ResponseWriter, r *http.Request) {
    log.Printf("")
    log.Printf("Request:")
    log.Printf("remote address: %s", r.RemoteAddr)
    log.Printf("referer: %s", r.Referer())
  })

  http.Handle("/", http.FileServer(http.Dir("build")))
  log.Printf("Listening at http://127.0.0.1:4000")
  http.ListenAndServe(":4000", nil)
}
