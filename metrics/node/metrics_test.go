package node

import (
	"testing"

	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/stretchr/testify/require"

	prom "github.com/prometheus/client_golang/prometheus"
)

func TestParsingLabelsFromNodeName(t *testing.T) {
	var labels prom.Labels
	var err error

	// mobile name
	labels, err = labelsFromNodeName("StatusIM/v0.30.1-beta.2/android-arm/go1.12")
	require.NoError(t, err)
	require.Equal(t, labels,
		prom.Labels{
			"platform": "android-arm",
			"type":     "StatusIM",
			"version":  "v0.30.1-beta.2",
		})
	// desktop name
	labels, err = labelsFromNodeName("Statusd/v0.29.0-beta.2/linux-amd64/go1.11")
	require.NoError(t, err)
	require.Equal(t, labels,
		prom.Labels{
			"platform": "linux-amd64",
			"type":     "Statusd",
			"version":  "v0.29.0-beta.2",
		})
	// missing version
	labels, err = labelsFromNodeName("StatusIM/android-arm64/go1.11")
	require.NoError(t, err)
	require.Equal(t, labels,
		prom.Labels{
			"platform": "android-arm64",
			"type":     "StatusIM",
			"version":  "unknown",
		})
}

func TestUpdateNodeMetricsPeersCounter(t *testing.T) {
	var err error

	n, err := node.New(&node.Config{
		P2P: p2p.Config{
			MaxPeers: 10,
		},
		NoUSB: true,
	})
	require.NoError(t, err)
	require.NoError(t, n.Start())
	defer func() { require.NoError(t, n.Stop()) }()
	server := n.Server()

	change, err := computeMetrics(server, p2p.PeerEventTypeAdd)
	require.NoError(t, err)
	require.Equal(t, float64(1), change.Counter)
	require.Equal(t, float64(10), change.Max)

	// skip other events
	change, err = computeMetrics(server, p2p.PeerEventTypeMsgRecv)
	require.NoError(t, err)
	require.Equal(t, float64(0), change.Counter)
	change, err = computeMetrics(server, p2p.PeerEventTypeMsgSend)
	require.NoError(t, err)
	require.Equal(t, float64(0), change.Counter)

	change, err = computeMetrics(server, p2p.PeerEventTypeDrop)
	require.NoError(t, err)
	require.Equal(t, float64(-1), change.Counter)

	server.MaxPeers = 20
	change, err = computeMetrics(server, p2p.PeerEventTypeDrop)
	require.NoError(t, err)
	require.Equal(t, float64(20), change.Max)
}
