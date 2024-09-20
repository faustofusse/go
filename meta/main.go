package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
)

func main() {
    http.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
        res.Header().Add("Content-Type", "text/html")
        path := req.URL.Path
        io.WriteString(res, fmt.Sprintf(`
            <html>
                <head>
                    <meta name="go-import" content="go.fausto.ar git https://github.com/faustofusse/go%v">
                </head>
                <body>
                </body>
            </html>
        `, path))
    })
    // http.ListenAndServe(":3000", http.DefaultServeMux)
    lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
