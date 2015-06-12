klikin-labs/booking
==============
Booking API microservice

## Basic Usage

- use `npm install` to download dependencies
- use `npm test` to run unit and E2E tests
- use `npm start` to run the service

Check `package.json` and `Makefile` to see available scripts and build targets. NPM run scripts are mostly aliases to their Make counterparts.

## Docker

### Base Docker Image
- [klikindockerhub/iojs:onbuild](https://registry.hub.docker.com/klikindockerhub/iojs/)

This service is available in DockerHub as [klikindockerhub/booking](https://registry.hub.docker.com/u/klikindockerhub/booking/).

Remember you can pass environment variables when running your container, eg:

```
$ docker run -e NODE_ENV=dev -p 80:8080 klikindockerhub/booking

$ docker run -e NODE_ENV=prod klikindockerhub/booking
```
