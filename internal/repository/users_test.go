package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/IDarar/test-task/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
)

var postgresTestURL = "postgres://root:secret@localhost:5432/cashback?sslmode=disable"

const (
	users = `
	create table if not exists users(
		"id" serial,
		"name" VARCHAR(24) UNIQUE,
		"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	
		CONSTRAINT "scan_tx_pk" PRIMARY KEY ("id")
	);
	`

	dropUsers = "drop table if exists users;"
)

func GetTestDB() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(postgresTestURL)
	if err != nil {
		log.Fatal(err)
	}

	config.MaxConns = 100
	config.ConnConfig.PreferSimpleProtocol = true

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to connect to database:", err)
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		fmt.Printf("postgres does not respond: %v\n", err)
		return nil, err
	}

	return pool, nil
}

func SetUpTables(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), users)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func TearDownTestDb(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), dropUsers)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func TestCreate(t *testing.T) {
	db, err := GetTestDB()
	require.NoError(t, err)

	//for testing unexpected closing of connections
	otherDb, err := GetTestDB()
	require.NoError(t, err)

	err = SetUpTables(db)
	require.NoError(t, err)

	defer func() {
		err = TearDownTestDb(otherDb)
		require.NoError(t, err)
	}()

	r := NewUsersRepo(db)

	u, err := r.Create("Name1")
	require.NoError(t, err)

	require.Equal(t, u.Name, "Name1")
	require.Equal(t, u.ID, 1)

	//duplicate
	u, err = r.Create("Name1")
	require.ErrorIs(t, err, domain.ErrUserAlreadyExists)
}

func TestDelete(t *testing.T) {
	db, err := GetTestDB()
	require.NoError(t, err)

	//for testing unexpected closing of connections
	otherDb, err := GetTestDB()
	require.NoError(t, err)

	err = SetUpTables(db)
	require.NoError(t, err)

	defer func() {
		err = TearDownTestDb(otherDb)
		require.NoError(t, err)
	}()

	r := NewUsersRepo(db)

	//user does not exists
	err = r.Delete(1)
	require.Error(t, err)

	_, err = db.Exec(context.Background(), "insert into users DEFAULT VALUES")
	require.NoError(t, err)

	err = r.Delete(1)
	require.NoError(t, err)
}

func TestList(t *testing.T) {
	db, err := GetTestDB()
	require.NoError(t, err)

	//for testing unexpected closing of connections
	otherDb, err := GetTestDB()
	require.NoError(t, err)

	err = SetUpTables(db)
	require.NoError(t, err)

	defer func() {
		err = TearDownTestDb(otherDb)
		require.NoError(t, err)
	}()

	r := NewUsersRepo(db)

	//no rows
	uList, err := r.List()
	require.NoError(t, err)

	require.Equal(t, len(uList), 0)

	for i := 0; i < 10; i++ {
		_, err := db.Exec(context.Background(), "insert into users (name) values ($1)", fmt.Sprint(i))
		require.NoError(t, err)
	}

	uList, err = r.List()
	require.NoError(t, err)

	require.Equal(t, len(uList), 10)

	require.Equal(t, uList[4].ID, 5)
}
