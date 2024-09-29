package utils

import (
	"encoding/json"

	"redis-cache-training/logging"
)

func CompressToJsonBytes(obj interface{}) ([]byte, error) {
	raw, err := json.Marshal(obj)
	if err != nil {
		logging.Logger.Error().Stack().Err(err).Msg("Failed to serialize struct to JSON")
		return nil, err
	}
	return raw, nil
}

func UnCompressJsonBytes(jsonData []byte, obj interface{}) error {
	err := json.Unmarshal(jsonData, obj)
	if err != nil {
		logging.Logger.Error().Stack().Err(err).Msg("Failed to deserialize JSON to struct")
		return err
	}
	return nil
}
