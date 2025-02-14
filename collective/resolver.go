package collective

type contentResolver func(name string, version int) ([]byte, error)

type relationResolver func(name string) (relation, error)
