package entities

type Fight_Results struct {
	Date {}
	Participant1Name string
	Participant2Name string
	FightInfo {}
	Event-Place  Event_Place 
	UnderCard_Events {}
}

type Event_Place struct {
	Stadium_Name string
	Country_Name string
}