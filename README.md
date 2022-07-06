# Go Pin Code Generator
## foruscommunity/gopincode

[![Forus App](https://forus.app/icons/icon-128x128.png)](https://forus.app)

---

Package `foruscommunity/gopincode` implements a 6-digit pin code generator.

---

* [Install](#install)
* [Examples](#examples)
* [License](#license)

---

## Install

```sh
go get -u github.com/foruscommunity/gopincode
```

## Examples

Generate a 6-digit pin code:

```go
import (
        "fmt"
        "github.com/foruscommunity/gopincode"
)

func main() {
        c, s := gopincode.GenSix()
      
        fmt.Printf("int: %d\n", c)
        fmt.Printf("string: %s\n", s)
}
```

## License

BSD licensed. See the LICENSE file for details.