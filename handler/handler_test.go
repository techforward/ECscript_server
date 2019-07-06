package handler

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/alexsasharegan/dotenv"
	"github.com/techforward/ECscript_server/util"
)

func TestMain(m *testing.M) {

	err := dotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// terminalWidth := goterm.Width()
	// a := (terminalWidth - 12) / 2
	b := strings.Repeat("=", 20)
	c := strings.Repeat("=", 52)

	println(c)
	println(b + "            " + b)
	println(b + " Start TEST " + b)
	println(b + "            " + b)
	println(c)
	testAddressID = util.NewUlid()
	testCartID = util.NewUlid()
	testItemID = util.NewUlid()
	testOrderID = util.NewUlid()
	testUserID = util.NewUlid()
	testFirebaseUID = util.NewUlid()
	testBoughtAt = time.Now().Format("2006-01-02T15:04:05Z")

	fmt.Printf("=== Generate testAddressID: %v\n", testAddressID)
	fmt.Printf("=== Generate testCartID   : %v\n", testCartID)
	fmt.Printf("=== Generate testBoughtAt : %v\n", testBoughtAt)
	fmt.Printf("=== Generate testItemID   : %v\n", testItemID)
	fmt.Printf("=== Generate testOrderID  : %v\n", testOrderID)
	fmt.Printf("=== Generate testUserID   : %v\n", testUserID)
	fmt.Printf("=== Generate testFirebaseUID   : %v\n", testFirebaseUID)

	code := m.Run()

	println(c)
	println(b + "            " + b)
	println(b + "  End TEST  " + b)
	println(b + "            " + b)
	println(c)

	os.Exit(code)
}
