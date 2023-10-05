package service

func (k *KVService) Set(key string, value string) error {
	_, err := k.processOnlyQueryRequest("POST", setPath, map[string]string{
		"key":   key,
		"value": value,
	})
	return err
}
