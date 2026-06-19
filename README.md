# Compogo Validator

[![Go Reference](https://pkg.go.dev/badge/github.com/Compogo/validator.svg)](https://pkg.go.dev/github.com/Compogo/validator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Адаптер [go-playground/validator](https://github.com/go-playground/validator) для фреймворка [Compogo](https://github.com/Compogo/compogo).

Регистрирует валидатор в DI-контейнере для проверки структур в HTTP-хендлерах, GRPC-сервисах и других компонентах.

## Установка

```shell
go get github.com/Compogo/validator
```

## Быстрый старт

```go
package main

import (
    "github.com/Compogo/compogo"
    "github.com/Compogo/validator"
)

func main() {
    app := compogo.NewApp("myapp",
        compogo.WithComponents(&validator.Component),
    )

    if err := app.Serve(); err != nil {
        panic(err)
    }
}
```

Зависимости

* [Compogo](https://github.com/Compogo/compogo) — основной фреймворк
* [go-playground/validator](https://github.com/go-playground/validator) — библиотека валидации

## Лицензия

```plantuml
MIT License

Copyright (c) 2026 Compogo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
