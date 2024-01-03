package token

type Token[ID ~string, T IdGetter[ID]] interface {
	Create(T) (ID, error)
	Decode(ID) (T, error)
	GetId(ID) (ID, error)
}

type IdGetter[ID ~string] interface {
	GetId() ID
}
