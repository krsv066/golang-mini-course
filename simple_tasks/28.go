package main

type HashMap struct {
	buckets []*bucket
	size    int
}

type bucket struct {
	key   string
	value string
}

func NewHashMap() *HashMap {
	return &HashMap{
		buckets: make([]*bucket, 0),
		size:    0,
	}
}

func (hm *HashMap) Put(key string, value string) {
	index := hm.getIndex(key)
	if index == -1 {
		hm.buckets = append(hm.buckets, &bucket{key: key, value: value})
		hm.size++
	} else {
		hm.buckets[index].value = value
	}
}

func (hm *HashMap) Get(key string) (string, bool) {
	index := hm.getIndex(key)
	if index != -1 {
		return hm.buckets[index].value, true
	}

	return "", false
}

func (hm *HashMap) Remove(key string) bool {
	index := hm.getIndex(key)
	if index != -1 {
		hm.buckets = append(hm.buckets[:index], hm.buckets[index+1:]...)
		hm.size--

		return true
	}

	return false
}

func (hm *HashMap) Size() int {
	return hm.size
}

func (hm *HashMap) getIndex(key string) int {
	for i, b := range hm.buckets {
		if b.key == key {
			return i
		}
	}

	return -1
}
