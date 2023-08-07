package database

import (
	"encoding/json"
	"server/models"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const usersCacheKey = "users"

func SetCachedUsers(users []*models.User) error {
	cacheDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",              
		DB:       0,             
	})

	_, err := cacheDB.Ping().Result()
	if err != nil {
		log.Error("Error connecting to cache:", err)
		return err
	}

	jsonUserData, err := json.Marshal(users)
	if err != nil {
		log.Error("Error marshalling users for cache:", err)
		return err
	}

	log.Info("Caching users")
	return cacheDB.Set(usersCacheKey, jsonUserData, 0).Err()
}

func GetCachedUsers() ([]*models.User, error) {
	cacheDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",              
		DB:       0,              
	})

	_, err := cacheDB.Ping().Result()
	if err != nil {
		log.Error("Error connecting to cache:", err)
		return nil, err
	}
	
	jsonUserData, err := cacheDB.Get(usersCacheKey).Result()
	if err != nil {
		if err == redis.Nil {
			log.Info("Users not found in cache")
			return nil, nil
		}

		log.Error("Error getting users from cache:", err)
		return nil, err
	}

	var users []*models.User

	err = json.Unmarshal([]byte(jsonUserData), &users)
	if err != nil {
		log.Error("Error unmarshalling users from cache:", err)
		return nil, err
	}

	log.Info("Fetching users from cache")
	return users, nil
}

func ClearCache() error {
	cacheDB := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := cacheDB.Ping().Result()
	if err != nil {
		return err
	}

	err = cacheDB.Del(usersCacheKey).Err()
	if err != nil {
		return err
	}

	log.Info("Cache cleared successfully")
	return nil
}