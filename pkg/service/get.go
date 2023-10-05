package service

func (k *KVService) Get(key string) (string, error) {
	return k.processOnlyQueryRequest("GET", getPath, map[string]string{
		"key": key,
	})
}
