package DAG

// readGradient TODO Read from File or Transfer between Processes.
func (that *Peer) readGradient(path string) []float64 {
	return []float64{}
}

func (that *Peer) aggregateGradient() *Peer {
	// TODO Accumulative to Threshold. Max Circle, Max Time or Gradient Size.
	return that
}

func (that *Peer) ReleaseModel() *Peer {
	// TODO Calculate Timestamp 'Start'.
	// TODO HE Calculate.
	// TODO Calculate Timestamp 'End'.
	// TODO Broadcast via TCP.
	// TODO 1.Send Id, 2. Confirm Not Reached, 3. Transfer.
	return that
}

func (that *Peer) receiveModel() *Peer {
	// TODO Calculate Timestamp 'Start'.
	// TODO HE Verify.
	// TODO Calculate Timestamp 'End'.
	return that
}

func (that *Peer) verifyModel() *Peer {
	// TODO Communicate between Processes.
	return that
}

// Try Train with Test Data and Get Result.
func (that *Peer) Try() {
	// TODO Communicate between Processes.
	that.verifyModel()
}
