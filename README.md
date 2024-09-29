
# Implementation of Caching in Redis with GoFiber 

1. Clone Repository ไปที่เครื่องของคุณ

2. เข้าไปที่ Directory ของ Project และทำการ Run Docker Compose ที่เตรียมไว้ เพื่อ Deploy Service Database และ Redis
```shell
docker compose up -d
```

3. ใช้คำสั่งเพื่อ Run Project GoFiber
```shell
go run .
```

4. เริ่มการ Implement Redis กับ API Get product by id โดยทำการลง Libary สำหรับใช้งาน Redis
```shell
go get github.com/redis/go-redis/v9
```

5. สร้างไฟล์ cahce/redis.go
```go
package cache

import (
	"os"
	"strconv"

	"redis-cache-training/utils"

	"github.com/redis/go-redis/v9"
)

func RedisConnection() (*redis.Client, error) {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisConnURL, err := utils.ConnectionURLBuilder("redis")
	if err != nil {
		return nil, err
	}

	options := &redis.Options{
		Addr:     redisConnURL,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNumber,
	}

	return redis.NewClient(options), nil
}
```

6. ปรับปรุง Method GetProductByID ให้มีการใช้งาน Redis
```go
func (s *ProductServiceImpl) GetProductByID(c *fiber.Ctx, id int) (*models.Product, error) {
	logger := logging.Logger.With().Str("method", "GetProductByID").Logger()

	connRedis, err := cache.RedisConnection()
	if err != nil {
		logger.Error().Stack().Err(err).Msg("Failed to connect to Redis")
		return nil, err
	}
	cacheKey := fmt.Sprintf("product:%d", id)
	ctx := c.Context()
	var product *models.Product

	val, err := connRedis.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		product, err = s.repo.GetProductByID(c, id)
		if err != nil {
			logger.Error().Stack().Err(err).Msg("Failed to fetch from database")
			return product, err
		}

		jsonData, err := utils.CompressToJsonBytes(product)
		if err != nil {
			logger.Error().Stack().Err(err).Msg("Failed to serialize product to JSON")
			return product, errors.Errorf(errorsConstant.ERROR_CODE_FAIL_COMPRESS_JSON)
		}

		err = connRedis.Set(ctx, cacheKey, jsonData, 24*time.Hour).Err()
		if err != nil {
			logger.Error().Stack().Err(err).Msg("Failed to set product data in Redis")
			return product, errors.Errorf(errorsConstant.ERROR_CODE_FAIL_SET_DATA_REDIS)
		}

	} else if err != nil {
		logger.Error().Stack().Err(err).Msg("Error fetching product from Redis")
		return product, errors.Errorf(errorsConstant.ERROR_CODE_FAIL_FETCH_REDIS)
	} else {
		logger.Info().Msg("Success fetching product from Redis")
		err = utils.UnCompressJsonBytes([]byte(val), &product)
		if err != nil {
			logger.Error().Stack().Err(err).Msg("Failed to deserialize product from JSON")
			return product, errors.Errorf(errorsConstant.ERROR_CODE_FAIL_UNCOMPRESS_JSON)
		}
	}

	return product, nil
}
```
