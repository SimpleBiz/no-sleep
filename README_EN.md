# No-Sleep
[![license](https://img.shields.io/badge/license-apache-brightgreen.svg?style=flat)](https://github.com/SimpleBiz/no-sleep/blob/main/LICENSE)  [![Go](https://img.shields.io/badge/language-Go-blue.svg)]()

[简体中文](README.md) | [English](README_EN.md)

Sleepless tonight...

This is a simple Go program that can be used to learn the basics of Go concurrency.

More importantly, when it's running, it can bypass some global management sleep controls, effectively preventing the computer from going to sleep.

## Installation

Ensure you have the Go environment installed, recommended Go 1.19 or higher.

1. Clone this repository:
git clone [your repository address]

2. Navigate to the project directory and install dependencies:
```
cd no-sleep
go mod tidy
```

3. Run the program:
```
go run main.go
```

## Features

- Prevents sleep based on simulating device input.
- Activates the sleep prevention mechanism if there's no activity for 5 seconds.

## Questions

Please submit an issue directly in this repository.

## Contribution

Contributions to this project are welcome! Please fork this repository first, then submit a Pull Request.

## License

Apache License Version 2.0

