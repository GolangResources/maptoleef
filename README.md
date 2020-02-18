# maptoleef
map[string]interface{} to LEEF 2.0 string

## Example

```
package main

import (
        "log"
        "github.com/GolangResources/maptoleef/v1"
)

func main() {
        mapt := maptoleef.LEEF{
                Manufacter: "Manufacter",
                ProductName: "ProductName",
                ProductVersion: "1",
                EventID: "1",
                Delimiter: "",
        }
        msg := map[string]interface{} {
                "key": "value",
        }
        log.Println(mapt.MapToLEEF(msg))
}
```
