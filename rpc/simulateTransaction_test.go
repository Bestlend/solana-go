package rpc

import (
	"context"
	"fmt"
	"github.com/Bestlend/solana-go"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_SimulateTransactionWithOpts(t *testing.T) {
	t.Skip()
	client := New(MainNetBeta_RPC)

	txData := "AZss8pCBB5FQjmE+y8F9j44UWOZJ30UOeVpVOj6MPrzz6AicwKY4rdmq1xc7ByCeWAzKriTZkY+Q7msFy4CfegiAAQADDAw1GjRlHOLKPlEoqeV3VmVCvhzuXs3CNEJ1iyO6X0wbsdiswKnfXPKs7dPU3Kl6QOcadUtfhfFMXn3IhYZ4KzHQmZr1rJDtclN66upO9QSMHI4vf7LTTsWl50j1u/O3iwR4MSYHL0KpzWat0Y0bh7sxC/WDC8z1X2ot8XwBHNQM7fG2+iH3t5gWGtqvv2kSAeuvExP26KXeM6eEUp/5rz6qJkX9VXz7Kq5f18qks/qugGJSguKeYM2ZBzd6vgU9rFe4BFSYF6bt6iQkL2zVIe0Cw7ImZFTuMhtwhvtwoYV+iIgylC2bo4fyZdKChvPiNXrXTjnGWdbeC/n2vvL+VogUDAnDiGwcaoRPMpMbzwSqwxeFGbHA/lPU/ypxDrTePAR51VvyMcBu7nTFbs5oFQf9sbLeo/SOUQKxzaJWvBOPAwZGb+UhFzL/7K26csOb57yM5bvF9xJrLEObOkAAAAAI4JNdDrH8DXYq2tEruBOIIu3b7wrVZMYIgFlCyxPHHPHZR512vnfEiX3/ZiWJNIrTEgshc7xK8nrntkSJ1kYQBwoACQMIGggAAAAAAAoABQKAuSoACwUAECABIhThFA86Cm7jOwAAAAAAAAAAAAAAAAsVACEQEi4iIyQlARMUFRMmFycoKSorEEtGzLJBU4uSgIQeAAAAAAAJJS4IABICAxYjLAkJLQkvLggZBBoCGwUGBzAdHgwfDQ4IBAMPHC4pwSCbM0HWnIEBAgAAABEAZAABMGQBAoCEHgAAAAAAaH0eAAAAAAAZAAALDAAhEBYRLiwkJQEXGBFqxTFULbajzmh9HgAAAAAAAAsFABAgASIIXIf8Uj+QVE4ED426La5P7U2ZI5iKXZ5DR8PMNGlo1ELnW1nSjqdO0g8EODI2IgMQFDPegtlf4l9tXF0mNS98e2hau8RUdHGgJkGfJAJQTuiR+QICCAIDAdQePuO6Ce1veEMhyp5PD1/PFN246RRaZk2E7soZfJRFByd8pZgmtH4MHg0DBHNuk02DsAw18hrkWGR9PR33RP5C/L+kJfh9N/rTvvZTJplYPDtXOoAEfXl68AMCd34="
	tx, err := solana.TransactionFromBase64(txData)
	require.NoError(t, err)

	out, err := client.SimulateTransactionWithOpts(
		context.Background(), tx, &SimulateTransactionOpts{
			SigVerify:              false,
			InnerInstructions:      true,
			ReplaceRecentBlockhash: true,
		},
	)
	require.NoError(t, err)

	for _, innerIx := range out.Value.InnerInstructions {
		fmt.Printf("%+v\n", innerIx.Instructions)
	}
}
