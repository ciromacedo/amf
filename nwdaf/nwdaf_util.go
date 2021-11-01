package nwdaf

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/free5gc/amf/context"
	"golang.org/x/net/http2"
	"log"
	"net"
	"net/http"
)


const NwdafUrl = "http://127.0.0.1:29599";
const RegistrationAcceptApi = "/datacollection/amf-contexts/registration-accept";


func RegistrationAccept(amfUe *context.AmfUe, amfSelf *context.AMFContext) {

	m := map[string]interface{}{}
	amfSelf.EventSubscriptions.Range(func(key, value interface{}) bool {
		m[fmt.Sprint(key)] = value
		return true
	})

	b, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))


	url := getUrlConnection(RegistrationAcceptApi)

	jsonData := `
    {	
		"Amf" : { 	"id"	: "`+ amfUe.ServingAMF().NfId +`", 
					"locale": "`+ amfUe.ServingAMF().Locality +`"},
		"Ue"  : { 	"suci"	: "`+ amfUe.Suci+`", 
					"supi"	: "`+ amfUe.Supi+`"},
		"PlmnId" : { "mcc"	: "`+ amfUe.PlmnId.Mcc+`", 
					 "mnc"	: "`+ amfUe.PlmnId.Mnc+`"}
	}
	`

	var jsonStr = []byte(jsonData)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := getConnection()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}

func getUrlConnection(str string)(string){
	return NwdafUrl +str;
}


func getConnection()(http.Client){
	client := http.Client{
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	return client
}
