package dbclient

import (
	"fmt"
	"log"
	"time"

	"github.com/AXR8ELW/data-pump-file/calorietracker/internal/app/model"
	"github.com/boltdb/bolt"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryWeight() (model.Weight, error)
	Seed()
}
type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("calorietracker.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
}
func (bc *BoltClient) InitializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("WEIGHT"))
		if err != nil {
			return fmt.Errorf("could not create weight bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("ENTRIES"))
		if err != nil {
			return fmt.Errorf("could not create days bucket: %v", err)
		}
		return nil
	})
	fmt.Println("DB Setup Done")
}
func (bc *BoltClient) SeedWeight() {
	total := 10
	for i := 0; i < total; i++ {
		key := strconv.Itoa(1 + i)
		weight := model.Weight{
			Id:   key,
			Date = now.AddDate(0, i, 0)
			fmt.Println("\nAdd 1 Month:", Date)
			Weight: 50.5 + i,
		}
		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(weight)
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("WEIGHT"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v weight\n", total)
}
func (bc *BoltClient) QueryWeight() (model.Weights, error) {
	// Allocate an empty Weight instance
	weight := model.Weight{}
	weights := model.Weights{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("WEIGHT"))

		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			json.Unmarshal(v, &weight)
			weights.Weight = append(weights.Weight, weight)
			return nil
		})

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Weights{}, err
	}
	// Return the Account struct and nil as error.
	return weights, nil
}
