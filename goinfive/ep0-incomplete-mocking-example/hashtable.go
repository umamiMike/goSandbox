package ep0

import "errors"

var (
	ErrNotFound = errors.New("not found")
)

//hahtable is the interface fora simple hash table.  It is designed int such a way
//that some libraries immediately adhere to it with no extra code.
//on such librarar is godoc.org/github.com.hoisie/redis
type HashTable interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
