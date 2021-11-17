//go:generate $GOBIN/mockgen -source=$GOFILE -destination=mock/mock_$GOFILE -package=mock_$GOPACKAGE

package datetime

const (
	nullString = "null"
)

type Manager interface {
	Clock() Clock
	Calendar() Calendar
}

func NewManager() Manager {
	return &manager{}
}

type manager struct{}

func (m *manager) Clock() Clock {
	return NewClock()
}

func (m *manager) Calendar() Calendar {
	return NewCalendar()
}
