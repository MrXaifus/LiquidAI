package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// BlockchainClient handles blockchain connection and interactions
type BlockchainClient struct {
	client      *ethclient.Client
	privateKey  *ecdsa.PrivateKey
	publicAddr  common.Address
	networkName string
}

// NewBlockchainClient initializes a new BlockchainClient
func NewBlockchainClient(rpcURL, privateKeyHex, networkName string) (*BlockchainClient, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RPC: %v", err)
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	publicAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	return &BlockchainClient{
		client:      client,
		privateKey:  privateKey,
		publicAddr:  publicAddr,
		networkName: networkName,
	}, nil
}

// GetWalletBalance fetches ETH balance of the given wallet address
func (bc *BlockchainClient) GetWalletBalance(ctx context.Context, address common.Address) (*big.Float, error) {
	balance, err := bc.client.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("error fetching balance: %v", err)
	}
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))
	return ethBalance, nil
}

// SimulateAIPrediction pretends to analyze blockchain data and predict market trend
func (bc *BlockchainClient) SimulateAIPrediction(balance *big.Float) (string, float64) {
	// Very simple fake AI model: if wallet balance > 1 ETH, predict "Bullish", else "Bearish"
	threshold := big.NewFloat(1.0)
	cmp := balance.Cmp(threshold)

	if cmp >= 0 {
		return "Bullish", 0.82 // 82% confidence
	} else {
		return "Bearish", 0.76 // 76% confidence
	}
}

// SendDummyTransaction sends a small transaction (simulation of "trading action")
func (bc *BlockchainClient) SendDummyTransaction(ctx context.Context, to common.Address) (string, error) {
	nonce, err := bc.client.PendingNonceAt(ctx, bc.publicAddr)
	if err != nil {
		return "", fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := bc.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to suggest gas price: %v", err)
	}

	tx := bind.NewKeyedTransactor(bc.privateKey)
	value := big.NewInt(1e15) // 0.001 ETH

	auth := bind.TransactOpts{
		From:     bc.publicAddr,
		Signer:   tx.Signer,
		Value:    value,
		GasPrice: gasPrice,
		GasLimit: uint64(21000),
		Nonce:    big.NewInt(int64(nonce)),
		Context:  ctx,
		NoSend:   false,
	}

	// Actually here, we'd build and send a real tx, but we fake a transaction hash for now
	fakeHash := "0xDEADBEEF1234567890ABCDEF"
	return fakeHash, nil
}

// HTTP Handlers
func balanceHandler(bc *BlockchainClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		addressHex := r.URL.Query().Get("address")
		if addressHex == "" {
			http.Error(w, "Missing address parameter", http.StatusBadRequest)
			return
		}

		addr := common.HexToAddress(addressHex)
		balance, err := bc.GetWalletBalance(ctx, addr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"address": addressHex,
			"balance": balance.String(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func predictHandler(bc *BlockchainClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		addressHex := r.URL.Query().Get("address")
		if addressHex == "" {
			http.Error(w, "Missing address parameter", http.StatusBadRequest)
			return
		}

		addr := common.HexToAddress(addressHex)
		balance, err := bc.GetWalletBalance(ctx, addr)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
			return
		}

		trend, confidence := bc.SimulateAIPrediction(balance)

		response := map[string]interface{}{
			"address":   addressHex,
			"trend":     trend,
			"confidence": fmt.Sprintf("%.2f%%", confidence*100),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	rpcURL := os.Getenv("RPC_URL")
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	networkName := os.Getenv("NETWORK_NAME")

	if rpcURL == "" || privateKeyHex == "" {
		log.Fatal("Missing environment variables: RPC_URL and PRIVATE_KEY are required")
	}

	client, err := NewBlockchainClient(rpcURL, privateKeyHex, networkName)
	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	http.HandleFunc("/balance", balanceHandler(client))
	http.HandleFunc("/predict", predictHandler(client))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("LiquidAI API is running on port %s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
