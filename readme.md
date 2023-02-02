# Personal Ethereum Client
```Work in progress!``` In this project am trying to build the underlying backend system for an ethereum dapp using the clients provided by the ethereum protocols, i'm abstracting functionalities from ethereum's own abstraction too.

# Usage
Clone or download the project and run ```sh metro-eth.sh``` in the terminal and watch the magic (this is only temporary as its still wip) after the development, the mode of running the project would change. the results would be as below.

# Sending Ether

```
pkey, err := crypto.ToECDSA(privkey [string])
if err != nil {
    log.Printf("error parsing private key:v\n", err)
    return
}

// fetch the public key from the private key
pubkey := pkey.Public()
publickey, ok := pubkey.(*ecdsa.Publickey)
if !ok{
    log.Printf("error casting private to public")
    return
}

// fetch the address from publickey
fromaddress := crypto.PubKeyToAddress(publickey)

// find the nonce [uint64]
nonce, err := client.PendingNonceAt(fromaddress)

check ```custom/sendether.go``` to see complete!
```