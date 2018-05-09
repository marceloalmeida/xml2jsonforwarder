# xml2jsonforwarder

Simple webserver that is listening for webhooks on XML format and converts it to JSON format.

[![Build Status](https://travis-ci.org/marceloalmeida/xml2jsonforwarder.svg?branch=master)](https://travis-ci.org/marceloalmeida/xml2jsonforwarder)

## Installation

If you are using Go 1.6+ (or 1.5 with the `GO15VENDOREXPERIMENT=1` environment variable), you can install `xml2jsonforwarder` with the following command:

```bash
$ go get -u github.com/marceloalmeida/xml2jsonforwarder
```

## Usage

```bash
$ ./xml2jsonforwarder --help
Usage of ./xml2jsonforwarder:
  -addr="0.0.0.0:8080": IP/port for the HTTP server
  -forwarder-url="": URL to which payload should be forwarded
  -return-response-body="": Return response body for debug purposes.

$ ./xml2jsonforwarder -forwarder-url "http://127.0.0.1:80" -addr "127.0.0.1:8080"
```
## Contributing

All contributions are welcome, but if you are considering significant changes, please open an issue beforehand and discuss it with us.

## License

MIT. See the `LICENSE` file for more information.
