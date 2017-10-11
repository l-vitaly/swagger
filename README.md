# Swaggre Handler

Get swagger yaml/json file.

# Use 

``` golang
import (
  "fmt"
  "net/http"

  "github.com/l-vitaly/swagger"
)

func main() {
  http.Handle("/swagger.yaml", swagger.Handler("./path/to/swagger.yaml"))	
  fmt.Println(http.ListenAndServe(":9000", nil))
}
```