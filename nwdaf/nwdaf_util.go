package nwdaf

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/free5gc/amf/context"
	"golang.org/x/net/http2"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)


const NwdafUrl = "http://127.0.0.1:29599";
const RegistrationAcceptApi = "/amf-contexts/registration-accept";
const DateFormat = "02/Jan/2006 15:04:05";

func RegistrationAccept(amfUe *context.AmfUe){
	url := getUrlCoonection(RegistrationAcceptApi)

	fmt.Println(amfUe.ServingAMF().NfId)

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

	//criando o JSOn
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

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func getUrlCoonection(str string)(string){
	return NwdafUrl +str;
}




func getConnection()(http.Client){
	client := http.Client{
		Transport: &http2.Transport{
			// So http2.Transport doesn't complain the URL scheme isn't 'https'
			AllowHTTP: true,
			// Pretend we are dialing a TLS endpoint. (Note, we ignore the passed tls.Config)
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}
	return client
}

func NWDAFConnection() {
	url := NwdafUrl +"/amf-contexts/registration-accept";

	//criando o JSOn
	var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := getConnection()

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}