package bookshelf

var (
	DB BookDatabase
)

func init() {
	DB = newMemoryDB()
}
