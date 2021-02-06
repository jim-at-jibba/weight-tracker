# An easy and practical approach to structuring Golang applications

[Article](https://mbvisti.medium.com/a-practical-approach-to-structuring-go-applications-7f77d7f9c189)

## Project structure

```
weight-tracker
- cmd
  - server
    main.go
- pkg
  - api
    user.go
    weight.go
  - app
    server.go
    handlers.go
    routes.go
  - repository
    storage.go
```
