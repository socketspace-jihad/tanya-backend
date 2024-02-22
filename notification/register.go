package notification

import "errors"

type NotifierFactory func() Notifier

var (
	NotifierMap map[string]NotifierFactory = make(map[string]NotifierFactory)
)

func RegisterNotifier(name string, fact NotifierFactory) {
	NotifierMap[name] = fact
}

func GetNotifier(name string) (NotifierFactory, error) {
	if _, ok := NotifierMap[name]; !ok {
		return nil, errors.New("notifier not found")
	}
	return NotifierMap[name], nil
}
