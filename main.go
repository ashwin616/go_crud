package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

func main() {
	var err error

	// ✅ Connect to PostgreSQL
	// db, err = pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/go_crud")
	db, err = pgx.Connect(context.Background(), "postgres://postgres:1234@localhost:5432/go_crud?sslmode=disable")

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close(context.Background())

	fmt.Println("✅ Connected to PostgreSQL")

	// 🔄 CRUD Operations
	CreateUser("Alice", "alice@example.com")
	CreateUser("Bob", "bob@example.com")

	fmt.Println("\n📄 All Users:")
	GetUsers()

	fmt.Println("\n✏️ Update User ID 1")
	UpdateUser(1, "alice_new@example.com")

	fmt.Println("\n📄 All Users After Update:")
	GetUsers()

	fmt.Println("\n❌ Delete User ID 2")
	DeleteUser(2)

	fmt.Println("\n📄 All Users After Delete:")
	GetUsers()
}

func CreateUser(name, email string) {
	_, err := db.Exec(context.Background(),
		"INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		log.Printf("❌ CreateUser error: %v\n", err)
	} else {
		fmt.Println("✅ User created:", name)
	}
}

func GetUsers() {
	rows, err := db.Query(context.Background(), "SELECT id, name, email FROM users")
	if err != nil {
		log.Printf("❌ GetUsers error: %v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Println("❌ Scan error:", err)
			continue
		}
		fmt.Printf("👤 ID: %d | Name: %s | Email: %s\n", id, name, email)
	}
}

func UpdateUser(id int, newEmail string) {
	_, err := db.Exec(context.Background(),
		"UPDATE users SET email = $1 WHERE id = $2", newEmail, id)
	if err != nil {
		log.Printf("❌ UpdateUser error: %v\n", err)
	} else {
		fmt.Println("✅ User updated")
	}
}

func DeleteUser(id int) {
	_, err := db.Exec(context.Background(),
		"DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Printf("❌ DeleteUser error: %v\n", err)
	} else {
		fmt.Println("✅ User deleted")
	}
}
