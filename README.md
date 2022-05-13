# An XDR to Go compiler

For [Buildbarn Remote Execution](https://github.com/buildbarn/bb-remote-execution)
we have implemented a userspace [NFSv4](https://en.wikipedia.org/wiki/Network_File_System)
server. NFS makes use of [XDR (External Data Representation)](https://en.wikipedia.org/wiki/External_Data_Representation)
to describe its wire format. The XDR description for NFSv4 can be found
in [RFC 7531](https://datatracker.ietf.org/doc/html/rfc7531).

This repository provides a compiler that is capable of converting the
XDR description of NFSv4 to Go code. For each data type in the XDR
description, a Go language equivalent is created. Each data type in Go
implements the `io.ReaderFrom` and `io.WriterTo` interfaces, allowing
instances of the data type to be deserialized and serialized,
respectively.

Furthermore, the XDR to Go compiler is capable of emitting interfaces
for ONC RPCv2 ("Sun RPC") program definitions, as described in
[RFC 5531](https://datatracker.ietf.org/doc/html/rfc5531). A simple RPC
server implementation is also provided.

## Grammar accepted by the XDR to Go compiler

The grammar supported by the XDR to Go compiler should match with what's
described in RFCs 4506 and 5531. Minor extensions have been made:

- Every XDR description file should start with a `package` statement,
  containing the full name of the resulting Go package.
- After the `package` statement, there may be `import` statements. These
  can be used to express dependencies between XDR description files,
  similar to Protobuf schema files can also depend on each other.
- As XDR's string type only permits ASCII characters, an `utf8string`
  type has been added for UTF-8.

Please refer to the `pkg/compiler/tests` and `pkg/protocols` directories
for examples.

## Bundled protocols

The `pkg/protocols` directory includes pregenerated copies of commonly
used XDR descriptions, such as ONC RPCv2 and NFSv4. Also included is a
copy of the NFS mount arguments protocol used by macOS.

## Why not use...

There are already a couple of other XDR libraries for Go. Examples
include:

- [github.com/davecgh/go-xdr](https://github.com/davecgh/go-xdr)
- [github.com/stellar/go/xdr](https://github.com/stellar/go/tree/master/xdr)
- [github.com/xdrpp/goxdr](https://github.com/xdrpp/goxdr)

All of these were considered, but turned out to be unsuitable for
Buildbarn's use case for one or more of the following reasons:

- They depend on runtime reflection,
- They don't include an XDR description compiler, requiring definitions
  to be translated to Go code manually,
- They don't support ONC RPCv2 `program` definitions,
- They can't compile XDR descriptions that depend on definitions coming
  from other protocols (e.g., NFSv4's dependency on ONC RPCv2).

## License

The XDR to Go compiler is Apache v2 licensed, just like other Buildbarn
components. **Protocols that are bundled under `pkg/protocols` are
subject to their original license terms.**
