# Guide for project

template for golang project

## Getting Started
This is a golang project template realized by using gin web framework and gorm database model, which can be used in online version.

### 1.step: Clone project
```bash
git clone https://github.com/bodii/go-gin-gorm-project-template
```

### 2.step: Init project
```bash
bash ./init_project.sh -h
```

## 3.set the daemon on the release version
[daemon setting guide](daemon.md)

----
#### MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```