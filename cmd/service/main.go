package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")

	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Printf("Error fetching block number: %v\n", err)
		panic(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Printf("Error fetching chain ID: %v\n", err)
		panic(err)
	}

	address := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		fmt.Printf("Error fetching balance: %v\n", err)
		panic(err)
	}

	fmt.Printf("Current block number: %d\n", blockNumber)
	fmt.Printf("Chain ID: %s\n", chainID)
	fmt.Printf("Balance: %s\n", balance.String())

	defer client.Close()
}

//var (
//	topic  = "user_click"
//	reader *kafka.Reader
//)
//
//func writeKafka(ctx context.Context) {
//	writer := kafka.Writer{
//		Addr:                   kafka.TCP("localhost:9092"),
//		Topic:                  topic,
//		Balancer:               &kafka.Hash{},
//		WriteTimeout:           1 * time.Second,
//		RequiredAcks:           kafka.RequireNone,
//		AllowAutoTopicCreation: true,
//	}
//
//	defer writer.Close()
//
//	for i := 0; i < 3; i++ {
//		err := writer.WriteMessages(
//			ctx,
//			kafka.Message{Key: []byte("1"), Value: []byte("P")},
//			kafka.Message{Key: []byte("2"), Value: []byte("L")},
//			kafka.Message{Key: []byte("3"), Value: []byte("A")},
//			kafka.Message{Key: []byte("1"), Value: []byte("Y")},
//			kafka.Message{Key: []byte("2"), Value: []byte("E")},
//			kafka.Message{Key: []byte("3"), Value: []byte("R")},
//			kafka.Message{Key: []byte("1"), Value: []byte("G")},
//			kafka.Message{Key: []byte("2"), Value: []byte("S")},
//		)
//
//		if err != nil {
//			if errors.Is(err, kafka.LeaderNotAvailable) {
//				time.Sleep(500 * time.Millisecond)
//				continue
//			} else {
//				fmt.Printf("Error writing to kafka: %v\n", err)
//			}
//		} else {
//			break
//		}
//	}
//}
//
//func readKafka(ctx context.Context) {
//	reader = kafka.NewReader(kafka.ReaderConfig{
//		Brokers:        []string{"localhost:9092"},
//		Topic:          topic,
//		CommitInterval: 1 * time.Second,
//		GroupID:        "rec_team",
//		StartOffset:    kafka.FirstOffset,
//		Partition:      0,
//	})
//
//	//defer reader.Close()
//
//	for {
//		msg, err := reader.ReadMessage(ctx)
//
//		if err != nil {
//			fmt.Printf("Error reading message from kafka: %v\n", err)
//			continue
//		} else {
//			fmt.Printf("topic=%s, partition=%d, offset=%d, key=%s, value=%s timestamp=%T\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value), msg.Time)
//		}
//	}
//
//}
//
//func listenSignal() {
//	c := make(chan os.Signal, 1)
//
//	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
//	sig := <-c
//	fmt.Printf("Got signal: %v, exiting...\n", sig)
//	if reader != nil {
//		reader.Close()
//	}
//	os.Exit(0)
//}
//
//func main() {
//	ctx := context.Background()
//
//	//writeKafka(ctx)
//
//	go listenSignal()
//	readKafka(ctx)
//}
