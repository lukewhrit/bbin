# BBin

Experimental self-hosted pastebin service written in Golang using Chi and Badger.

BBin is entirely self-contained and requires no external services or dependencies. It makes use of Badger &mdash; an embeddable, persistent and fast key-value (KV) database.

## Installation

```sh
# Clone repository from git remote
$ git clone https://github.com/lukewhrit/bbin.git
$ cd bbin

# Build and run docker image on port 80
$ sudo docker build -t bbin .
$ sudo docker run -d \
	--restart=always \
	-e BBIN_HOSTNAME=0.0.0.0 \
	-e BBIN_PORT=80 \
	-p 80:80 \
	bbin
```

## Usage

### Configuration

You configure BBin solely via environment variables:

* **`BBIN_HOSTNAME`** controls the hostname BBin serves on.
* **`BBIN_PORT`** controls the port BBin serves on.
* `BBIN_DB_LOCATION` controls the location of the Badger database directory. Default: `/tmp/badger`.
* `BBIN_ID_LENGTH` dictates the length of document IDs. Default: `8`.
* `BBIN_MAX_PAYLOAD_SIZE` dictates the maximum length of request entities in [mebibytes](https://en.wikipedia.org/wiki/Mebibyte). Default: `256 MiB`.

Bolded keys are required.

## API Endpoints

## Contributing

1. [Fork it](https://lukewhrit/bbin/fork).
2. Create your feature branch (`git checkout -b my-new-feature`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin my-new-feature`).
5. Create a new [Pull Request](https://github.com/lukewhrit/bbin/pulls).

## Authors

* [Luke Whrit <lukewhrit@pm.me>](https://github.com/lukewhrit) - Original developer and maintainer.

## License

BBin is provided to the public under the specific terms of the [MIT license](LICENSE).
