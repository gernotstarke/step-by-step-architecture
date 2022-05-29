# 3. go-package-structure

Date: 2022-05-08

## Status

Accepted

## Context

How to structure Golang code?

> "The cmd layout pattern is very useful"
> 
> quoted from [1]
> 
> 
## Decision

* `cmd` package for main application
* `internal` package to contain other go modules

resulting in:

````
.
├── cmd
├── doc
├── go.mod
├── go.sum
├── internal
└── resources
````


See these references:
* [1] [Eli Benderskys](https://github.com/eliben/modlib) overview of this approach.
* [2] [this article](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2) for an introduction.

## Consequences
