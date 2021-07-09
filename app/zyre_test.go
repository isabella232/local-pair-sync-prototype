package app

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-zeromq/zyre"
	"sync"
	"testing"
	"time"
)

func TestZyre(t *testing.T) {
	z1, cncl1 := makeNode()
	defer cncl1()

	z2, cncl2 := makeNode()
	defer cncl2()

	z1.SetVerbose()
	z2.SetVerbose()

	err := z1.Start()
	if err != nil {
		spew.Dump(err)
	}

	defer z1.Stop()

	err = z2.Start()
	if err != nil {
		spew.Dump(err)
	}

	defer z2.Stop()

	wg := new(sync.WaitGroup)
	//wg.Add(1)
	go watchEvents(t, wg, z1)
	//wg.Add(1)
	go watchEvents(t, wg, z2)

	z1.Join("TALK")
	z2.Join("TALK")

	time.Sleep(time.Second)
	z1.Shout("TALK", []byte("I am here"))
	time.Sleep(time.Second)
	z1.Shout("TALK", []byte("Hi world"))

	//wg.Wait()
	time.Sleep(time.Second)

	spew.Dump(z1.Peers(), z2.Peers())
}

func makeNode() (*zyre.Zyre, context.CancelFunc) {
	ctx, cncl := context.WithCancel(context.Background())
	z := zyre.NewZyre(ctx)

	return z, cncl
}

func watchEvents(t *testing.T, wg *sync.WaitGroup, z *zyre.Zyre) {
	select {
	case e := <- z.Events():
		t.Log(spew.Sdump(e))
		if string(e.Msg) == "Hi world" {
			//wg.Done()
			return
		}
	}
}