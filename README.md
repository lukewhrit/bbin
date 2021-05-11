# Sojourner

[![GitHub last commit](https://img.shields.io/github/last-commit/lukewhrit/sojourner)](https://github.com/lukewhrit/sojourner/commits/main) [![Project License](https://img.shields.io/github/license/lukewhrit/sojourner)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/lukewhrit/sojourner)](https://goreportcard.com/report/github.com/lukewhrit/sojourner)

Self-hosted Pastebin service written in Golang using Chi and Badger. Sojourner is fully compatible with [Spacebin](https://docs.spaceb.in), allowing any Spacebin client to work perfectly with Sojourner and vice versa.

Sojourner is entirely self-contained and requires no external services or dependencies. It makes use of Badger &mdash; an embeddable, persistent and fast key-value (KV) database.

## Table of Contents

- [Sojourner](#sojourner)
  - [Table of Contents](#table-of-contents)
  - [Documentation](#documentation)
    - [Installation](#installation)
    - [Usage](#usage)
    - [Configuration](#configuration)
  - [Contributing](#contributing)
  - [Authors](#authors)
  - [Copyright and Licensing](#copyright-and-licensing)

## Documentation

### Installation

```sh
# Clone repository from git remote
$ git clone https://github.com/lukewhrit/sojourner.git
$ cd sojourner

# Build and run docker image on port 80
$ sudo docker build -t sojourner .
$ sudo docker run -d \
  --restart=always \
  -e SOJOURNER_HOSTNAME=0.0.0.0 \
  -e SOJOURNER_PORT=80 \
  -p 80:80 \
  sojourner
```

### Usage

### Configuration

You configure Sojourner solely via environment variables:

* **`SOJOURNER_HOSTNAME`** controls the hostname Sojourner serves on.
* **`SOJOURNER_PORT`** controls the port Sojourner serves on.
* `SOJOURNER_DB_LOCATION` controls the location of the Badger database directory. Default: `/tmp/badger`.
* `SOJOURNER_ID_LENGTH` dictates the length of document IDs. Default: `8`.
* `SOJOURNER_MAX_PAYLOAD_SIZE` dictates the maximum length of request entities in [mebibytes](https://en.wikipedia.org/wiki/Mebibyte). Default: `256 MiB`.

Bolded keys are required.

## Contributing

1. [Fork it](https://lukewhrit/sojourner/fork).
2. Create your feature branch (`git checkout -b my-new-feature`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin my-new-feature`).
5. Create a new [Pull Request](https://github.com/lukewhrit/sojourner/pulls).

## Authors

* [Luke Whrit <lukewhrit@pm.me>](https://github.com/lukewhrit) - Original developer and maintainer.

## Copyright and Licensing

Sojourner is provided to the public under the specific terms of the [MIT license](LICENSE).
