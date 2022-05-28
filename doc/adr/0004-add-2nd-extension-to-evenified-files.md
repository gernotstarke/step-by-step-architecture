# 4. add 2nd extension to evenified files

Date: 2022-05-27

## Status

Accepted

## Context

When evenifying (adding a single page to files with odd pagecount) we need a new filename to ensure the original file is not altered.

## Decision

Several options for the newly created file:

1. create in different directory: seems to complicated, as we need a second directory
2. add pre- or suffix to filename, e.g. sample.pdf -> sample-even.pdf
3. add 2nd extension, e.g. sample.pdf -> sample.evenified.pdf

diagrams.net (formerly draw.io) is using approach 3 for their png export with embedded source, which we really appreciate.

## Consequences

* Determining the target filename requires parsing the original filename. Maybe some golang standard functions will help.
* 