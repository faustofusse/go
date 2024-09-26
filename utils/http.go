package utils

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"

    "github.com/gorilla/schema"
	"github.com/labstack/echo/v4"
)

func Bind[T any](req *http.Request) (*T, error) {
    contentType := req.Header.Get("Content-Type")
    if req.ContentLength == 0 {
        return nil, errors.New("empty body")
    }
    result := new(T)
    switch {
        case strings.HasPrefix(contentType, "application/json"):
            err := json.NewDecoder(req.Body).Decode(result)
            return result, err
        case strings.HasPrefix(contentType, "application/xml"):
            err := xml.NewDecoder(req.Body).Decode(result)
            return result, err
        case strings.HasPrefix(contentType, "application/x-www-form-urlencoded"):
            err := req.ParseForm()
            if err != nil { return nil, err }
            decoder := schema.NewDecoder()
            decoder.SetAliasTag("form")
            err = decoder.Decode(result, req.Form)
            return result, err
        case strings.HasPrefix(contentType, "multipart/form-data"):
            err := req.ParseMultipartForm(req.ContentLength)
            if err != nil { return nil, err }
            decoder := schema.NewDecoder()
            decoder.SetAliasTag("form")
            err = decoder.Decode(result, req.MultipartForm.Value)
            return result, err
    }
    return nil, errors.New("unsupported content type")
}

func GetBody[T any](ctx echo.Context) (*T, error) {
    var body T
    err := ctx.Bind(&body)
    if err != nil {
        return nil, errors.New(fmt.Sprintf("Error: body could not be casted properly: %v", err.Error()))
    }
    return &body, err
}

func GetVariable[T any](ctx echo.Context, key string) (*T, error) {
    value := ctx.Get(key)
    casted, ok := value.(T)
    if !ok {
        return nil, errors.New(fmt.Sprintf("Error: variable '%s' could not be casted properly", key))
    }
    return &casted, nil
}

func GetFormBody[T any](ctx echo.Context) (*T, error) {
    var body T
    err := ctx.Bind(&body)
    // TODO: ineficiente ?
    if err == nil {
        if bytes, e := json.Marshal(body); e == nil {
            fmt.Printf("[BODY] %+v\n", string(bytes))
        }
    }
    return &body, err
}

func GetAuthToken(ctx echo.Context) string {
    reqToken := ctx.Request().Header.Get("Authorization")
    if reqToken == "" {
        return ""
    }
    splitToken := strings.Split(reqToken, "Bearer")
    return strings.TrimSpace(splitToken[1])
}

// func GetBody[T any](ctx *gin.Context) (*T, error) {
//     var body T
//     err := ctx.BindJSON(&body)
//     // TODO: ineficiente ?
//     if err == nil {
//         if bytes, e := json.Marshal(body); e == nil {
//             fmt.Printf("[BODY] %+v\n", string(bytes))
//         }
//     }
//     return &body, err
// }

// func GetVariable[T any](ctx *gin.Context, key string) (*T, error) {
//     value, exists := ctx.Get(key)
//     if !exists {
//         return nil, errors.New(fmt.Sprintf("Error: variable '%s' not defined in context", key))
//     }
//     casted := value.(T)
//     return &casted, nil
// }
