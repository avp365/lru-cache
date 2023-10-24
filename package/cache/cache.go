package cache

import (
	"errors"
	"fmt"
	"unsafe"
)

type LRUCacheInterface interface {
	// Добавляет новое значение с ключом в кеш (с наивысшим приоритетом), возвращает true, если все прошло успешно
	// В случае дублирования ключа вернуть false
	// В случае превышения размера - вытесняется наименее приоритетный элемент
	Add(key, value string) bool

	// Возвращает значение под ключом и флаг его наличия в кеше
	// В случае наличия в кеше элемента повышает его приоритет
	Get(key string) (value string, ok bool)

	// Удаляет элемент из кеша, в случае успеха возврашает true, в случае отсутствия элемента - false
	Remove(key string) (ok bool)
}

type LRUCache struct {
	Cap      int
	CacheMap map[string]*[]string
	Cache    [][]string
}

func NewLRUCache(n int) (*LRUCache, error) {

	if n < 1 {
		return nil, errors.New("n must equal 1 or more!")
	}
	return &LRUCache{n, make(map[string]*[]string), make([][]string, 0, n)}, nil
}

func (с *LRUCache) Add(key, value string) bool {

	if _, ok := с.CacheMap[key]; ok == false {

		с.Cache = append(с.Cache, []string{key, value})
		с.CacheMap[key] = &с.Cache[(len(с.Cache) - 1)]

		if len(с.Cache) > с.Cap {
			keyDelete := с.Cache[0][0]
			с.Remove(keyDelete)
			с.Cache = с.Cache[1:]
		}

		return true
	}

	return false
}
func (с *LRUCache) Get(key string) (value string, ok bool) {

	if _, ok := с.CacheMap[key]; ok == true {
		i := с.getIndexByPtr(с.CacheMap[key])

		с.Cache = append(с.Cache, с.Cache[i])
		с.Cache = append(с.Cache[:i], с.Cache[i+1:]...)
		return с.Cache[i][1], true
	}

	return "", false

}
func (с *LRUCache) Remove(key string) (ok bool) {

	if _, ok := с.CacheMap[key]; ok == true {

		i := с.getIndexByPtr(с.CacheMap[key])
		с.Cache = append(с.Cache[:i], с.Cache[i+1:]...)
		delete(с.CacheMap, key)

		return true
	}

	return false
}

func (с *LRUCache) getIndexByPtr(ptr *[]string) int {
	size := unsafe.Sizeof(с.Cache[0])
	offset := uintptr(unsafe.Pointer(ptr)) - uintptr(unsafe.Pointer(&с.Cache[0]))

	fmt.Println(offset / size)
	return 0
}
