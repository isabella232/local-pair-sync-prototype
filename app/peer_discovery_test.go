package app

import (
	"fmt"
	"log"
	mrand "math/rand"
	"sync"
	"testing"
	"time"

	"github.com/schollz/peerdiscovery"
)

func TestPeerDiscovery(t *testing.T) {
	//wg := new(sync.WaitGroup)

	for x:=0;x<10;x++{
		t.Logf("attempting discovery %d", x+1)
		ds, err := peerdiscovery.Discover(peerdiscovery.Settings{Limit: 1})
		if err != nil {
			t.Error(err)
		}

		//t.Logf("Read peer list of %s", "peer 1")
		for _, d := range ds {
			fmt.Printf("discovered '%s'\n", d.Address)
		}
	}

	/*wg.Add(1)
	go peerDiscovery(t, "peer 1", wg)

	wg.Add(1)
	go peerDiscovery(t, "peer 2", wg)
	 */

	//wg.Wait()
}

func peerDiscovery(t *testing.T, name string, wg *sync.WaitGroup) {
	t.Logf("Begin discovery %s", name)
	ds, err := peerdiscovery.Discover(peerdiscovery.Settings{Limit: 1})
	if err != nil {
		t.Error(err)
	}

	t.Logf("Read peer list of %s", name)
	for _, d := range ds {
		fmt.Printf("discovered '%s'\n", d.Address)
	}
	wg.Done()
}

func TestAnother(t *testing.T) {
	fmt.Println("Scanning for 10 seconds to find LAN peers")
	pl := randStringBytesMaskImprSrc(10)
	fmt.Printf("Payload sending : '%s'\n", pl)

	// discover peers
	discoveries, err := peerdiscovery.Discover(peerdiscovery.Settings{
		Limit:     -1,
		Payload:   []byte(pl),
		Delay:     500 * time.Millisecond,
		TimeLimit: 10 * time.Second,
		Notify: func(d peerdiscovery.Discovered) {
			log.Println(d)
		},
	})

	// print out results
	if err != nil {
		log.Fatal(err)
	} else {
		if len(discoveries) > 0 {
			fmt.Printf("Found %d other computers\n", len(discoveries))
			for i, d := range discoveries {
				fmt.Printf("%d) '%s' with payload '%s'\n", i, d.Address, d.Payload)
			}
		} else {
			fmt.Println("Found no devices. You need to run this on another computer at the same time.")
		}
	}
}

// src is seeds the random generator for generating random strings
var src = mrand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandStringBytesMaskImprSrc prints a random string
func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
