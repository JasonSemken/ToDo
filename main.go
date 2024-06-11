package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

var DB *bolt.DB

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	homePage(db)
	if err != nil {
		log.Fatal(err)
	}
}

func setupDB() (*bolt.DB, error) {

	// Load database, one will be created if it does not exist
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to load db")
	}

	// Create bucket
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("failed to create bucket: %s", err)
		}
		return nil
	})
	fmt.Println("DB connection OK")
	return db, nil
}

func updateItem(db *bolt.DB, key []byte, value []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}

func viewItem(db *bolt.DB, key []byte) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte(key))
		fmt.Printf("%s is %s\n", key, v)
		return nil
	})

	return err
}

func returnInputItem(db *bolt.DB, key []byte) error {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte(key))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})

	return err
}

func returnAllItems(db *bolt.DB) {
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("%s: %s\n", k, v)
			return nil
		})
		return nil
	})
}

func homePage(db *bolt.DB) {
	var input int
	var key []byte

	fmt.Println("")
	fmt.Println("Enter 1 to view list item")
	fmt.Println("Enter 2 to enter item")
	fmt.Println("Enter 3 to view all items")
	fmt.Println("Enter 4 to exit")
	fmt.Scanln(&input)
	switch input {
	case 1: // returns users item based on key to them. TODO check for blank item
		fmt.Println("Input key")
		fmt.Scanln(&key)
		viewItem(db, key)
		homePage(db)
	case 2: // Allows user to enter a item into the db
		fmt.Println("Input key")
		fmt.Scanln(&key)
		fmt.Println("Input value")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		value := scanner.Text()
		updateItem(db, key, []byte(value))
		returnInputItem(db, key)
		homePage(db)
	case 3: // returns all entries in the db
		returnAllItems(db)
		homePage(db)
	case 4: // closes the program
		os.Exit(0)
	}
}
