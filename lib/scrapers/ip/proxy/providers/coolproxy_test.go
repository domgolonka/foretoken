package providers

import (
	"testing"
)

var c = NewCoolProxy()
var raw = []byte(`
<div id="main">
  <table>
    <tr>
      <th class="tHeader" title="Ip Address of Proxy server. Use together with Port"><a href="/proxies/http_proxy_list/sort:ip/direction:asc">IP Address</a></th>
      <th class="tHeader" title="Port of Proxy server. Should be used together with IP"><a href="/proxies/http_proxy_list/sort:port/direction:asc">Port</a></th>
      <th class="tHeader" title="Location of Proxy server"><a href="/proxies/http_proxy_list/sort:country_code/direction:asc">Flag</a></th>
      <th class="tHeader" title="Location of Proxy server"><a href="/proxies/http_proxy_list/sort:country_name/direction:asc">Country</a></th>
      <th class="tHeader" title="How many stars did we give to this Proxy server. Take proxies with at least 2 stars"><a href="/proxies/http_proxy_list/sort:score/direction:asc" class="desc">Rating</a></th>
      <th class="tHeader" title="Will you be anonymous if you use this Proxy server"><a href="/proxies/http_proxy_list/sort:anonymous/direction:desc">Anonymous</a></th>
      <th class="tHeader" title="Percentage of time Proxy server was available"><a href="/proxies/http_proxy_list/sort:working_average/direction:desc">Working (%)</a></th>
      <th class="tHeader" title="The time it took for the Proxy server to respond in seconds"><a href="/proxies/http_proxy_list/sort:response_time_average/direction:asc">Response Time (s)</a></th>
      <th class="tHeader" title="Average download speed of this Proxy Server"><a href="/proxies/http_proxy_list/sort:download_speed_average/direction:desc">Download Speed (KB/s)</a></th>
      <th class="tHeader" title="How many minutes:seconds ago did we check this Proxy server"><a href="/proxies/http_proxy_list/sort:update_time/direction:desc">Last Check</a></th>
    </tr>
    <!-- Here is where we loop through our $posts array, printing out post info -->
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGZ4YwL4YwRlZP4lZQR=")))</script>138.68.120.201
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:99.138626%;background:			        
          #0167c0;">0.09</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">1029</span></div>
      </td>
      <td class="time">80:38</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("AwphZwN1YwRmZv4lAQR=")))</script>67.205.132.241
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:98.72776%;background:			        
          #0266bf;">0.13</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">151</span></div>
      </td>
      <td class="time">71:31</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZwN3YwR1AP4lZmRhZwRm")))</script>207.154.231.213
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:98.17561%;background:			        
          #0366bf;">0.18</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">1165</span></div>
      </td>
      <td class="time">70:45</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZwN3YwR1AP4lZmRhZwN4")))</script>207.154.231.208
      </td>
      <td>8080</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:97.23378%;background:			        
          #0565be;">0.28</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">1766</span></div>
      </td>
      <td class="time">71:41</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGp0YwRmBP41AP40BD==")))</script>174.138.54.49
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:96.83904%;background:			        
          #0664bd;">0.32</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">142</span></div>
      </td>
      <td class="time">70:30</td>
    </tr>
    <tr>
      <td colspan="10">
        <script type="text/javascript"><!--
          google_ad_client = "ca-pub-4076174243377816";
          google_ad_slot = "5199943694";
          google_ad_width = 728;
          google_ad_height = 90;
          //-->
        </script>
        <script type="text/javascript" src="//pagead2.googlesyndication.com/pagead/show_ads.js"></script>
      </td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("AwphZwN1YwR0Av43Aj==")))</script>67.205.146.77
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:95.55247%;background:			        
          #0863bc;">0.44</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">139</span></div>
      </td>
      <td class="time">68:48</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGH5YwVjZl4kAwRhZGD0")))</script>159.203.161.144
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:95.30818%;background:			        
          #0963bc;">0.47</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">148</span></div>
      </td>
      <td class="time">67:24</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGZ5Ywx5Ywx5YwR1")))</script>139.99.99.15
      </td>
      <td>8888</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:95.28015%;background:			        
          #0963bc;">0.47</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">125</span></div>
      </td>
      <td class="time">71:32</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("BGHhBQHhAGthZGH0")))</script>95.85.58.154
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-nl" alt="Netherlands flag"></td>
      <td>Netherlands</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:95.17411%;background:			        
          #0963bc;">0.48</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">752</span></div>
      </td>
      <td class="time">84:54</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("AwVhAl44AF4lZmD=")))</script>62.7.85.234
      </td>
      <td>8080</td>
      <td><img src="/img/spacer.gif" class="flag flag-gb" alt="United Kingdom flag"></td>
      <td>United Kingdom</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:95.03737%;background:			        
          #0962bb;">0.5</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">542</span></div>
      </td>
      <td class="time">92:52</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZwN3YwR1AP4lZmRhZwRk")))</script>207.154.231.211
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:94.3324%;background:			        
          #0a62bb;">0.57</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">1419</span></div>
      </td>
      <td class="time">69:25</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("AGDhZmphZGZkYwR1")))</script>54.37.131.15
      </td>
      <td>8080</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:94.30769%;background:			        
          #0a62bb;">0.57</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">1062</span></div>
      </td>
      <td class="time">80:53</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGt1YwRkBP4lAv4kZN==")))</script>185.118.26.10
      </td>
      <td>80</td>
      <td><img src="/img/spacer.gif" class="flag flag---" alt="Unknown flag"></td>
      <td>Unknown</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>No</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:93.77083%;background:			        
          #0c61ba;">0.62</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">141020</span></div>
      </td>
      <td class="time">64:13</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGZ4YwL4YwR2ZF42ZN==")))</script>138.68.161.60
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:93.50701%;background:			        
          #0c61ba;">0.65</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">227</span></div>
      </td>
      <td class="time">64:31</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGN0YwV0BP4kZwDhZwL=")))</script>104.248.124.26
      </td>
      <td>8080</td>
      <td><img src="/img/spacer.gif" class="flag flag---" alt="Unknown flag"></td>
      <td>Unknown</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:93.38901%;background:			        
          #0c61ba;">0.66</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">167</span></div>
      </td>
      <td class="time">87:36</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZwRlYwVmZl4kZGxhAQV=")))</script>212.233.119.42
      </td>
      <td>60379</td>
      <td><img src="/img/spacer.gif" class="flag flag-ru" alt="Russian Federation flag"></td>
      <td>Russian Federation</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:93.32185%;background:			        
          #0c61ba;">0.67</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">241</span></div>
      </td>
      <td class="time">68:24</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGp0YwRmBP40BP4kZmD=")))</script>174.138.48.134
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:93.2192%;background:			        
          #0d61ba;">0.68</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">145</span></div>
      </td>
      <td class="time">83:45</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGN0YwV0BP42ZF4lZGR=")))</script>104.248.61.211
      </td>
      <td>80</td>
      <td><img src="/img/spacer.gif" class="flag flag---" alt="Unknown flag"></td>
      <td>Unknown</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:95.45277%;background:			        
          #0863bc;">0.45</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">234</span></div>
      </td>
      <td class="time">78:28</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGDlYwxmYwL2YwH5")))</script>142.93.66.59
      </td>
      <td>8080</td>
      <td><img src="/img/spacer.gif" class="flag flag-ca" alt="Canada flag"></td>
      <td>Canada</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:92.8863%;background:			        
          #0d60b9;">0.71</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">166</span></div>
      </td>
      <td class="time">84:53</td>
    </tr>
    <tr>
      <td style="text-align:left; font-weight:bold;">
        <script type="text/javascript">document.write(Base64.decode(str_rot13("ZGZ4YwL4YwR2ZF4kAGp=")))</script>138.68.161.157
      </td>
      <td>3128</td>
      <td><img src="/img/spacer.gif" class="flag flag-us" alt="United States flag"></td>
      <td>United States</td>
      <td><img src="/img/stars/5stars1.png" alt="5 star proxy"></td>
      <td>Anonymous</td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0067c0;">100</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:92.70892%;background:			        
          #0e60b9;">0.73</span></div>
      </td>
      <td style="text-align:left;">
        <div class="graph"><span class="bar" style="color:#fff; width:100%;background:
          #0068c1;">198</span></div>
      </td>
      <td class="time">93:07</td>
    </tr>
    <tr>
      <th colspan="10" class="pagination">
        <span class="prev">&lt;&lt; Prev</span>&nbsp;&nbsp;
        <span class="current">1</span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:2">2</a></span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:3">3</a></span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:4">4</a></span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:5">5</a></span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:6">6</a></span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:7">7</a></span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:8">8</a></span> | <span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:9">9</a></span>...<span><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:169">169</a></span>&nbsp;&nbsp;
        <span class="next"><a href="/proxies/http_proxy_list/sort:score/direction:desc/page:2" rel="next">Next &gt;&gt;</a></span>			            			        
      </th>
    </tr>
  </table>
</div> 
`)

func TestCoolProxy_Load(t *testing.T) {
	ips, err := c.Load(raw)
	if err != nil {
		t.Fatal(err)
	}

	if len(ips) != 20 {
		t.Fatalf("expected ips len to be %d, but got %d", 20, len(ips))
	}
}

func BenchmarkCoolProxy_Load(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := c.Load(raw)
		if err != nil {
			b.Fatal(err)
		}
	}
}
