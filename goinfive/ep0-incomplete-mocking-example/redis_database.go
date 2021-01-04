package ep0
import	(
	"fmt"
	"gopkg.in/redis.v3"

)

//RedisClientWrapper is a wrapper for the CLient type the gokg needs )in the import)
// which means get and set uset the param and retrun types mush
//match the the hashtable interface

type RedisClientWrapper struct {
	Client *redis.Client
}

func (i RedisClientWrapper Get(key string) ([]byte, error) {
val, ok := i.Client.Get(key).Result()
	if ok != nil {
		return nil, fmt.Errorf("Error: %s", ok)
	}
	return []byte(val),nil

}

//set method of the RedisClientWrapper type
func (i RedisClientWrapper) Set(key string, val []byte error) {
	err := i.Client.Set(key, val, 0).Err()
	return err
}
