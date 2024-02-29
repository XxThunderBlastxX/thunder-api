package repository

import (
	"errors"
	"github.com/XxThunderBlast/thunder-api/domain"
	"github.com/XxThunderBlast/thunder-api/types"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type kvRepository struct {
	BaseUrl string
	AuthKey string
}

func NewKVRepository(baseUrl string, authKey string) domain.KVRepository {
	return &kvRepository{
		BaseUrl: baseUrl,
		AuthKey: authKey,
	}
}

func (kv *kvRepository) GetValue(key string) (string, error) {
	reqURl := kv.BaseUrl + "/values/" + key

	agent := fiber.Get(reqURl)
	agent.Set("Authorization", "Bearer "+kv.AuthKey)
	agent.ContentType("application/json")

	statusCode, body, err := agent.Bytes()

	if err != nil {
		return "", err[0]
	}

	if statusCode != 200 {
		return "", errors.New("key not found")
	}

	return string(body), nil
}

func (kv *kvRepository) SetKeyValue(key string, value string) error {
	var res map[string]interface{}

	reqURl := kv.BaseUrl + "bulk"

	agent := fiber.Put(reqURl)
	agent.Set("Authorization", "Bearer "+kv.AuthKey)
	agent.ContentType("application/json")

	if reqBody, err := json.Marshal(types.KVRequest{Base64: false, Key: key, Value: value}); err != nil {
		return err
	} else {
		agent.Body(reqBody)
	}

	statusCode, body, err := agent.Bytes()

	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	if err != nil {
		return err[0]
	}

	if statusCode != 200 {
		errMsg := res["errors"].(map[string]interface{})["message"]
		return errors.New(errMsg.(string))
	}

	return nil
}

func (kv *kvRepository) ListKeys() (keys []string, err error) {
	var res map[string]interface{}
	reqURl := kv.BaseUrl + "/keys"

	agent := fiber.Get(reqURl)
	agent.Set("Authorization", "Bearer "+kv.AuthKey)
	agent.ContentType("application/json")

	statusCode, body, resErr := agent.Bytes()

	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if resErr != nil {
		return nil, resErr[0]
	}

	if statusCode != 200 {
		errMsg := res["errors"].(map[string]interface{})["message"]
		return nil, errors.New(errMsg.(string))
	}

	results := res["result"].([]map[string]interface{})
	for _, result := range results {
		keys = append(keys, result["name"].(string))
	}

	return keys, nil
}

func (kv *kvRepository) DeleteKey(key string) error {
	var res map[string]interface{}
	reqURl := kv.BaseUrl + "/values/" + key

	agent := fiber.Delete(reqURl)
	agent.Set("Authorization", "Bearer "+kv.AuthKey)
	agent.ContentType("application/json")

	statusCode, body, err := agent.Bytes()

	if err := json.Unmarshal(body, &res); err != nil {
		return err
	}

	if err != nil {
		return err[0]
	}

	if statusCode != 200 {
		errMsg := res["errors"].(map[string]interface{})["message"]
		return errors.New(errMsg.(string))
	}

	return nil
}
