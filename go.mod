module github.com/giangnamnabka/neutrino

require (
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/davecgh/go-spew v1.1.1
	github.com/giangnamnabka/btcd v0.21.0-beta.0.20210509080720-d664d3d3dfd7
	github.com/giangnamnabka/btcutil v1.0.3-0.20210509081240-f80751d24cbd
	github.com/giangnamnabka/btcwallet/wallet/txauthor v1.0.1-0.20210420072038-a00c5ff72602
	github.com/giangnamnabka/btcwallet/walletdb v1.3.4
	github.com/giangnamnabka/btcwallet/wtxmgr v1.2.0
	github.com/lightningnetwork/lnd/queue v1.0.1
	github.com/stretchr/testify v1.5.1
)

replace github.com/giangnamnabka/btcwallet/walletdb => /home/tinhnguyenhn/Desktop/20210508/btcwallet/walletdb

replace github.com/giangnamnabka/btcwallet/wtxmgr => /home/tinhnguyenhn/Desktop/20210508/btcwallet/wtxmgr

go 1.13
