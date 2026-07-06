package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type Result struct {
	Domain      string `json:"domain"`
	HasMX       bool   `json:"hasMx"`
	HasSPF      bool   `json:"hasSPF"`
	SPFRecord   string `json:"spfRecord"`
	HasDMARC    bool   `json:"hasDMARC"`
	DMARCRecord string `json:"dmarcRecord"`
}

func main() {

	http.HandleFunc("/check", checkHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func checkHandler(w http.ResponseWriter, r *http.Request) {

	domain := r.URL.Query().Get("domain")
	if domain == "" {
		http.Error(w, "domain is required", http.StatusBadRequest)
		return
	}

	hasMx, hasSPF, spf, hasDMARC, dmarc := checkDomain(domain)

	result := Result{
		Domain:      domain,
		HasMX:       hasMx,
		HasSPF:      hasSPF,
		SPFRecord:   spf,
		HasDMARC:    hasDMARC,
		DMARCRecord: dmarc,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func checkDomain(domain string) (bool, bool, string, bool, string) {

	var hasMx, hasSPF, hasDMARC bool
	var sprRecords, demarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("MX Error: %v", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("TXT Error: %v", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			sprRecords = record
			break
		}
	}

	demarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("DMARC Error: %v", err)
	}

	for _, record := range demarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			demarcRecord = record
			break
		}
	}

	return hasMx, hasSPF, sprRecords, hasDMARC, demarcRecord
}
