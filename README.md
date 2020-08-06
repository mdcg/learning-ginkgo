# learning-ginkgo

## Installing

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

## Bootstrapping a Suite

To write Ginkgo tests for a package you must first bootstrap a Ginkgo test suite. Say you have a package named `books`:

```
$ cd path/to/books
$ ginkgo bootstrap
```

## Adding Specs to a Suite

Example to generate `book_test.go`:

```
$ ginkgo generate book
```

# Running Tests

```
$ ginkgo
```