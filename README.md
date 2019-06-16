# rachio

**WORK IN PROGRESS**

[![wercker status](https://app.wercker.com/status/fc566cfb19f155d3d4261e186f8142e7/s/master "wercker status")](https://app.wercker.com/project/byKey/fc566cfb19f155d3d4261e186f8142e7)

<p align="center">
  <img src="https://raw.githubusercontent.com/jsgoecke/rachio/master/images/rachio.png">
</p>

This library provides a wrapper around the API to easily query and command the a [Rachio Irrigation Controller](https://www.rachio.com) in Go.

## Library Documentation

[https://godoc.org/github.com/jsgoecke/rachio](https://godoc.org/github.com/jsgoecke/rachio)

## API Documentation

[View Rachio JSON API Documentation](https://rachio.readme.io/docs/getting-started)

This is official documentation of the Rachio JSON API. The API provides functionality to monitor and control the Rachio Irrigation Controller. The project provides both a documention of the API and a Go library for accessing it.

## Installation

```
go get github.com/jsgoecke/rachio
```

## Usage

Here's an example (more in the /examples project directory):

```go
```

## ToDo

* Implement [Device](https://rachio.readme.io/docs/publicdeviceid)
* Implement [Zone](https://rachio.readme.io/docs/publiczoneid)
* Implement [ScheduleRule](https://rachio.readme.io/docs/publicscheduleruleid)
* Implement [FlexScheduleRule](https://rachio.readme.io/docs/publicflexscheduleruleid)
* Implement [Notification](https://rachio.readme.io/docs/publicflexscheduleruleid)

## Copyright & License

Copyright (c) 2019-Present Jason Goecke. Released under the terms of the MIT license. See LICENSE for details.
