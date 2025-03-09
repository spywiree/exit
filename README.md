# Golang exit-with-defer

Before:

```go
package main

import (
	"log"
	"fmt"
)

func main() {
	defer fmt.Println("defer") // does not run

	err := doSomething()
	if err != nil {
		log.Fataln(err)
	}
}
```

After:

```go
package main

import (
	"github.com/spywiree/exit"
	"fmt"
)

func main() {
	defer exit.Handle()
	defer fmt.Println("defer")

	err := doSomething()
	if err != nil {
		exit.Fatalln(err)
	}
}
```
