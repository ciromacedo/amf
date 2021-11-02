package producer

import (
	"bytes"
	"github.com/free5gc/amf/context"
	"github.com/free5gc/amf/nwdaf"
	"log"
	"net/http"
)

func RegistrationAccept(amfUe *context.AmfUe, amfSelf *context.AMFContext) {

	/* verifica se existe registro de instância NWDAF */
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

		client := nwdaf.GetConnection()

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()
	}
}