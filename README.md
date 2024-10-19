# Boot Progress Decoder

## Overview

The Boot Progress Decoder is a Go application designed to decode UEFI boot progress and error codes from EDK2.
It interprets hexadecimal progress/error code string and provides human-readable descriptions for class, subclass, and operation.

## Supported Formats

The decoder supports the following formats:

- **Progress Code**: `PROGRESS CODE: V03020003 I0`
- **Error Code**: `ERROR: C000000002:V03058002 I0 6D33944A-EC75-4855-A54D-809C75241F6C`

## Features

- Decodes UEFI boot progress codes and error codes.
- Provides detailed descriptions for each code, including class, subclass, operation, and severity.

## Build

To build the Boot Progress Decoder, ensure you have Go installed on your machine. Then, clone the repository and build the application:

```
go build -o bpd cmds/bpd/main.go
```

## Download

You can download the latest version of the Boot Progress Decoder from the
[Releases](https://github.com/nhivp/boot-progress-decoder/releases) page.

## Usage

To use the Boot Progress Decoder, run the application with a progress or error code as an argument:

```
./bpd "PROGRESS CODE: V03020003 I0"
PROGRESS CODE: V03020003 I0
Class     :  Software
Subclass  :  PEI Core
Operation :  Init End
```

or

```
./bpd "ERROR: C000000002:V03058002 I0 6D33944A-EC75-4855-A54D-809C75241F6C"
C000000002:V03058002 I0 6D33944A-EC75-4855-A54D-809C75241F6C
Severity  :
Class     :  Software
Subclass  :  DXE Boot Driver
Operation :  OEM Specific Error Code
Module    :  6D33944A-EC75-4855-A54D-809C75241F6C
```

If you need help with the usage, you can run the application without arguments:

```
./bpd
```

## Contributing

Contributions are welcome! If you have suggestions for improvements
or find bugs, please open an issue or submit a pull request.

## License

This project is licensed under the BSD-3-Clause License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

The Status Code Type/Value format is based on the UEFI Specification,
as implemented in the file
[edk2/MdePkg/Include/Pi/PiStatusCode.h](https://github.com/tianocore/edk2/blob/master/MdePkg/Include/Pi/PiStatusCode.h)
