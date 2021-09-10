package animation

/*
	Event for when an animation has been cancelled.
*/
type AnimationCanceled struct {
	Id string `json:"id"`
}

/*
	Event for each animation that has been created.
*/
type AnimationCreated struct {
	Id string `json:"id"`
}

/*
	Event for animation that has been started.
*/
type AnimationStarted struct {
	Animation *Animation `json:"animation"`
}
