# resizer [![Build Status](https://travis-ci.org/ssola/resizer.svg)](https://travis-ci.org/ssola/resizer)

This is a naive approach to build an image resizing service. At the moment given few parameters the system returns the image resized.

At the moment this service supports those versions of Go:

- 1.3
- 1.4
- latest stable version

#### How it works?

By now it listen automatically to port 8080 by default (this should be changed in the near future). 

Resizing endpoint:

GET host:8080/resize

**Parameters**:
- image: Current image url you want to change
- width: New width of the image
- height: New image height

#### Configuration

By default we provide a dummy config.json file with some useless default values. In this configuration you can do:

- Configure default port to listen
- List of white hosts allowed to resize the image
- Max height and width for the new image

About the hosts by default all of them are restricted. But you can add as many hosts as you want and you can use regular expressions!

For example you can do something like this:

```json
{
    "hostwhitelist": [
        "([a-z]+).supercdn.com"
    ]
}
```

The previous example show you how to allow any something.supercdn.com host.

#### Dependencies

This service relies on top of some greate packages like:

- https://github.com/spf13/viper
- https://github.com/nfnt/resize

#### TODO

- [x] Resize a given image with width/height parameters
- [x] Create some unit tests
- [ ] Gopher even more this code
- [x] Configure server with configuration files
- [x] Move validators to another Go file
- [x] Allow to find hosts by regex patterns
