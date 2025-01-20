package workerpool

type CsvWorker struct {
	GlobalRank     int    `json:"globalrank"`
	TldRank        int    `json:"tldrank"`
	Domain         string `json:"domain"`
	Tld            string `json:"tld"`
	RefSubNets     int    `json:"refsubnets"`
	RefIPs         int    `json:"refips"`
	IDN_Domain     string `json:"idn_domain"`
	IDN_TLD        string `json:"idn_tld"`
	PrevGlobalRank int    `json:"prevglobalrank"`
	PrevTldRank    int    `json:"prevtldrand"`
	PrevRefSubNets int    `json:"prevrefsubnets"`
	PrevRefIPs     int    `json:"prevrefips"`
}
