# Jsonix

[中文文档](https://github.com/JsonLee12138/jsonix/blob/main/README.md)

Jsonix is a modern framework developed with Go and Fiber, designed to improve project development efficiency.

The framework provides a series of CLI tools for service management, project initialization, code generation, and automated database migration. With a modular design that supports controllers, services, repositories, and more, Jsonix offers rapid code scaffolding along with rich functional modules.

## Installation

To install Jsonix, run:

```
go install github.com/JsonLee12138/jsonix@latest
```

## Overview

- CLI Tools Support: Built-in commands for various operations.
- Web Functionality: Basic web support is provided.
- Fiber Framework Support: Integrated with the Fiber web framework.
- Apifox Integration: Automatic payload upload to Apifox.
- Swagger Generation: Built-in support for generating Swagger documentation.
- Gorm ORM: Database ORM support via Gorm.
- Redis Caching: Redis support for caching.
- i18n Support: Internationalization support.
- Logging: Robust logging functionality.
- CORS: Cross-Origin Resource Sharing support.
- Live Reload: Jsonix supports hot reloading.
- Dependency Injection: Supported via dig.
- Configuration Reading: Built-in support for reading configuration files.

The framework provides several configuration examples:
- [ApifoxConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/apifox.go)
- [CorsConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/cors.go)
- [I18nConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/i18n.go)
- [LogConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/log.go)
- [MysqlConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/mysql.go)
- [RedisConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/redis.go)
- [SwaggerConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/swagger.go)
- [SystemConfig](https://github.com/JsonLee12138/jsonix/blob/main/pkg/configs/system.go)

## CLI Usage

Running `jsonix` provides several commands; some of the primary usages are listed below:

### Server Command

```
$ jsonix --help

  Usage: jsonix server

  Options:
    -e, --env <string>                  Set the current service environment (dev, prod, test)
    -w, --watch                         Enable configuration file hot reloading
    -s, --show <string>                 Check if a specific port is in use
    -k, --kill <string>                 Kill the process occupying a specified port
    -h, --help                          Display help information

  Usage: jsonix gen

  Options:
    -m, --module <string>               Specify the module name
    -s, --service                       Generate service code
    -c, --controller                    Generate controller code
    -r, --repository                    Generate repository code
    -e, --entity                        Generate entity code
    -o, --override                      Overwrite existing files
    -h, --help                          Display help information

  Usage: jsonix migrate                 Database Migration Command

  Options:
    -r, --root <string>                 Specify the project root directory
    -d, --dest <string>                 Specify the output directory for migration files
    -h, --help                          Display help information

  Usage: jsonix init                    Initialize project

  Options:
    -h, --help                          Display help information
```

## Project File Structure

Below is an overview of the typical project structure:

```
├─ .vscode                     # Recommended VSCode configuration
├─ .run                        # GoLand run command configurations
├─ apps                        # Application modules
│  ├─ example                  # Example module
│  │  ├─ controller            # Controllers
│  │  ├─ repository            # Repositories
│  │  ├─ service               # Services
│  │  └─ entry                 # Entry files
│  └─ ...                      # Other modules
├─ auto_migrate                # Auto-migration directory (generated via 'jsonix migrate')
│  └─ ...                      # Auto-migration files
├─ config                      # Configuration file directory
│  ├─ config.yaml              # Main configuration file
│  ├─ config.dev.yaml          # Development environment configuration
│  ├─ config.test.yaml         # Testing environment configuration
│  ├─ config.prod.yaml         # Production environment configuration
│  ├─ config.local.yaml        # Local configuration file
│  ├─ config.dev.local.yaml    # Local development configuration
│  ├─ config.test.local.yaml   # Local testing configuration
│  ├─ config.prod.local.yaml   # Local production configuration
│  ├─ regexes.yaml             # UA parser configuration file
│  └─ ...                      # Other configuration files
├─ configs                     # Example configuration instances directory
│  ├─ configs.yaml             # All configuration instances
│  └─ ...                      # Other configuration instance files
├─ core                        # Core module directory
├─ docs                        # Swagger documentation directory
├─ locales                     # Localization files directory
├─ logs                        # Logs directory
├─ middleware                  # Middleware directory
├─ tmp                         # Temporary files for air startup (do not modify or commit)
├─ utils                       # Utility functions directory
├─ .air.toml                   # air configuration file
├─ main.go                     # Main application entry file
├─ go.mod                      # Go module definition
└─ go.sum                      # Go module checksum file
```

## Feedback & Suggestions

If you encounter any issues or have suggestions for new features during your usage, please feel free to provide your feedback through one of the following methods:

- Submit an Issue on GitHub
- Send an email to: lijunsong2@gmail.com
- Personal Discord: json_lee12138
