# resizer

This is a naive approach to build an image resizing service. At the moment given few parameters the system returns the image resized.

#### How it works?

By now it listen automatically to port 8080 by default (this should be changed in the near future). 

Resizing endpoint:

GET host:8080/resize

**Parameters**:
- image: Current image url you want to change
- width: New width of the image
- height: New image height

#### TODO

- [x] Resize a given image with width/height parameters
- [ ] Create some unit tests
- [ ] Gopher even more this code
- [ ] Add support to upload final image result to AWS S3 or similar service (creating adapters)
- [ ] Configure server with configuration files
