package cache

type GenericCache interface {
	Get(seriesID uint64) ([]interface{}, error)
	Set(seriesID uint64, setDTO interface{}) (interface{}, error)
}

type Cache struct {
	GenericCache
}

func NewCache(cache GenericCache) Cache {
	return Cache{cache}
}

func (c Cache) SetSeries(seriesID uint64, setDTO interface{}) (interface{}, error) {
	return c.GenericCache.Set(seriesID, setDTO)
}

func (c Cache) GetSeries(seriesID uint64) ([]interface{}, error) {
	return c.GenericCache.Get(seriesID)
}
