package hyperledger

import (
	"fmt"
	"testing"
)

func Test_ConnectionTest_Success(t *testing.T) {

	clients := NewClientMap(
		"test-network",
		"/home/shwetha/go/src/github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/hyperledger/config.yaml",
	)

	_, err := clients.AddClient(
		"Admin",
		"ibm",
		"mainchannel",
	)
	if err != nil {
		t.Fatal(err)
		return
	}

	res, err := clients.Query("ibm", "mainchannel", "resources", "index", [][]byte{
		[]byte(""),
		[]byte(""),
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Println(string(res))
}
