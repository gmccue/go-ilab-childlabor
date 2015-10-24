# go-ilab-childlabor - A Go wrapper library for the ILAB Child Labor API.

[![Build Status](https://api.travis-ci.org/gmccue/go-ilab-childlabor.png?branch=master)](https://travis-ci.org/gmccue/go-ilab-childlabor)
[![GoDoc](https://godoc.org/github.com/gmccue/go-ilab-childlabor?status.svg)](https://godoc.org/github.com/gmccue/go-ilab-childlabor)

go-ilab-childlabor provides programmatic acces to the [ILAB Child Labor API](http://developer.dol.gov/others/sweat-and-toil).

## Installation
```
go get github.com/gmccue/go-ilab-childlabor
```

## Usage
In order to use this library, you must have a valid API token for the DOL Public API. You can register for an API key at the [US DOL website](https://devtools.dol.gov/developer/).

Sample usage:
```
import laborstats "github.com/gmccue/go-ilab-childlabor"

func test() {
	api := lstats.NewLaborStatsAPI("{your API token}")
	api.Debug = true

	countryData, err := api.QueryCountryData()
	if err != nil {
		log.Println(err)
	}

	log.Printf("%v", countryData)
}
```

### Configurable fields
| Field     | Type   | Description                                                            | Example |
|-----------|--------|------------------------------------------------------------------------|---------|
| Debug     | Bool   | Output detailed information related to an API request. Uses pkg `log`. | api.Debug(true)
| SecretKey | String | Your API token.                                                        | api.SecretKey("123abc")

Detailed struct field information can be found [in the wiki]().
