package providers_test

import (
	"testing"

	"github.com/soluchok/freeproxy/providers"
)

var decodeParams = "a=0;s=1;d=2;f=3;g=4;h=5;j=6;k=7;l=8;z=9;"
var mapDecodeParams = map[byte]byte{115: 49, 103: 52, 106: 54, 107: 55, 108: 56, 97: 48, 100: 50, 102: 51, 104: 53, 122: 57}
var xs = providers.NewXseoIn()
var body = []byte(`
<script type="text/javascript">h=0;s=1;d=2;z=3;c=4;g=5;p=6;o=7;m=8;u=9;</script>
<tr class=cls81>
	<td colspan=1>
	   <font class=cls1>
		  47.89.53.92<font class=cls4>:</font><script type="text/javascript">document.write(""+z+s+d+m)</script>
	   </font>
	</td>
</tr>
<tr class=cls8>
	<td colspan=1>
	   <font class=cls1>
		  203.188.225.174<font class=cls4>:</font><script type="text/javascript">document.write(""+m+h+m+h)</script>
	   </font>
	</td>
</tr>
<tr class=cls81>
	<td colspan=1>
	   <font class=cls1>
		  138.201.240.238<font class=cls4>:</font><script type="text/javascript">document.write(""+m+h)</script>
	   </font>
	</td>
</tr>
<tr class=cls8>
	<td colspan=1>
	   <font class=cls1>
		  63.85.203.16<font class=cls4>:</font><script type="text/javascript">document.write(""+m+h+m+h)</script>
	   </font>
	</td>
</tr>`)

func TestXseoInDecodePort(t *testing.T) {
	port := xs.DecodePort(mapDecodeParams, "+a+s+d+f")
	if port == nil {
		t.Fatal("Port is nil")
	}
	if string(port) != "0123" {
		t.Fatalf("Post is %s must be 0123", string(port))
	}

	port = xs.DecodePort(mapDecodeParams, "+g+h+j+k")
	if port == nil {
		t.Fatal("Port is nil")
	}
	if string(port) != "4567" {
		t.Fatalf("Post is %s must be 4567", string(port))
	}

	port = xs.DecodePort(mapDecodeParams, "+j+k+l+z")
	if port == nil {
		t.Fatal("Port is nil")
	}
	if string(port) != "6789" {
		t.Fatalf("Post is %s must be 6789", string(port))
	}

	port = xs.DecodePort(mapDecodeParams, "+j+k")
	if port == nil {
		t.Fatal("Port is nil")
	}
	if string(port) != "67" {
		t.Fatalf("Post is %s must be 67", string(port))
	}
}

func TestXseoInLoad(t *testing.T) {
	ips, err := xs.Load(body)
	if err != nil {
		t.Fatalf("Load is fail %s", err.Error())
	}
	if len(ips) != 4 {
		t.Fatalf("Len ips is %d must be 4", len(ips))
	}
}

func BenchmarkXseoInDecodeParamsToMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xs.DecodeParamsToMap(decodeParams)
	}
}

func BenchmarkXseoInDecodePort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xs.DecodePort(mapDecodeParams, "+a+z+s+l")
	}
}

func BenchmarkXseoInLoad(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xs.Load(body)
	}
}
