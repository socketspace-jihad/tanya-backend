package notification

type Notification struct {
	Title    string            `json:"title"`
	Icon     string            `json:"icon"`
	Subtitle string            `json:"subtitle"`
	Topic    string            `json:"topic"`
	Data     map[string]string `json:"data"`
}

type Notifier interface {
	Notify(*Notification) error
}
