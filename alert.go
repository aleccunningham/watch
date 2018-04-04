package watch

const (
	Alert0 AlertType = "ALERT_0"
	Alert1 AlertType = "ALERT_1"
	Alert2 AlertType = "ALERT_2"
	Alert3 AlertType = "ALERT_3"
	Alert4 AlertType = "ALERT_4"
	Alert5 AlertType = "ALERT_5"
	Alert6 AlertType = "ALERT_6"
	Alert7 AlertType = "ALERT_7"
	Alert8 AlertType = "ALERT_8"
	Alert9 AlertType = "ALERT_9"
)

var AlertType struct {
	level int
}

func Alert(alert AlertType) (Event, error) {
	alertEvent := NewAlertEvent(alert)

	return alertEvent
}

func NewAlertEvent(alert AlertType) AlertType {
	return &AlertType{
		level: alert,
	}
}
