package book

type Book struct {
	Name string `json:"name"`
}

func NewBook(name string) *Book {
	return &Book{
		Name: name,
	}
}

func (b Book) GetName() string {
	return b.Name
}
