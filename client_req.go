package requestclient

import ( 
   "net/http"
   "log"
   "io/ioutil"
   "encoding/json"
)


type RequestClient struct {
 url string
}

func reqJSON(url string) RequestClient {
 resp, err := http.Get(url)
 if err != nil {
  log.Fatal(err)
 }

 defer resp.Body.Close()

 byteValue, _ := ioutil.ReadAll(resp.Body)

    var result map[string]interface{}
    json.Unmarshal([]byte(byteValue), &result)

    return result
}
