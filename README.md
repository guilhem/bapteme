# bapteme

[![Build Status](https://travis-ci.org/optiflows/bapteme.png)](https://travis-ci.org/optiflows/bapteme)
[![Gobuild Download](http://gobuild.io/badge/github.com/optiflows/bapteme/download.png)](http://gobuild.io/github.com/optiflows/bapteme)
[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/optiflows/bapteme)

## Description

`bapteme` is a go http server who can be used to named server with a random name.

## Use case

`Bapteme` is mainly used in conjonction with "static" deployment script (preseed, bash, chef, puppet, whatever).

## Usage

### Server

```shell
$ ./bapteme -h
Usage of ./bapteme:
  -bind="": Address to bind. Format IP:PORT
  -d=false: turn on debug info
  -size=10: Default final hostname size
```

### Client

|HTTP Parameter |Description |default |
|---------------|------------|--------|
|size     | size of hostname | `size` server command line option |
|prefix   | prefix to add on hostname | if User-Agent:<br>`win` for windows<br>`lin` for linux<br> else `srv` |
|instance | service to add after prefix | |
|id       | Unique identifier to generate same hostname<br>Can be MAC address etc. | |

## Example

### Server
```shell
$ ./bapteme -bind="127.0.0.1:8086" -size="15"
2014/08/28 16:32:41 INFO Bind to 127.0.0.1:8086
```

### Client
```shell
$ curl -X GET 'http://127.0.0.1:8086?size=42'
srvIuKHEt3Kg5ES88kWx8WDCplhzvBM2xEkfRWcNXE
$ curl -X GET 'http://127.0.0.1:8086?size=42'
srvBzi2mhH6y4fKOWXEH9v7EwpYMM4NesXPHY1tseW
```

```shell
$ curl -X GET 'http://127.0.0.1:8086?prefix=rhel&instance=bind&id=01:02:03:04:05:06'
rhelbindMDE6MDI
$ curl -X GET 'http://127.0.0.1:8086?prefix=rhel&instance=bind&id=01:02:03:04:05:06'
rhelbindMDE6MDI
```

```shell
$ curl -X GET 'http://127.0.0.1:8086?prefix=AZERTY&size=2'
name too long
# HTTP 500
# len(AZERTY) >> 2
```

## Dev

We are using [`gopm`](https://github.com/gpmgo/gopm) tool to manage vendoring.

Just compile / [download](http://gobuild.io/download/github.com/gpmgo/gopm) gopm.

### Build

```shell
$ gopm build
```
That's all folks!
