package jobapp

func InitJob() {
	listJob := []func(){
		NewEmailJob().handle,
	}

	for _, j := range listJob {
		go func(job func()) {
			job()
		}(j)
	}
}
