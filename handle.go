package watch

const (
		Alert EventType = "ALERT_"
)

var EventType string


func Handler(Event) error {
		eventType := getEventType(*Event)

		switch {
		case eventType == Alert:
				handleAlert, error := handleAlert(Alert)
				if err != nil {
						panic("failed to broadcast alert")
				}
		case default:
				handle, error := handleEvent(Event)
				if err != nil {
						panic("failed to handle event")
				}
		}
}
