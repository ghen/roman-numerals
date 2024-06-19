# Roman numerals parser (Go)

Implements a `Roman` type that supports text marshalling / unmarshalling.

## Usage

```go
  import (
    "fmt"
    
    "github.com/ghen/roman-numerals/roman"
  )

  func main() {

    var r Roman
    if err := r.UnmarshalText([]byte("XIV")); err != nil {
      panic(err)
    }

    fmt.Printf("%s = %d", r.String(), r.Value)
  }
```

## Notes

Not covered in `roman` package:
* Text tokenization (trim the whitespace before parsing the value)
* Text normalization (only upper case ASCII characters are accepted)
* Validation (roman number characters are parsed as long as they do not exceed the maximum allowed value of 3,999)