package cloudflare

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (connInfo *ConnectInformation) createHTTPClient() (err error) {
	connInfo.HTTPClient = &http.Client{}
	return
}

func (connInfo ConnectInformation) newQuery(req *http.Request) (body []byte, err error) {
	if connInfo.HTTPClient == nil {
		connInfo.createHTTPClient()
	}
	req.Header.Add("X-Auth-Email", connInfo.AuthEmail)
	req.Header.Add("X-Auth-Key", connInfo.AuthKey)
	resp, err := connInfo.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// GetZone - Get all zones and return a struct with all information
func (connInfo ConnectInformation) GetZone(record string) (zoneResponse ZoneResponse, err error) {
	req, err := http.NewRequest("GET", connInfo.URL+"zones?name="+record, nil)
	if err != nil {
		return
	}

	body, err := connInfo.newQuery(req)
	if err == nil {
		err = json.Unmarshal(body, &zoneResponse)
	}
	return
}

// GetRecord - Get one record and return a struct with that information
func (connInfo ConnectInformation) GetRecord(recordType string, record string) (dnsRecordResponse DNSRecordsResponse, err error) {
	req, err := http.NewRequest("GET", connInfo.URL+"zones/"+connInfo.ZoneIdentifier+"/dns_records?type="+recordType+"&name="+record, nil)
	if err != nil {
		return
	}

	body, err := connInfo.newQuery(req)
	if err == nil {
		err = json.Unmarshal(body, &dnsRecordResponse)
	}
	return
}

// GetAllRecords - Get all records in a zone and return a struct
func (connInfo ConnectInformation) GetAllRecords() (dnsRecordResponse DNSRecordsResponse, err error) {
	req, err := http.NewRequest("GET", connInfo.URL+"zones/"+connInfo.ZoneIdentifier+"/dns_records", nil)
	if err != nil {
		return
	}

	body, err := connInfo.newQuery(req)
	if err == nil {
		err = json.Unmarshal(body, &dnsRecordResponse)
	}
	return
}

// CreateRecord - Create a DNS record and return the response in a struct
func (connInfo ConnectInformation) CreateRecord(record CreateDNSRecord) (dnsRecordResponse DNSRecordResponse, err error) {
	postData, _ := json.Marshal(record)
	req, err := http.NewRequest("POST", connInfo.URL+"zones/"+connInfo.ZoneIdentifier+"/dns_records", bytes.NewReader(postData))
	if err != nil {
		return
	}

	body, err := connInfo.newQuery(req)
	if err == nil {
		err = json.Unmarshal(body, &dnsRecordResponse)
	}
	return
}

// UpdateRecord - Update DNS record and return the response in a struct 
func (connInfo ConnectInformation) UpdateRecord(record DNSRecord) (dnsRecordResponse DNSRecordResponse, err error) {
	putData, err := json.Marshal(record)
	if err != nil {
		return
	}
	req, err := http.NewRequest("PUT", connInfo.URL+"zones/"+connInfo.ZoneIdentifier+"/dns_records/"+record.ID, bytes.NewReader(putData))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return
	}

	body, err := connInfo.newQuery(req)
	if err == nil {
		err = json.Unmarshal(body, &dnsRecordResponse)
	}
	return
}
