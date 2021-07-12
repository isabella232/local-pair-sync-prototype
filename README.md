## Native Go Zyre

### Repo:
[go-zeromq/zyre](https://github.com/go-zeromq/zyre)

### Test code:
[`lib_test/zyre_test.go`](./lib_tests/zyre_test.go)

### Test Output(s)

**macos: wifi: single device peers:**
```text
=== RUN   TestZyre
    TestZyre: zyre_test.go:67: (*zyre.Event)(0xc000234000)({
         Type: (string) (len=5) "ENTER",
         PeerUUID: (string) (len=32) "B1D08EEE664CAB5955546B5BC5AD18B1",
         PeerName: (string) (len=6) "B1D08E",
         PeerAddr: (string) (len=24) "tcp://192.168.0.14:58858",
         Headers: (map[string]string) {
         },
         Group: (string) "",
         Msg: ([]uint8) <nil>
        })
        
    TestZyre: zyre_test.go:67: (*zyre.Event)(0xc0000c6000)({
         Type: (string) (len=5) "ENTER",
         PeerUUID: (string) (len=32) "945D8DD9CE3C20527CEFCC4BC7A21343",
         PeerName: (string) (len=6) "945D8D",
         PeerAddr: (string) (len=24) "tcp://192.168.0.14:58859",
         Headers: (map[string]string) {
         },
         Group: (string) "",
         Msg: ([]uint8) <nil>
        })
        
([]string) (len=1 cap=1) {
 (string) (len=32) "945D8DD9CE3C20527CEFCC4BC7A21343"
}
([]string) (len=1 cap=1) {
 (string) (len=32) "B1D08EEE664CAB5955546B5BC5AD18B1"
}
--- PASS: TestZyre (9.01s)
PASS
```

**Notes:**

The peers can see the `ENTER` event but the `SHOUT` and `JOIN` events are not received

## Native Go ZMQ4

### Repo:
[go-zeromq/zmq4](https://github.com/go-zeromq/zmq4)

### Test code:
[`lib_test/zmq4_test.go`](./lib_tests/zmq4_test.go)

### Test Output(s):

#### TestZMQ4ReqRep():
**macos: wifi: single device peers:**
```text
=== RUN   TestZMQ4ReqRep
rrworker: 2021/07/09 16:43:01 received request: [Hello]
rrclient: 2021/07/09 16:43:02 received reply 0 [World]
rrworker: 2021/07/09 16:43:02 received request: [Hello]
rrclient: 2021/07/09 16:43:03 received reply 1 [World]
rrworker: 2021/07/09 16:43:03 received request: [Hello]
rrclient: 2021/07/09 16:43:04 received reply 2 [World]
rrworker: 2021/07/09 16:43:04 received request: [Hello]
rrclient: 2021/07/09 16:43:05 received reply 3 [World]
rrworker: 2021/07/09 16:43:05 received request: [Hello]
rrclient: 2021/07/09 16:43:06 received reply 4 [World]
rrworker: 2021/07/09 16:43:06 received request: [Hello]
rrclient: 2021/07/09 16:43:07 received reply 5 [World]
rrworker: 2021/07/09 16:43:07 received request: [Hello]
rrclient: 2021/07/09 16:43:08 received reply 6 [World]
rrworker: 2021/07/09 16:43:08 received request: [Hello]
rrclient: 2021/07/09 16:43:09 received reply 7 [World]
rrworker: 2021/07/09 16:43:09 received request: [Hello]
rrclient: 2021/07/09 16:43:10 received reply 8 [World]
rrworker: 2021/07/09 16:43:10 received request: [Hello]
rrclient: 2021/07/09 16:43:11 received reply 9 [World]
--- PASS: TestZMQ4ReqRep (15.00s)
PASS
```

## Peer Discovery

### Repo:
[schollz/peerdiscovery](https://github.com/schollz/peerdiscovery)

### Test code:
[`lib_test/peer_discovery_test.go`](./lib_tests/peer_discovery_test.go)

### Test Output(s):

**macos: wifi: 1 peer per device:**
```text
Scanning for 10 seconds to find LAN peers
Payload sending : 'ahQBlMvUYE'
2021/07/12 14:24:32 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:32 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:33 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:33 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:34 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:34 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:35 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:35 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:36 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:36 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:37 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:37 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:38 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:38 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:39 address: 192.168.0.17, payload: DNDVvRmirX
2021/07/12 14:24:39 address: 192.168.0.17, payload: DNDVvRmirX
Found 1 other computers
0) '192.168.0.17' with payload 'DNDVvRmirX'
--- PASS: TestAnother (10.00s)
PASS
```

**Windows: ethernet cable: 1 peer per device:**
```text
=== RUN   TestAnother
Scanning for 10 seconds to find LAN peers
Payload sending : 'DNDVvRmirX'
2021/07/09 16:54:54 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:54 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:55 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:55 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:55 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:55 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:56 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:56 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:56 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:56 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:57 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:57 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:57 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:57 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:58 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:58 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:58 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:58 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:59 address: 192.168.0.14, payload: ahQBlMvUYE
2021/07/09 16:54:59 address: 192.168.0.14, payload: ahQBlMvUYE
Found 1 other computers
0) '192.168.0.14' with payload 'ahQBlMvUYE'
--- PASS: TestAnother (10.34s)
PASS
```

**Notes:**

~~Windows/ethernet peer device can see and receive payload from the macOS/Wifi device peer. But the MacOS/Wifi device can not see the Windows/ethernet device and/or receive the payload.~~

Follow up, the issue is caused by the system firewall silently blocking all network calls for my application.
