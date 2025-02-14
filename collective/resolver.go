package collective

/*
type resolver interface {
	resolve(name string, version int) ([]byte, error)
}
*/

type contentResolver func(name string, version int) ([]byte, error)

type relationResolver func(name string) (relation, error)
