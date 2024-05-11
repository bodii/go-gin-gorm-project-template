package bootstrap

import (
	"fmt"
	"template-project-name/internal/types"
	"template-project-name/internal/utils"
	"log"
	"path"
	"sync"

	"context"

	"github.com/redis/go-redis/v9"
)

var (
	_redisInitOnce sync.Once
	redisDBs       map[string]*types.RedisDBT
)

// redisList config struct type
type redisListConfT struct {
	Servs []redisConfT `toml:"redis-server"`
}

// redis config struct type
type redisConfT struct {
	ServerName string `toml:"server_name"`
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	Auth       string `toml:"auth"`
	Db         int    `toml:"db"`
}

func RedisOnceInit() map[string]*types.RedisDBT {

	_redisInitOnce.Do(initRedis)

	return redisDBs
}

// read  cache.yaml config and set var
func loadRedisListConfig() redisListConfT {
	confPath := path.Join(CONFIG_DIR, "redis.toml")
	return utils.ReadTomlConfig[redisListConfT](confPath)
}

func initRedis() {
	// loading redis config info
	confList := loadRedisListConfig()

	// fmt.Printf("%#v\n", redisList)
	for _, db := range confList.Servs {
		RedisClient := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", db.Host, db.Port),
			Password: db.Auth,
			DB:       db.Db,
		})

		RedisContext := context.Background()
		ping, err := RedisClient.Ping(RedisContext).Result()
		if err != nil || ping != "PONG" {
			RedisClient.Close()
			log.Fatalf("redis connent error: %v", err)
		}

		redisDBs[db.ServerName] = &types.RedisDBT{DB: RedisClient, Ctx: RedisContext}
	}

	log.Println("init redis db success!")
}

func GetRedis(ServerName string) (*types.RedisDBT, error) {
	db, ok := redisDBs[ServerName]
	if !ok {
		return nil, fmt.Errorf(
			"the currently specified redis server name '%s' does not exist",
			ServerName)
	}

	return db, nil
}

func GetRedisAllServerName() []string {
	all := []string{}
	for serverName := range redisDBs {
		all = append(all, serverName)
	}

	return all
}

func GetRedisAllDB() map[string]*types.RedisDBT {
	return redisDBs
}
