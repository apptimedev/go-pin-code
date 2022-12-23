# Go Pin Code Generator
## foruscommunity/go-pin-code

[![Forus App](https://forus.app/icons/icon-128x128.png)](https://forus.app)

---

Package `foruscommunity/go-pin-code` implements a 6-digit pin code generator.

---

* [Install](#install)
* [Examples](#examples)
* [License](#license)

---

## Install

```sh
go get -u github.com/foruscommunity/go-pin-code
```

## Examples

Generate a 6-digit pin code:

```go
import (
        "fmt"
        "github.com/foruscommunity/go-pin-code"
)

func main() {
        c, s := pincode.GenSix()
      
        fmt.Printf("int: %d\n", c)
        fmt.Printf("string: %s\n", s)
}
```

## License

BSD licensed. See the LICENSE file for details.
