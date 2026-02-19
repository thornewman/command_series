package lib

type MessageSync struct {
	update chan any
	cont   chan bool
}

func NewMessageSync() *MessageSync {
	return &MessageSync{
		update: make(chan any),
		cont:   make(chan bool)}
}

func (s *MessageSync) SendUpdate(msg any) bool {
	s.update <- msg
	return <-s.cont
}
func (s *MessageSync) Wait() bool {
	return <-s.cont
}
func (s *MessageSync) GetUpdate() any {
	s.cont <- true
	return <-s.update
}
func (s *MessageSync) Stop() {
	s.cont <- false
}
