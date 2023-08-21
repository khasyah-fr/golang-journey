package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func main() {

	// config
	redisAddress := "localhost:6379"
	redisPassword := ""

	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       0,
	})

	// Caching usecase
	cacheKey := "visited"
	cacheValue := 1028

	err := client.Set(ctx, cacheKey, cacheValue, 10*time.Minute).Err()
	if err != nil {
		log.Fatal(err)
	}

	cachedData, err := client.Get(ctx, cacheKey).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cached Data: ", cachedData)

	// Session usecase
	sessionId := "totallyrandomid"
	sessionData := "jwt_token"

	err = client.Set(ctx, sessionId, sessionData, 24*time.Hour).Err()
	if err != nil {
		log.Fatal(err)
	}

	sessionValue, err := client.Get(ctx, sessionId).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Session value: ", sessionValue)

	// Leaderboard usecase
	leaderboardKey := "scores"
	client.ZAdd(ctx, leaderboardKey, &redis.Z{Score: 100, Member: "Alucard"})
	client.ZAdd(ctx, leaderboardKey, &redis.Z{Score: 200, Member: "Balmond"})
	client.ZAdd(ctx, leaderboardKey, &redis.Z{Score: 120, Member: "Chang'e"})
	client.ZAdd(ctx, leaderboardKey, &redis.Z{Score: 170, Member: "Diggie"})
	client.ZAdd(ctx, leaderboardKey, &redis.Z{Score: 150, Member: "Edith"})

	top, err := client.ZRevRangeWithScores(ctx, leaderboardKey, 0, 2).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top 3 heroes: ")
	for _, player := range top {
		fmt.Printf("Hero %v with the score of %v\n", player.Member, player.Score)
	}
}
