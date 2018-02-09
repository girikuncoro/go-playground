package bookshelf

import "fmt"

var (
	DB BookDatabase
)

func main() {
	DB = newMemoryDB()
	fmt.Println("Running DB forever...")
	for {
	}
}
