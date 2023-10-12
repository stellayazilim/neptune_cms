# NeptuneCms

## Description
A content management system on premise

## Table of Contents
- [Getting Started](#getting-started)
- [Prerequisites](#prerequisites)
- [Usage](#usage)
- [Commands](#commands)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites
List any software, dependencies, or tools that users need to install or configure before using your project. Make sure to specify versions if necessary.

- Docker
- Go (Golang)
- Goose (database migration tool)
- Other dependencies...

### Usage


#### 1. Clone the repository:

ssh
```shell
git clone git@github.com:stellayazilim/neptune_cms.git
cd neptun_cms
```
http
```shell
git clone https://github.com/stellayazilim/neptune_cms.git
cd neptun_cms
```
#### 2. Set up your database environment by modifying the following fields in the Makefile:

* `POSTGRES_USER`
* `POSTGRES_PASSWORD`
* `POSTGRES_DB`
* `POSTGRES_PORT`
* `POSTGRES_HOST`   



#### 3. Start enviroment

```shell
make spinup
```

#### 4. Run the migrations to set up your database:
```shell
make up
```

#### 5. Start the application in development mode with hot reload:
```shell
make dev
```

## Commands

Each of the Makefile commands available in project.

* `spinup`: Spin up the database and other services using Docker Compose.
* `spindown`: Spin down (stop and remove) the Docker containers.
* `dev`: Start the application in development mode with hot reload.
* `create`: Create a new migration file.
* `up`: Migrate the database to the most recent version.
* `up-by-one`: Migrate the database up by one version.
* `up-to`: Migrate the database to a specific version.
* `down`: Roll back the database to the previous version.
* `down-to`: Roll back the database to a specific version.
* `redo`: Re-run the latest migration.
* `reset`: Roll back all migrations.
* `status`: Show the migration status for the current database.
* `version`: Display the current version of the database.
* `fix`: Apply sequential ordering to migrations.
* `validate`: Check migration files without running them.
* `test_db`: Run tests on the database.
* `test_unit`: Run unit tests.
* `test_cov`: Run code coverage tests.


## Testing

#### 1. To run unit tests:
```shell
make test_unit
```


#### 2. To run code coverage tests and generate a coverage report:
```shell
make test_cov
```

#### 3. To test database migrations, you can use:

note: you need a fresh database instance
```shell
make test_db
```
## Contributing
Any helps are appreciated, feel free to fork this repository
## License

[AGPL-3.0](LICENSE)