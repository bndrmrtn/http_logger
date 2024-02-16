# HTTP Logger

It's a simple tool to monitor your requests during development.

## How to?

SUPER EASY!
Just define a variable `logger` and then use it's `Log` method:

```go
var logger = http_logger.New()

func someFunc(res http.ResponseWriter, req *http.Request) {
    logger.Log(req.Method, req.URL.String())
}
```

With `Echo`:
```go
var logger = http_logger.New()

func Logger(next echo.HandleFunc) echo.HandleFunc {
    return func(c echo.Context) error {
        logger.Log(c.Request().Method, c.Request().URL.String())
        return next(c)
    }
}
```

That's all, and you get beautiful log messages on your terminal under development.