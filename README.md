#Go Mixmax
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/ghmeier/go-mixmax) [![Go project version](https://badge.fury.io/go/github.com%2Fghmeier%2Fgo-mixmax.svg)](https://badge.fury.io/go/github.com%2Fghmeier%2Fgo-mixmax) [![Go Report Card](https://goreportcard.com/badge/github.com/ghmeier/go-mixmax)](https://goreportcard.com/report/github.com/ghmeier/go-mixmax)

##Summary

A golang Mixmax client library.

##Installation

```
go get github.com/ghmeier/go-mixmax
```

##Documentation

Check out the [Mixmax API documentation](http://developer.mixmax.com/docs/) for an exhaustive list of examples and explanations.

View the [GoDoc](http://godoc.org/github.com/ghmeier/go-mixmax) for implementation details on go-mixmax

Here are some simple examples:

###Contacts

```
mixmax := mixmax.New(apiKey)
params := &models.ContactParams{
    Email: "test@test.com",
    Name:  "test",
}

err := mixmax.Contacts.New(params)
```

###Send

```
mixmax := mixmax.New(apiKey)
params := &models.Send{
    To:      []*Recipient{&Recipient{Email: "test@test.com"}},
    From:    "me@test.com",
    Subject: "Welcome to the Mixmax API",
    HTML:    "The body of the email",
}

err := mixmax.Send.New(params)
```

##Usage

Each major API resource is contained under the main mixmax client, so usage requires initializing the mixmax client with your [Mixmax API Token](http://developer.mixmax.com/docs/getting-started-with-the-api) and importing models that will be used to send requests. This example uses the `Contact` resource, but other resources follow the same pattern:

```
import (
    "github.com/ghmeier/go-mixmax"
    mixmodels "github.com/ghmeier/go-mixmax/models"
)

mm := mixmax.New(MIXMAX_API_TOKEN)

err := mm.Contact.New(mixmodels.ContactParams)

contacts, err := mm.Contact.List()
```


##Development

Go Mixmax is still in early development with changes likely to come. If you spot problems or missing functionality [open an issue](https://github.com/ghmeier/go-mixmax/issues/new) or [submit a pull request](https://github.com/ghmeier/go-mixmax/compare).

_Note: The /insightsreport resource is not yet implemented_
