package service

func (k *KVService) Get(key string) (string, error) {
	res, err := k.processOnlyQueryRequest("GET", getPath, map[string]string{
		"key": key,
	})
	if err != nil {
		return "", err
	}
	ans, err := answerFromString[getAnswer](res)
	if err != nil {
		return "", err
	}
	return ans.Value, nil
}
