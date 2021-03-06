# Helloworld Service

This is the Helloworld service

Generated with

```
micro new --namespace=go.micro --type=service github.com/equanimityid/micro-helloworld
```

## Getting Started

- [Helloworld Service](#helloworld-service)
  - [Getting Started](#getting-started)
  - [Configuration](#configuration)
  - [Dependencies](#dependencies)
  - [Usage](#usage)

## Configuration

- FQDN: go.micro.svc.helloworld
- Type: service
- Alias: helloworld

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./helloworld_service
```

Build a docker image
```
make docker
```