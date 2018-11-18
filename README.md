# regenea [![Build Status][travis-image]][travis-url] [![Coverage Status][coveralls-image]][coveralls-url] [![Go Report Card][report-image]][report-url] [![GoDoc][doc-image]][doc-url]
A program to manage your family tree. It uses a JSON-based format called `genea`.

# Install it

```
go get -u github.com/LouisBrunner/regenea
```

# Usage

Here is the list of supported actions:

## Check

This command checks that your family tree is correctly formatted, it will return a zero status code if it's valid and a non-zero otherwise.

Using a file:
```
regenea check tree.genea
```

Using stdin:
```
cat tree.genea | regenea check
```

## Report

This command prints a report of interesting facts and statistics that your family tree contains.

Using a file:
```
regenea report tree.genea
```

Using stdin:
```
cat tree.genea | regenea report
```

## Report

This command prints a report of interesting facts and statistics that your family tree contains.

Using a file:
```
regenea report tree.genea
```

Using stdin:
```
cat tree.genea | regenea report
```

## Transform

This command transforms a family tree from one format to another. Currently only `genea` is supported (V1 and V2).

Using files:
```
regenea transform -in treeV1.genea -inform genea -out treeV2.genea -outform genea
```

Using stdin/stdout:
```
cat treeV1.genea | regenea report -inform genea -outform genea > treeV2.genea
```

*Note:* you can use this command to upgrade/downgrade between `genea` versions.

Upgrade:
```
regenea transform -in treeV1.genea -inform genea -out treeV2.genea -outform genea
```

Downgrade (will result in loss of data):
```
regenea transform -in treeV2.genea -inform genea -out treeV1.genea -outform genea -outversion 1
```

## Display

Coming soon

# `genea` format

Coming soon

[travis-image]: https://travis-ci.org/LouisBrunner/regenea.svg?branch=master
[travis-url]: https://travis-ci.org/LouisBrunner/regenea
[coveralls-image]: https://coveralls.io/repos/github/LouisBrunner/regenea/badge.svg?branch=master
[coveralls-url]: https://coveralls.io/github/LouisBrunner/regenea?branch=master
[report-image]: https://goreportcard.com/badge/github.com/LouisBrunner/regenea?branch=master
[report-url]: https://goreportcard.com/report/github.com/LouisBrunner/regenea?branch=master
[doc-image]: https://godoc.org/github.com/LouisBrunner/regenea?status.svg?branch=master
[doc-url]: https://godoc.org/github.com/LouisBrunner/regenea?branch=master
