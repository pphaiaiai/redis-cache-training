
# การใช้งาน Caching ใน Redis กับ GoFiber

โครงการนี้แสดงวิธีการใช้งาน Redis caching ในแอปพลิเคชัน GoFiber โดยมีเป้าหมายเพื่อแคชผลลัพธ์ที่ดึงมาจากฐานข้อมูล เพื่อเพิ่มประสิทธิภาพของ API โดยเฉพาะในการดึงข้อมูลสินค้าตาม ID

## สิ่งที่ต้องเตรียมก่อนเริ่ม

- [Go](https://golang.org/doc/install) (เวอร์ชัน 1.16 ขึ้นไป)
- ติดตั้ง Docker และ Docker Compose
- Redis (จัดการผ่าน Docker Compose)
- ฐานข้อมูลที่พร้อมใช้งาน (จัดการผ่าน Docker Compose)

## ขั้นตอนการตั้งค่าและรันโปรเจกต์

### 1. Clone โปรเจกต์ไปยังเครื่องของคุณ

เริ่มต้นโดยการ clone โปรเจกต์นี้ไปยังเครื่องของคุณ:

```shell
git clone https://github.com/pphaiaiai/redis-cache-training.git
```

### 2. สร้า Service ผ่าน Docker

เข้าไปในไดเรกทอรีของโปรเจกต์และใช้คำสั่ง Docker Compose เพื่อเริ่มบริการฐานข้อมูลและ Redis:

```shell
cd redis-cache-training
docker compose up -d
```

คำสั่งนี้จะทำให้บริการรันในโหมด background

### 3. รันแอปพลิเคชัน GoFiber

หลังจากที่บริการทั้งหมดถูกเริ่มต้นแล้ว ให้รันแอปพลิเคชัน GoFiber:

```shell
go run .
```

API Server จะเริ่มทำงานบนพอร์ตที่กำหนด พร้อมให้บริการรับส่งคำขอ

### 4. ติดตั้งไลบรารี Redis สำหรับ Go

ก่อนที่จะเริ่มการใช้งาน Redis ตรวจสอบให้แน่ใจว่าได้ติดตั้งไลบรารี Redis สำหรับ Go ด้วยคำสั่งนี้:

```shell
go get github.com/redis/go-redis/v9
```

คำสั่งนี้จะเพิ่มการรองรับ Redis ในแอปพลิเคชัน GoFiber ของคุณ

### 5. สร้างการเชื่อมต่อ Redis

สร้างไฟล์การเชื่อมต่อ Redis และเพิ่มโค้ดใน `cache/redis.go` ดังนี้:

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

ฟังก์ชันนี้จะทำหน้าที่สร้างการเชื่อมต่อกับ Redis โดยใช้ค่าคอนฟิกที่ได้จาก environment variables

### 6. ปรับปรุง Method `GetProductByID` ให้มีการใช้งาน Redis

แก้ไข `GetProductByID` ใน struct `ProductServiceImpl` เพื่อใช้ Redis ในการ caching:

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
        // ดึงข้อมูลจากฐานข้อมูลและแคชผลลัพธ์
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

ใน method นี้ เราจะพยายามดึงข้อมูลจาก Redis โดยใช้ `product ID` เป็น key ถ้าไม่พบข้อมูลใน Redis จะดึงข้อมูลจากฐานข้อมูล จากนั้นจะ serialize ข้อมูลและเก็บลงใน Redis เพื่อใช้ในคำขอครั้งถัดไป

## สรุป

การนำ Redis caching มาใช้งานช่วยให้ API ทำงานได้รวดเร็วขึ้น โดยลดภาระการดึงข้อมูลซ้ำๆ จากฐานข้อมูล ซึ่งเหมาะสมอย่างยิ่งสำหรับข้อมูลที่ไม่ค่อยเปลี่ยนแปลง เช่น รายละเอียดสินค้า คุณสามารถเพิ่มฟังก์ชันอื่นๆ หรือจัดการกับกลยุทธ์การจัดการ cache ต่อไปได้ตามความต้องการของโครงการ
