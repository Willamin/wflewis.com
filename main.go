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
  "time"
)

const (
  AWSLambdaFunctionVersion = "AWS_LAMBDA_FUNCTION_VERSION"
  PostgresUrl = "POSTGRES_URL"
)

type PageView struct {
    Id         int64
    Referer    string
    RemoteAddr string
    AccessedAt time.Time
}

func (u PageView) String() string {
    return fmt.Sprintf("PageView<%d %s %s %s>", u.Id, u.Referer, u.RemoteAddr, u.AccessedAt)
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

func SetupDB() *pg.DB {
  url, present := os.LookupEnv(PostgresUrl)
  if !present {
    panic(errors.New(fmt.Sprintf("missing env %s", PostgresUrl)))
  }

  options, err := pg.ParseURL(url)
  if err != nil {
    panic(err)
  }

  db := pg.Connect(options)

  db.CreateTable(&PageView{}, &orm.CreateTableOptions{})
  return db
}

func AWSHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  address, present := request.Headers["x-forwarded-for"]
  if !present {
    address = request.Headers["X-Forwarded-For"]
  }

  referer, present := request.Headers["referer"]
  if !present {
    referer = request.Headers["Referer"]
  }

  pv := PageView{
    Referer: referer,
    RemoteAddr: address,
    AccessedAt: time.Now(),
  }


  log.Printf("%v", pv)

  for k, v := range request.Headers {
    log.Printf("request.Headers[\"%s\"]=\"%s\"", k, v)
  }

  db := SetupDB()

  err := db.Insert(&pv)
  if err != nil {
      panic(err)
  }
  db.Close()

  return events.APIGatewayProxyResponse{
    Body:       "",
    StatusCode: 200,
  }, nil
}

func MockHandler() {
  http.HandleFunc("/analytics.css", func(w http.ResponseWriter, r *http.Request) {
    pv := PageView{
      Referer: r.RemoteAddr,
      RemoteAddr: r.Referer(),
      AccessedAt: time.Now(),
    }
    log.Printf("%v", pv)

    db := SetupDB()

    err := db.Insert(&pv)
    if err != nil {
        panic(err)
    }
    db.Close()
  })

  http.Handle("/", http.FileServer(http.Dir("build")))
  log.Printf("Listening at http://127.0.0.1:4000")
  http.ListenAndServe(":4000", nil)
}
