# 3. go-package-structure

Date: 2022-05-08

## Status

Accepted

## Context

How to structure Golang code?

## Decision

* `cmd` package for main application
* `internal` package to contain other go modules

resulting in:

````
.
├── cmd
├── docs
├── go.mod
├── go.sum
├── internal
└── resources
````


See [this article](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) for an introduction.
## Consequences

What becomes easier or more difficult to do and any risks introduced by the change that will need to be mitigated.
