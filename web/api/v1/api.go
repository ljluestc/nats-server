// AlertmanagerDiscovery has all the active Alertmanager endpoints.
type AlertmanagerDiscovery struct {
	ActiveAlertmanagers  []*AlertmanagerTarget `json:"activeAlertmanagers"`
	DroppedAlertmanagers []*AlertmanagerTarget `json:"droppedAlertmanagers"`
}

// LabelMatchers is a list of matchers for labels that define a series.
type LabelMatchers string
