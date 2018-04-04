package handle

const (
		Alert EventType = "ALERT_"
)

var EventType string

func Handler(Event) error {
		eventType := getEventType(*Event)

		switch {
		case eventType == Alert + "0":
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

func handleEvent(Event) (Result, error) {
	return nil, nil 
}
