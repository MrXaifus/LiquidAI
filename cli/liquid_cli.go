package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// Global CLI flags
var (
	serverURL string
)

func init() {
	flag.StringVar(&serverURL, "server", "http://localhost:8080", "LiquidAI API server URL")
}

// handleRegisterAgent simulates agent registration to LiquidAI
func handleRegisterAgent(agentName string) {
	fmt.Printf("🔵 Registering agent '%s' to LiquidAI at %s...\n", agentName, serverURL)
	// Here you would normally POST to /agents/register
	fmt.Println("✅ Agent registered successfully.")
}

// handlePredictMarket simulates requesting a market prediction
func handlePredictMarket(walletAddress string) {
	fmt.Printf("🔵 Requesting market prediction for wallet '%s' at %s...\n", walletAddress, serverURL)
	// Here you would normally GET from /predict?address=walletAddress
	fmt.Println("📈 Predicted trend: Bullish (82% confidence)")
}

// handleQueryBalance simulates querying on-chain wallet balance
func handleQueryBalance(walletAddress string) {
	fmt.Printf("🔵 Querying balance for wallet '%s' at %s...\n", walletAddress, serverURL)
	// Here you would normally GET from /balance?address=walletAddress
	fmt.Println("💰 Wallet Balance: 2.37 ETH")
}

// handleHelp prints available CLI commands
func handleHelp() {
	fmt.Println(`LiquidAI CLI - Command Line Interface

Usage:
  liquid_cli [command] [arguments]

Available Commands:
  register-agent   Register a new AI agent into LiquidAI network
  predict-market   Request a market trend prediction for a wallet
  query-balance    Query blockchain wallet balance
  help             Show this help menu

Examples:
  liquid_cli register-agent --name=AgentX
  liquid_cli predict-market --wallet=0xDEADBEEF...
  liquid_cli query-balance --wallet=0xDEADBEEF...

Use "-server=http://your-server.com" to specify API server.
`)
}

func main() {
	// Parse initial global flags
	flag.Parse()

	if len(os.Args) < 2 {
		handleHelp()
		os.Exit(1)
	}

	// Determine subcommand
	cmd := os.Args[1]

	switch cmd {
	case "register-agent":
		registerCmd := flag.NewFlagSet("register-agent", flag.ExitOnError)
		agentName := registerCmd.String("name", "", "Name of the agent")
		registerCmd.Parse(os.Args[2:])

		if *agentName == "" {
			fmt.Println("❌ Missing required argument: --name")
			registerCmd.Usage()
			os.Exit(1)
		}
		handleRegisterAgent(*agentName)

	case "predict-market":
		predictCmd := flag.NewFlagSet("predict-market", flag.ExitOnError)
		walletAddress := predictCmd.String("wallet", "", "Wallet address to predict")
		predictCmd.Parse(os.Args[2:])

		if *walletAddress == "" {
			fmt.Println("❌ Missing required argument: --wallet")
			predictCmd.Usage()
			os.Exit(1)
		}
		handlePredictMarket(*walletAddress)

	case "query-balance":
		balanceCmd := flag.NewFlagSet("query-balance", flag.ExitOnError)
		walletAddress := balanceCmd.String("wallet", "", "Wallet address to query")
		balanceCmd.Parse(os.Args[2:])

		if *walletAddress == "" {
			fmt.Println("❌ Missing required argument: --wallet")
			balanceCmd.Usage()
			os.Exit(1)
		}
		handleQueryBalance(*walletAddress)

	case "help":
		handleHelp()

	default:
		fmt.Printf("❌ Unknown command: %s\n\n", cmd)
		handleHelp()
		os.Exit(1)
	}
}
