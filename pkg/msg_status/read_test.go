package msg_status

import (
	"bytes"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRead(t *testing.T) {
	bs, err := os.ReadFile(filepath.Join("testdata", "msg-status-recv.xml"))
	require.NoError(t, err)

	var msgStatus Reader
	err = xml.NewDecoder(bytes.NewReader(bs)).Decode(&msgStatus)
	require.NoError(t, err)

	require.Equal(t, "M20230222200000057A1BFFF87758970177", msgStatus.MsgId())
	require.Equal(t, "M20230222234567891T1BOTS01850058244", msgStatus.OrgnlMsgId())
	require.Equal(t, "20230222200000057A1B123856846970177", msgStatus.OrgnlInstrId())
	require.Equal(t, "20230222200000057A1B123856846970177", msgStatus.OrgnlTxId())
	require.Equal(t, "RCVD", msgStatus.TxSts())
	require.Equal(t, "", msgStatus.StsRsn())
	require.Equal(t, 0.0, msgStatus.IntrBkSttlmAmt())
	require.Equal(t, "0001-01-01T00:00:00Z", msgStatus.IntrBkSttlmDt().Format(time.RFC3339))
}
