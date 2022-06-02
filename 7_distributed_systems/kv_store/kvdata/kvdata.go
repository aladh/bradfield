package kvdata

type KVData struct {
	Data map[string]string
}

func LoadOrInitialize() *KVData {
	return &KVData{
		Data: make(map[string]string),
	}
}

func (kv *KVData) Get(key string) string {
	return kv.Data[key]
}

func (kv *KVData) Set(key string, value string) {
	kv.Data[key] = value
}
