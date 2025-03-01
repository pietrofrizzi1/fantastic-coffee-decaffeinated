package database

import (
	"database/sql"
)

func (db *appdbimpl) SetFollow(f Follow) (Follow, error) {
	_, err := db.c.Exec(`INSERT INTO followers (Id, followerId, userId) VALUES (?, ?, ?)`, f.FollowId, f.FollowedId, f.UserId)
	if err != nil {
		return f, err
	}
	return f, nil
}

func (db *appdbimpl) RemoveFollow(UserId uint64, FollowedId uint64) error {
	res, err := db.c.Exec(`DELETE FROM followers WHERE followerId=? AND userId=? `, FollowedId, UserId)
	if err != nil {
		return err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return nil
	}
	return nil
}

func (db *appdbimpl) GetFollowingId(user1 uint64, user2 uint64) (Follow, error) {
	var follow Follow
	if err := db.c.QueryRow(`SELECT Id, followerId, userId FROM followers WHERE followerId=? AND userId = ?`, user1, user2).Scan(&follow.FollowId, &follow.FollowedId, &follow.UserId); err != nil {
		if err == sql.ErrNoRows {
			return follow, ErrLikeNotFound
		}
	}
	return follow, nil
}

func (db *appdbimpl) GetFollowers(u User, token uint64) (Follow, error) {
	var follow Follow
	if err := db.c.QueryRow(`SELECT Id, followerId, userId FROM followers WHERE followerId=? AND userId = ?`, u.Id, token).Scan(&follow.FollowId, &follow.FollowedId, &follow.UserId); err != nil {
		if err == sql.ErrNoRows {
			return follow, ErrLikeNotFound
		}
	}
	return follow, nil
}

func (db *appdbimpl) GetFollowersCount(id uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM followers WHERE followerId = ?`, id).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeNotFound
		}
	}
	return count, nil
}

func (db *appdbimpl) GetFollowingsCount(id uint64) (int, error) {
	var count int
	if err := db.c.QueryRow(`SELECT COUNT(*) FROM followers WHERE userId = ?`, id).Scan(&count); err != nil {
		if err == sql.ErrNoRows {
			return count, ErrLikeNotFound
		}
	}
	return count, nil
}
