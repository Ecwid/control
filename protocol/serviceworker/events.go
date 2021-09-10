package serviceworker

/*

 */
type WorkerErrorReported struct {
	ErrorMessage *ServiceWorkerErrorMessage `json:"errorMessage"`
}

/*

 */
type WorkerRegistrationUpdated struct {
	Registrations []*ServiceWorkerRegistration `json:"registrations"`
}

/*

 */
type WorkerVersionUpdated struct {
	Versions []*ServiceWorkerVersion `json:"versions"`
}
