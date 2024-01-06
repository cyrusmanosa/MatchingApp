// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: record.sql

package db

import (
    "context"
    "fmt"
    "time"
)

func (q *Queries) CreateChatTable(ctx context.Context, tablename string) error {
    createTable := fmt.Sprintf(`CREATE TABLE %s (
        "target_id" INT UNIQUE NOT NULL,
        "role_type" VARCHAR NOT NULL,
        "media_type" VARCHAR NOT NULL,
        "message" VARCHAR,
        "media" VARCHAR,
        "isRead" BOOLEAN,
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT (now())
    );`, tablename)

    _, err := q.db.Exec(ctx, createTable)
    return err
}


type CreateRecordParams struct {
    TargetID  int32     `json:"target_id"`
    RoleType  string    `json:"role_type"`
    MediaType string    `json:"media_type"`
    Message   string    `json:"message"`
    Media     string    `json:"media"`
    Isread    bool      `json:"isread"`
}

func (q *Queries) CreateRecord(ctx context.Context, arg CreateRecordParams, tablename string) (Record, error) {
    
    createRecord := fmt.Sprintf(`-- name: CreateRecord :one
    INSERT INTO %s (
        target_id,
        role_type,
        media_type,
        message,
        media,
        isread
    ) VALUES (
        $1,$2,$3,$4,$5,$6
    ) RETURNING target_id, role_type, media_type, message, media, isread,created_at
    `,tablename)
    
    row := q.db.QueryRow(ctx, createRecord,
        arg.TargetID,
        arg.RoleType,
        arg.MediaType,
        arg.Message,
        arg.Media,
        arg.Isread,
    )
    var i Record
    err := row.Scan(
        &i.TargetID,
        &i.RoleType,
        &i.MediaType,
        &i.Message,
        &i.Media,
        &i.Isread,
        &i.CreatedAt,
    )
    return i, err
}

type DeleteRecordParams struct {
    TargetID  int32     `json:"target_id"`
    CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) DeleteRecord(ctx context.Context, arg DeleteRecordParams, tablename string) error {
    deleteRecord := fmt.Sprintf(`-- name: DeleteRecord :exec
    DELETE FROM %s
    WHERE target_id = $1
    AND created_at = $2
    `,tablename)
    
    _, err := q.db.Exec(ctx, deleteRecord, arg.TargetID, arg.CreatedAt)
    return err
}

func (q *Queries) GetRecord(ctx context.Context, targetID int32, tablename string) ([]Record, error) {  
    getRecord := fmt.Sprintf(`-- name: GetRecord :many
    SELECT target_id, role_type, media_type, message, media, isread,created_at FROM %s
    WHERE target_id = $1
    ORDER BY created_at
    `,tablename)
    
    rows, err := q.db.Query(ctx, getRecord, targetID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    items := []Record{}
    for rows.Next() {
        var i Record
        if err := rows.Scan(
            &i.TargetID,
            &i.RoleType,
            &i.MediaType,
            &i.Message,
            &i.Media,
            &i.Isread,
            &i.CreatedAt,
        ); err != nil {
            return nil, err
        }
        items = append(items, i)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return items, nil
}



func (q *Queries) GetTargetID(ctx context.Context,tablename string) ([]int32, error) {
    getTargetID := fmt.Sprintf(`-- name: GetTargetID :many
    SELECT DISTINCT target_id FROM %s`,tablename)

    rows, err := q.db.Query(ctx, getTargetID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    items := []int32{}
    for rows.Next() {
        var target_id int32
        if err := rows.Scan(&target_id); err != nil {
            return nil, err
        }
        items = append(items, target_id)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return items, nil
}


type GetLastMsgRow struct {
    Message   string    `json:"message"`
    MediaType string    `json:"media_type"`
    Isread    bool      `json:"isread"`
}

func (q *Queries) GetLastMsg(ctx context.Context, targetID int32, tablename string) (GetLastMsgRow, error) {

	getLastMsg := fmt.Sprintf(`-- name: GetLastMsg :one
	SELECT message,media_type,isread FROM %s
	WHERE target_id = $1
	ORDER BY created_at DESC
	LIMIT 1
	`,tablename)

    row := q.db.QueryRow(ctx, getLastMsg, targetID)
    var i GetLastMsgRow
    err := row.Scan(&i.Message, &i.MediaType, &i.Isread)
    return i, err
}



type UpdateRecordParams struct {
    TargetID  int32         `json:"target_id"`
    MediaType string        `json:"media_type"`
    Message   string        `json:"message"`
    CreatedAt time.Time     `json:"created_at"`
}

func (q *Queries) UpdateRecord(ctx context.Context, arg UpdateRecordParams, tablename string) (Record, error) {
    
    updateRecord := fmt.Sprintf(`-- name: UpdateRecord :one
    UPDATE %s
    SET message = $4
    WHERE target_id = $1
    AND media_type = $2
    AND created_at = $3
    RETURNING target_id, role_type, media_type, message, media, isread, created_at
    `,tablename)

    row := q.db.QueryRow(ctx, updateRecord,
        arg.TargetID,
        arg.MediaType,
        arg.CreatedAt,
        arg.Message,
    )
    var i Record
    err := row.Scan(
        &i.TargetID,
        &i.RoleType,
        &i.MediaType,
        &i.Message,
        &i.Media,
        &i.Isread,
        &i.CreatedAt,
    )
    return i, err
}
