# laago
Go library for Windows' Large Address Aware detection and assignment.

## Installation
`go get -u github.com/karashiiro/laago`

## Usage
Please note that not all 32-bit binaries run without errors in Large Address Aware mode. Always make backups of your programs
before modifying them.

```go
// dummy.exe exists without LAA enabled
isLaa, _ := laago.IsLargeAddressAware("dummy.exe")
fmt.Println("LAA enabled: %t", isLaa) // LAA enabled: false

laago.SetLargeAddressAware("dummy.exe")
isLaa, _ = laago.IsLargeAddressAware("dummy.exe")
fmt.Println("LAA enabled: %t", isLaa) // LAA enabled: true

laago.UnsetLargeAddressAware("dummy.exe")
isLaa, _ = laago.IsLargeAddressAware("dummy.exe")
fmt.Println("LAA enabled: %t", isLaa) // LAA enabled: false
```