package entity

type Image struct {
	// Id   int
	Name string
}

// func (i *Image) GetId() int {
// 	return i.Id
// }

func (i *Image) GetName() string {
	return i.Name
}

func NewImage(name string) Image {
	return Image{
		// Id:   id,
		Name: name,
	}
}
