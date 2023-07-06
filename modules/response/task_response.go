package response

import "time"

type TaskResponse struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Date        time.Time `json:"date"`
}

type RedisResponse struct {
	Key        string
	CacheValue string
}

type PayloadPaginationType struct {
	Payload []TaskResponse `json:"payload"`
	Count   int64          `json:"count"`
	Limit   int            `json:"limit"`
	Page    int            `json:"page"`
}
