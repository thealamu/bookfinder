# BookFinder
Find quick details about any book

[![Build Status](https://travis-ci.com/thealamu/bookfinder.svg?branch=dev)](https://travis-ci.com/thealamu/bookfinder)

BookFinder is a little web app that collects books from two APIs based on a search query.

## Develop
To deploy, obtain a key from [GoodReads](https://www.goodreads.com/api/keys) and export as *GREADS_KEY*

```shell
export GREADS_KEY=[your_goodreads_api_key]
```

Clone the project, build and run
```shell
git clone https://github.com/thealamu/bookfinder.git
cd bookfinder
go build cmd/bookfinder.go
./bookfinder
```
