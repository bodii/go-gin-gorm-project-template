# Project go-gin-gorm-project-template

template for golang project

## Getting Started
template for golang project

### 1.step: Clone project
```bash
git clone https://github.com/bodii/go-gin-gorm-project-template
```

### 2.step: Init project
```bash
bash ./init_project.sh
```

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