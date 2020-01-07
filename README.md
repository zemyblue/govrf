# VRF-go
This is VRF library of golang version 
I forked [algorand's libsodium](https://github.com/algorand/libsodium) and extract only vrf module from go-algorand.

### How to compile libsodium of algorand
```shell script
cd ./crypto/libsodium
./autogen.sh
./configure
make && make check
sudo make install
```
And then, libsodium installed under /usr/local/lib/ if MacOS.

### Build and Test
```shell script
go build
go test -v
```


# VRF Types
```go
type VrfPrivkey [64]uint8
type VrfPubkey [32]uint8
```

# VRF functions
## make private key and public key
```go
func VrfKeygenFromSeed(seed [32]byte) (pub VrfPubkey, priv VrfPrivkey)
func VrfKeygen() (pub VrfPubkey, priv VrfPrivKey)
```

## VrfPrivKey
```go
func (sk VrfPrivkey) Pubkey() (pk VrfPubkey)
func (sk VrfPrivkey) Prove(message []byte) (proof VrfProof, ok bool)
func (proof VrfProof) Hash() (hash VrfOutput, ok bool)
```

## VrfPubKey
```go
func (pk VrfPubkey) Verify(p VrfProof, message []byte) (bool, VrfOutput
``` 

