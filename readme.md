go-graphql-api/
├── go.mod                # Like .csproj
├── go.sum                # Package lock file
├── server.go             # Entry point (Main)
├── graph/                # GraphQL Layer
│   ├── model/            # Generated Models (DTOs)
│   ├── schema.graphqls   # The "Contract"
│   ├── schema.resolvers.go # The "Controller" logic
│   └── generated.go      # Internal Gqlgen magic (Don't touch)
└── internal/             # Private business logic (Optional but good practice)
    └── database.go       # Mock Database / Service layer