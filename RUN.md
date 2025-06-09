# Run Guide

This project uses a `Makefile` to simplify development tasks such as setup, running the server, testing, and generating coverage reports.

## 📦 Setup

```sh
make setup
```

- Installs project dependencies using go mod tidy.
- Generates sample data (products, reviews, sellers) via data/generate_data.go.

## 🚀 Run the Application

```sh
make run
```

- Access the swagger in http://localhost:8080/swagger/index.html#/

## ✅ Run Tests

```sh
make test
```

- Executes all unit tests.
- Excludes mocks, docs, and data directories.
- Produces a coverage report in cover.out.
- +80% test coverage.

## 📊 View Coverage in Terminal

```sh
make coverage
```

- Shows a summary of code coverage in the terminal using the cover.out file.

## 🌐 View Coverage in Browser

```sh
make coverage-html
```

- Generates an HTML report from cover.out.
- Opens it in your default browser for visual inspection.

## 🧹 Clean Coverage Artifacts

```sh
make clean
```

- Deletes the cover.out file to clean up the workspace.