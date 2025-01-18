├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/           # controllers yerine handlers
│   │   ├── auth.go        # _handler suffix'i opsiyonel
│   │   └── user.go
│   ├── middleware/
│   │   └── auth.go
│   ├── models/
│   │   └── user.go
│   ├── repositories/
│   │   └── user.go
│   └── services/
│       ├── auth.go
│       └── user.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   └── utils/
│       ├── jwt.go
│       └── access.go
├── .env
├── go.mod
└── README.md