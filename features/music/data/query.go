package data

import (
	"context"
	"encoding/json"
	"fmt"
	"l3nmusic/app/cache"
	"l3nmusic/features/music"

	"gorm.io/gorm"
)

type musicQuery struct {
	db    *gorm.DB
	redis cache.Redis
}

func New(db *gorm.DB, redis cache.Redis) music.MusicDataInterface {
	return &musicQuery{
		db:    db,
		redis: redis,
	}
}

// Insert implements music.MusicDataInterface.
func (repo *musicQuery) Insert(userIdLogin int, input music.Core) error {
	musicInput := CoreToModel(input)
	musicInput.UserID = uint(userIdLogin)

	tx := repo.db.Create(&musicInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAll implements music.MusicDataInterface.
func (repo *musicQuery) SelectAll(ctx context.Context, page, limit int) ([]music.Core, error) {
	key := fmt.Sprintf("songs:%d:%d", page, limit)
	songsData, err := repo.redis.Get(ctx, key)
	if err == nil && songsData != "" {
		var songs []music.Core
		err = json.Unmarshal([]byte(songsData), &songs)
		if err == nil {
			return songs, nil
		}
	}

	var songs []Song
	offset := (page - 1) * limit
	result := repo.db.Offset(offset).Limit(limit).Find(&songs)
	if result.Error != nil {
		return nil, result.Error
	}

	var cores []music.Core
	for _, song := range songs {
		cores = append(cores, song.ModelToCore())
	}

	jsonData, err := json.Marshal(cores)
	if err == nil {
		songsData = string(jsonData)
		err = repo.redis.Set(ctx, key, songsData)
		if err != nil {
			return nil, err
		}
	}

	return cores, nil
}
