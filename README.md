# Loopback

[![](https://godoc.org/github.com/teemow/loopback?status.svg)](http://godoc.org/github.com/teemow/loopback) [![](https://img.shields.io/docker/pulls/teemow/loopback.svg)](http://hub.docker.com/teemow/loopback) [![IRC Channel](https://img.shields.io/badge/irc-%23giantswarm-blue.svg)](https://kiwiirc.com/client/irc.freenode.net/#giantswarm)

In case you don't have a spare disk or partition but would like to play around with eg btrfs this tool creates a loopback device from a file which you can mount whereever you want.

Use cases:
 * systemd machinectl
 * conair (https://github.com/teemow/conair)

## Prerequisites

## Getting Loopback

Download the latest release from here: https://github.com/teemow/loopback/releases/latest

Clone the latest git repository version from here: https://github.com/teemow/loopback.git

### How to build

#### Build Dependencies

 * `make`
 * `docker`

#### Building the standard way

```
make && sudo make install
```

## Usage

You can create a mounted loopback device with 15G size like this:
```
loopback create --fs=ext4 --name=loopback-test size=15
```

Whenever you are done with working in the device you can unmount it like this:
```
loopback umount --name=loopback-test
```

When you don't need the device and image anymore you can destroy it like this:
```
loopback destroy --name=loopback-test
```

### Using with Docker containers

You can also use loopback in Docker containers:
```
docker run \
    --rm \
    -it \
    --privileged \
    -v /:/host:shared \
    teemow/loopback \
    create --fs=ext4 --image-path=/host/var/lib/loopback --name=loopback-test
```

If you are not yet using Docker 1.10 you can use `nsenter` which is shipped with
the image:
```
docker run \
    --rm \
    -it \
    --privileged \
    -v /:/host \
    --entrypoint=/bin/bash \
    teemow/loopback \
    -c "/opt/loopback attach --fs=ext4 --image-path=/host/var/lib/loopback --name=loopback-test --create-missing-image && /opt/loopback list --no-headings | grep loopback-test.img | awk '{print \$1}' | xargs -I{} nsenter --mount=/host/proc/1/ns/mnt -- mount {} /mnt"
```

## Dependencies

 * `dd`
 * `losetup`
 * `mkfs` (eg for btrfs)
 * `mount`

Create a btrfs fs for machined

```
sudo loopback create --name=machined --size=10 --mount-path=/var/lib/machines
```

Remove the loopbacks

```
sudo loopback destroy --name=machined
```

## Further Steps

Check more detailed documentation: [docs](docs)

Check code documentation: [godoc](https://godoc.org/github.com/teemow/loopback)

## Contact

- Mailing list: [giantswarm](https://groups.google.com/forum/!forum/giantswarm)
- IRC: #[giantswarm](irc://irc.freenode.org:6667/#giantswarm) on freenode.org
- Bugs: [issues](https://github.com/teemow/loopback/issues)

## Contributing & Reporting Bugs

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches, the contribution workflow as well as reporting bugs.

## License

Loopback is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.
