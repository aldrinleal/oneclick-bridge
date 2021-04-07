package types

import "time"

type OneClickEvent struct {
	DeviceEvent struct {
		ButtonClicked struct {
			AdditionalInfo struct {
				Version string `json:"version"`
			} `json:"additionalInfo"`
			ClickType    string    `json:"clickType"`
			ReportedTime time.Time `json:"reportedTime"`
		} `json:"buttonClicked"`
	} `json:"deviceEvent"`
	DeviceInfo struct {
		Attributes struct {
			DeviceTemplateName string `json:"deviceTemplateName"`
			PlacementName      string `json:"placementName"`
			ProjectName        string `json:"projectName"`
			ProjectRegion      string `json:"projectRegion"`
		} `json:"attributes"`
		DeviceID      string `json:"deviceId"`
		RemainingLife int64  `json:"remainingLife"`
		Type          string `json:"type"`
	} `json:"deviceInfo"`
	DevicePayload struct {
		CertificateID string `json:"certificateId"`
		ClickType     string `json:"clickType"`
		RemainingLife int64  `json:"remainingLife"`
		ReportedTime  int64  `json:"reportedTime"`
		SerialNumber  string `json:"serialNumber"`
		Topic         string `json:"topic"`
		Version       string `json:"version"`
	} `json:"devicePayload"`
	PlacementInfo struct {
		Attributes struct {
			Location string `json:"location"`
		} `json:"attributes"`
		Devices struct {
			OneClickRequest string `json:"oneClickRequest"`
		} `json:"devices"`
		PlacementName string `json:"placementName"`
		ProjectName   string `json:"projectName"`
	} `json:"placementInfo"`
}
