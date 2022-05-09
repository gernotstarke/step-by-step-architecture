# 1. Golang and Fyne-for-gui

Date: 2022-04-10

## Status

Accepted

## Context

We need a graphical user interface running on Mac-OS, Windows and Linux, based upon a single codebase.


## Decision

We use Golang as programming language and [Fyne.io](https://fyne.io) as GUI framework.

## Alternatives

Alternative choices could have been:

* C++ and QT
* Python and wxWidgets
* Java and JavaFX

## Consequences

A prototype demonstrated that it is fairly easy to compile a demo application for Mac-OS, Windows and Linux.
