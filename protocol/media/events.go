package media

/*
	This can be called multiple times, and can be used to set / override /

remove player properties. A null propValue indicates removal.
*/
type PlayerPropertiesChanged struct {
	PlayerId   PlayerId          `json:"playerId"`
	Properties []*PlayerProperty `json:"properties"`
}

/*
	Send events as a list, allowing them to be batched on the browser for less

congestion. If batched, events must ALWAYS be in chronological order.
*/
type PlayerEventsAdded struct {
	PlayerId PlayerId       `json:"playerId"`
	Events   []*PlayerEvent `json:"events"`
}

/*
Send a list of any messages that need to be delivered.
*/
type PlayerMessagesLogged struct {
	PlayerId PlayerId         `json:"playerId"`
	Messages []*PlayerMessage `json:"messages"`
}

/*
Send a list of any errors that need to be delivered.
*/
type PlayerErrorsRaised struct {
	PlayerId PlayerId       `json:"playerId"`
	Errors   []*PlayerError `json:"errors"`
}

/*
	Called whenever a player is created, or when a new agent joins and receives

a list of active players. If an agent is restored, it will receive the full
list of player ids and all events again.
*/
type PlayersCreated struct {
	Players []PlayerId `json:"players"`
}
