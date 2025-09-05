# Neo4j SQL Generator

***in development***

A Domain Driven Design (DDD) application for designing Neo4j graph schemas visually and generating corresponding SQL database schemas with full table structures, relationships, and constraints.

## Overview

This project provides a comprehensive toolset for:
- **Visual Graph Design**: Interactive web-based designer for Neo4j graph schemas
- **Schema Analysis**: Analysis of existing Neo4j databases to understand structure
- **SQL Generation**: Automatic generation of SQL DDL scripts for multiple database engines
- **Bidirectional Workflow**: Integration with [mysql-graph-visualizer](../mysql-graph-visualizer) for complete SQL ↔ Neo4j ecosystem

## Key Features

- 🎨 **Interactive Graph Designer** - Visual drag-and-drop interface for designing graph schemas
- 🔍 **Neo4j Schema Analysis** - Analyze existing Neo4j databases and infer optimal SQL structures  
- 🗃️ **Multi-Database Support** - Generate SQL for MySQL, PostgreSQL, SQLite
- ⚡ **Template Engine** - Customizable SQL generation templates
- 🏗️ **DDD Architecture** - Clean, maintainable, and extensible codebase
- 🌐 **Multiple Interfaces** - Web UI, CLI, GraphQL API
- 🔄 **Real-time Preview** - Live preview of generated SQL schemas

## Tech Stack

- **Language**: Go 1.22.5+
- **Architecture**: Domain Driven Design (DDD)
- **Graph Database**: Neo4j 4.4+
- **Web Framework**: Gorilla Mux
- **Frontend**: Vanilla JS with D3.js/Vis.js for graph visualization
- **APIs**: REST, GraphQL (gqlgen)
- **Configuration**: Viper + YAML
- **Logging**: Logrus
- **Testing**: Testify

## Project Structure (DDD)

```
neo4j-sql-generator/
├── cmd/                        # Application entry point
├── internal/
│   ├── domain/                 # Domain Layer - Core business logic
│   │   ├── aggregates/         # Domain aggregates (Schema, Graph Design, Generation)
│   │   ├── entities/           # Domain entities
│   │   ├── value_objects/      # Value objects
│   │   ├── events/             # Domain events
│   │   ├── services/           # Domain services
│   │   └── specifications/     # Business rules specifications
│   ├── application/            # Application Layer - Use cases
│   │   ├── ports/              # Interface definitions (Ports)
│   │   ├── services/           # Application services
│   │   ├── commands/           # Command objects
│   │   ├── queries/            # Query objects
│   │   └── dto/                # Data Transfer Objects
│   ├── infrastructure/         # Infrastructure Layer - External concerns
│   │   ├── persistence/        # Repository implementations
│   │   ├── generators/         # SQL generation engines
│   │   ├── templates/          # Template engines
│   │   └── middleware/         # HTTP middleware
│   └── interfaces/             # Interface Layer - External interfaces
│       ├── web/                # Web interfaces (REST API, Static files)
│       ├── cli/                # Command-line interface
│       └── graphql/            # GraphQL interface
├── config/                     # Configuration files
├── docs/                       # Documentation
└── docker-compose.yml          # Development services
```

## Quick Start

### Prerequisites
- Go 1.22.5+
- Neo4j 4.4+ (Docker recommended)
- Modern web browser

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd neo4j-sql-generator

# Install dependencies
go mod tidy

# Start Neo4j (using Docker)
docker-compose up -d neo4j

# Run the application
go run cmd/main.go

# Access the web interface
open http://localhost:8080
```

## Usage

### Web Interface
1. Open `http://localhost:8080` in your browser
2. Use the visual designer to create your graph schema
3. Configure node types, relationships, and properties
4. Preview the generated SQL in real-time
5. Download the complete DDL scripts

### CLI Interface
```bash
# Analyze existing Neo4j database
./neo4j-sql-generator analyze --uri bolt://localhost:7687

# Generate SQL from design file
./neo4j-sql-generator generate --design ./my-design.json --target mysql

# Validate design
./neo4j-sql-generator validate --design ./my-design.json
```

### GraphQL API
```graphql
# Create a new design
mutation {
  createDesign(input: {
    name: "E-commerce Schema"
    description: "Product catalog with users and orders"
  }) {
    id
    name
  }
}

# Generate SQL schema
mutation {
  generateSchema(designId: "design-id", target: MYSQL) {
    ddl
    migrations
  }
}
```

## Configuration

### Main Configuration (`config/config.yml`)
```yaml
server:
  port: 8080
  host: localhost

neo4j:
  uri: bolt://localhost:7687
  username: neo4j
  password: password

generators:
  mysql:
    template_path: config/templates/mysql.yml
  postgresql:
    template_path: config/templates/postgresql.yml
  sqlite:
    template_path: config/templates/sqlite.yml

design:
  auto_save: true
  max_nodes: 100
  max_relationships: 200
```

## Development

### Architecture Principles

This project follows **Domain Driven Design (DDD)** principles:

- **Domain Layer**: Contains pure business logic, isolated from external concerns
- **Application Layer**: Orchestrates use cases and workflows  
- **Infrastructure Layer**: Implements technical details (databases, web servers)
- **Interface Layer**: Provides external access points (REST, GraphQL, CLI)

### Key Design Patterns
- **Ports & Adapters (Hexagonal Architecture)**
- **CQRS (Command Query Responsibility Segregation)**
- **Repository Pattern**
- **Domain Events**
- **Aggregate Pattern**

### Testing
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run integration tests
go test -tags=integration ./...
```

### Building
```bash
# Build for current platform
go build -o neo4j-sql-generator cmd/main.go

# Cross-compile for Linux
GOOS=linux GOARCH=amd64 go build -o neo4j-sql-generator-linux cmd/main.go
```

## API Documentation

### REST Endpoints
- `GET /api/designs` - List all designs
- `POST /api/designs` - Create new design
- `GET /api/designs/{id}` - Get design by ID
- `PUT /api/designs/{id}` - Update design
- `POST /api/generate` - Generate SQL from design
- `POST /api/analyze` - Analyze Neo4j database

### WebSocket Events
- `design.updated` - Design changed
- `generation.completed` - SQL generation finished
- `analysis.completed` - Database analysis finished

## Related Projects

- **[mysql-graph-visualizer](../mysql-graph-visualizer)** - Convert MySQL databases to Neo4j graphs
- Complete bidirectional ecosystem: SQL ↔ Neo4j

## Contributing

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Follow DDD principles and existing code patterns
4. Add tests for new functionality
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Roadmap

- [ ] Core DDD domain model
- [ ] Neo4j analysis engine
- [ ] Visual graph designer (Web UI)
- [ ] MySQL SQL generator
- [ ] PostgreSQL SQL generator
- [ ] SQLite SQL generator
- [ ] CLI interface
- [ ] GraphQL API
- [ ] Template customization
- [ ] Real-time collaboration
- [ ] Schema versioning
- [ ] Migration scripts generation
- [ ] Integration with mysql-graph-visualizer

---

**Status**: 🚧 In Active Development

For questions and support, please open an issue in the GitHub repository.
