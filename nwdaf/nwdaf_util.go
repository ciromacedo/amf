package nwdaf

import (
	"crypto/tls"
	"github.com/free5gc/openapi/models"
	"golang.org/x/net/http2"
	"net"
	"net/http"
)


//const NwdafUrl = "http://127.0.0.1:29599";
//const RegistrationAcceptApi = "/datacollection/amf-contexts/registration-accept";

/*
func RegistrationAccept(amfUe *context.AmfUe, amfSelf *context.AMFContext) {
	nwdafSubscriber, _ := amfSelf.FindEventSubscription("NWDAF")
	if nwdafSubscriber != nil {
		url := nwdafSubscriber.EventSubscription.EventNotifyUri
		jsonData := `
    {	
		"Amf" : { 	"id"	: "` + amfUe.ServingAMF().NfId + `", 
					"locale": "` + amfUe.ServingAMF().Locality + `"},
		"Ue"  : { 	"suci"	: "` + amfUe.Suci + `", 
					"supi"	: "` + amfUe.Supi + `"},
		"PlmnId" : { "mcc"	: "` + amfUe.PlmnId.Mcc + `", 
					 "mnc"	: "` + amfUe.PlmnId.Mnc + `"}
	}
	`

		var jsonStr = []byte(jsonData)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := GetConnection()

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
	}
}



 func getUrlConnection(str string)(string){
	return NwdafUrl +str;
}
*/

/* este método verifica se o evento consta na lista de interceptação */
func IsEventIntercept(events *[]models.AmfEvent, eventName string) bool {
	for _, element := range *events {
		if string(element.Type) == eventName{
			return true
		}
	}
	return false
}

/* retrona conexão HTTP cm TSL */
func GetConnection()(http.Client){
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
