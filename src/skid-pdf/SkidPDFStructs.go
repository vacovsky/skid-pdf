package main

type pdfRequest struct {
	URL            string `json:"url" form:"url"`
	Grayscale      bool   `json:"grayscale" form:"grayscale"`
	Landscape      bool   `json:"landscape" form:"landscape"`
	TargetFileName string `json:"targetFileName" form:"targetFileName"`
	TargetFileDest string `json:"targetFileDest" form:"targetFileDest"`
	Action         string `json:"action" form:"action"`
	Data           string `json:"data" form:"data"`
}
