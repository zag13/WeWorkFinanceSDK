# WeWork Finance SDK Go Wrapper
 
Written by cursor, Assisted by zag13.

This is a Go wrapper for the WeWork Finance SDK, which allows you to access WeWork chat records and media files.

## Prerequisites

- Go 1.20 or later
- C compiler (gcc/clang)
- WeWork Finance SDK C library

## Project Structure

```
.
├── lib/
│   ├── arm64_C_sdk/    # ARM64 architecture SDK files
│   └── amd64_C_sdk/    # x86_64 architecture SDK files
├── finance_sdk.go      # Main SDK implementation
├── finance_sdk_test.go # Test cases
├── go.mod             # Go module file
└── Dockerfile         # Docker build configuration
```

## Installation

1. Clone this repository:

```bash
git clone https://github.com/zag13/WeWorkFinanceSDK.git
cd WeWorkFinanceSDK
```

2. Place the WeWork Finance SDK C library files in the appropriate architecture directory under `lib/`:
   - For ARM64: `lib/arm64_C_sdk/`
     - `libWeWorkFinanceSdk_C.so`
     - `WeWorkFinanceSdk_C.h`
   - For x86_64: `lib/amd64_C_sdk/`
     - `libWeWorkFinanceSdk_C.so`
     - `WeWorkFinanceSdk_C.h`

## Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/zag13/WeWorkFinanceSDK"
)

func main() {
    // Initialize SDK
    sdk, err := finance.NewSDK("your_corp_id", "your_secret")
    if err != nil {
        log.Fatal(err)
    }
    defer sdk.Close()

    // Get chat data
    data, err := sdk.GetChatData(0, 100, "", "", 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Chat data: %s\n", string(data))

    // Decrypt data
    decrypted, err := sdk.DecryptData("encryption_key", "encrypted_message")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Decrypted data: %s\n", string(decrypted))

    // Get media data
    media, err := sdk.GetMediaData("index_buf", "sdk_file_id", "", "", 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Media data: %+v\n", media)
}
```

## Building with Docker

You can use the provided Dockerfile to build and test the SDK in a containerized environment:

```bash
# Build for ARM64 (default)
docker build -t wework-finance-sdk .

# Build for x86_64
docker build --build-arg TARGETPLATFORM=linux/amd64 -t wework-finance-sdk .

# Run the container
docker run -it wework-finance-sdk
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
