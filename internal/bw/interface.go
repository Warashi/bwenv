package bw

type Item struct {}
type SecureMemo struct {}
type Folder struct {}

type Wrapper interface {
	CreateItem(Item) error
	GetFolder(name string) (Folder, error)
	ListItemInFolder(folderID string) ([]Item, error)
}
