package providers_test

import (
	"testing"

	"github.com/soluchok/freeproxy/providers"
)

var fp = providers.NewFreeProxyList()
var fpBody = []byte(`
<div id="proxylisttable_wrapper" class="dataTables_wrapper">
   <div class="row">
      <div class="col-sm-12">
         <table class="table table-striped table-bordered dataTable" cellspacing="0" width="100%" id="proxylisttable" role="grid" aria-describedby="proxylisttable_info" style="width: 100%;">
            <tbody>
               <tr role="row" class="odd">
                  <td>89.187.217.118</td>
                  <td>80</td>
               </tr>
               <tr role="row" class="even">
                  <td>203.74.4.2</td>
                  <td>80</td>
               </tr>
               <tr role="row" class="odd">
                  <td>203.74.4.7</td>
                  <td>80</td>
               </tr>
               <tr role="row" class="even">
                  <td>203.74.4.3</td>
                  <td>80</td>
               </tr>
            </tbody>
         </table>
      </div>
   </div>
</div>`)

func TestFreeProxyListLoad(t *testing.T) {
	ips, err := fp.Load(fpBody)
	if err != nil {
		t.Fatalf("Load is fail %s", err.Error())
	}
	if len(ips) != 4 {
		t.Fatalf("Len ips is %d must be 4", len(ips))
	}
}

func BenchmarkFreeProxyListLoad(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fp.Load(fpBody)
	}
}
