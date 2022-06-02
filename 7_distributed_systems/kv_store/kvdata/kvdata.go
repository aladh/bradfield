package kvdata

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const filename = "kvdata.json"

type KVData struct {
	Data map[string]string `json:"data"`
	file *os.File
}

func Initialize() (*KVData, error) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, fmt.Errorf("error reading file info: %w", err)
	}

	if fi.Size() == 0 {
		return &KVData{Data: make(map[string]string), file: f}, nil
	} else {
		fileBytes, err := io.ReadAll(f)
		if err != nil {
			return nil, fmt.Errorf("error reading file: %w", err)
		}

		var kv KVData
		err = json.Unmarshal(fileBytes, &kv)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling file: %w", err)
		}

		kv.file = f

		return &kv, nil
	}
}

func (kv *KVData) Get(key string) string {
	return kv.Data[key]
}

func (kv *KVData) Set(key string, value string) error {
	kv.Data[key] = value

	err := kv.persist()
	if err != nil {
		return fmt.Errorf("error persisting kvdata: %s", err)
	}

	return nil
}

func (kv *KVData) persist() error {
	jsonData, err := json.Marshal(kv)
	if err != nil {
		return fmt.Errorf("error marshalling json: %s", err)
	}

	err = kv.file.Truncate(0)
	if err != nil {
		return fmt.Errorf("error truncating file: %s", err)
	}

	_, err = kv.file.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("error seeking file: %s", err)
	}

	_, err = kv.file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err)
	}

	return nil
}
