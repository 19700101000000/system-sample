package ark

type Auth struct {
	Name   string
	Pass   string
	Result chan<- *string
}
