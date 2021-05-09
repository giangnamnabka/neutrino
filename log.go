package neutrino

import (
	"github.com/btcsuite/btclog"
	"github.com/giangnamnabka/btcd/addrmgr"
	"github.com/giangnamnabka/btcd/blockchain"
	"github.com/giangnamnabka/btcd/connmgr"
	"github.com/giangnamnabka/btcd/peer"
	"github.com/giangnamnabka/btcd/txscript"
	"github.com/giangnamnabka/neutrino/blockntfns"
	"github.com/giangnamnabka/neutrino/pushtx"
	"github.com/giangnamnabka/neutrino/query"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log btclog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until either UseLogger or SetLogWriter are called.
func DisableLog() {
	log = btclog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
// This should be used in preference to SetLogWriter if the caller is also
// using btclog.
func UseLogger(logger btclog.Logger) {
	log = logger
	blockchain.UseLogger(logger)
	txscript.UseLogger(logger)
	peer.UseLogger(logger)
	addrmgr.UseLogger(logger)
	blockntfns.UseLogger(logger)
	pushtx.UseLogger(logger)
	connmgr.UseLogger(logger)
	query.UseLogger(logger)
}
