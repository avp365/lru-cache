package cache

import (
	"container/list"
	"errors"
	"sync"
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
type CacheData struct {
	Key   string
	Value string
}
type LRUCache struct {
	Cap      int
	CacheMap map[string]*list.Element
	Cache    *list.List
	mutex    sync.Mutex
}

func NewLRUCache(n int) (*LRUCache, error) {

	if n < 1 {
		return nil, errors.New("n must equal 1 or more!")
	}
	return &LRUCache{Cap: n, CacheMap: make(map[string]*list.Element), Cache: list.New()}, nil
}

func (c *LRUCache) Add(key, value string) bool {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.CacheMap[key]; ok == false {

		c.CacheMap[key] = c.Cache.PushBack(CacheData{key, value})

		if c.Cache.Len() > c.Cap {

			e := c.Cache.Front()
			key := e.Value.(CacheData).Key
			c.Cache.Remove(e)
			delete(c.CacheMap, key)
		}

		return true
	}

	return false
}
func (c *LRUCache) Get(key string) (value string, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.CacheMap[key]; ok == true {
		e := c.CacheMap[key]
		c.Cache.PushBack(e)

		return e.Value.(CacheData).Value, true
	}

	return "", false

}
func (c *LRUCache) Remove(key string) (ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, ok := c.CacheMap[key]; ok == true {

		c.Cache.Remove(c.CacheMap[key])
		delete(c.CacheMap, key)

		return true
	}

	return false
}
