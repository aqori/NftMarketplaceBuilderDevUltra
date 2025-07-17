// cmd/nftmarketplacebuilderdevultra/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplacebuilderdevultra/internal/nftmarketplacebuilderdevultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplacebuilderdevultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
