package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"go.etcd.io/bbolt"
)

// KV type
type KV struct {
	terminate chan bool
	config    *Config
	events    chan Message
	updates   chan Message
	sync      chan Message
	db        *bbolt.DB
	dbPath    string
	log       *Log
}

// KVUpdate type
type KVUpdate struct {
}

func newKV(app *Bunker) (*KV, error) {
	kv := &KV{
		terminate: make(chan bool, 1),
		config:    app.Config,
		events:    app.events,
		updates:   app.updates,
		sync:      app.sync,
		log:       app.Logger,
		dbPath:    app.Config.KV.DBPath,
	}
	if _, err := os.Stat(kv.dbPath); os.IsNotExist(err) {
		p := strings.Split(kv.dbPath, "/")
		if len(p) > 1 {
			s := p[:len(p)-1]
			q := strings.Join(s, "/")
			err := os.MkdirAll(q, 0755)
			if err != nil {
				return kv, err
			}
		}
	}
	options := &bbolt.Options{
		Timeout:      30 * time.Second,
		FreelistType: "hashmap",
	}
	db, err := bbolt.Open(kv.dbPath, 0755, options)
	if err != nil {
		return kv, err
	}
	kv.db = db
	return kv, nil
}

// Start func
func (kv *KV) Start() {
	for {
		select {
		case <-kv.terminate:
			return
		case msg := <-kv.updates:
			err := kv.handleUpdate(msg)
			if err != nil {
				fmt.Println(err)
			}

		case msg := <-kv.events:
			err := kv.handleEvent(msg)
			if err != nil {
				fmt.Println(err)
			}
		case msg := <-kv.sync:
			err := kv.handleSync(msg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (kv *KV) handleUpdate(msg Message) error {
	fmt.Println(msg)
	return nil
}

func (kv *KV) handleSync(msg Message) error {
	fmt.Println(msg)
	return nil
}

func (kv *KV) handleEvent(msg Message) error {
	fmt.Println(msg)
	return nil
}
