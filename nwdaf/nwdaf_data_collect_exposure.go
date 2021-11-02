package nwdaf

import (
	"bytes"
	"github.com/free5gc/amf/context"
	"log"
	"net/http"
)

func RegistrationAccept(amfUe *context.AmfUe, amfSelf *context.AMFContext) {

	/* verifica se existe registro de instância NWDAF */
	nwdafSubscriber, _ := amfSelf.FindEventSubscription("NWDAF")
	if nwdafSubscriber != nil {
		/* verifica se o evento será interceptado */
		eventList := nwdafSubscriber.EventSubscription.EventList
		if IsEventIntercept(eventList, "REGISTRATION_ACCEPT") == false {
			return
		}

		url := nwdafSubscriber.EventSubscription.EventNotifyUri
		jsonData := BuildRegistrationAcceptJson(amfUe)

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

func BuildRegistrationAcceptJson(amfUe *context.AmfUe) string {
	return `
    	{	
		"Amf" : { 	"id"	: "` + amfUe.ServingAMF().NfId + `", 
					"locale": "` + amfUe.ServingAMF().Locality + `"},
		"Ue"  : { 	"suci"	: "` + amfUe.Suci + `", 
					"supi"	: "` + amfUe.Supi + `"},
		"PlmnId" : { "mcc"	: "` + amfUe.PlmnId.Mcc + `", 
					 "mnc"	: "` + amfUe.PlmnId.Mnc + `"}
		}
		`
}