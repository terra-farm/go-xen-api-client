# Go XenAPI client library

This is a client library for the Xapi toolstack
(http://xapi-project.github.io/).

This library covers the entire XenAPI and I have successfully used it to
implement a Terraform plugin that interfaces Citrix XenServer. That being said,
this library is not production-ready yet. Use it at your own risk, and don't
expect everything in this library to work out of the box.

## Project status

The most important missing pieces before this library is production-ready are:

  - A strategy how to handle the differences in the XenAPI versions.
  - Tests, at least for the various data type conversions.
  - Embed XenAPI documentation as GoDoc in the generated code.
  - Better error messages.
  - Usage examples.

Contributions welcome!

Please note that I want to keep this library lean. I envision it to merely
provide a one-to-one mapping of XenAPI functions to Go functions. Because of
this, I will likely not accept pull requests that implement higher level
functionality.

## Implementation notes

Most of the code in this library is generated from a description of the XenAPI.
This description is the file `xenapi.json`, the source of which is the XenAPI
documentation at http://xapi-project.github.io/:

  - https://github.com/xapi-project/xapi-project.github.io/tree/master/_data

The list of error code constants in `error.go` is borrowed from xapi-projects
OCaml client:

  - https://github.com/xapi-project/xen-api/blob/master/ocaml/idl/api_errors.ml
