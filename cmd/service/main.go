package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// 初始化客户端
	client, err := ethclient.Dial(os.Getenv("BSC_TEST_RPC_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to the BSC network: %v", err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 公钥地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Printf("Public key: %s\n", fromAddress.Hex())

	// 获取链ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	fmt.Printf("Chain ID: %d\n", chainID)

	fmt.Println("Go Ethereum SDK初始化完成")
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
