package test

func InitializeTestStorages() (tearDown func(), retErr error) {
	conf := NewTestConfig()

	var tearDowns []func()
	tf := func() {
		if len(tearDowns) == 0 {
			return
		}
		for _, t := range tearDowns {
			t()
		}
	}

	defer func() {
		if retErr != nil {
			tf()
			tearDown = nil
		}
	}()

	tearDownDB, err := InitializeTestDB(conf)
	if err != nil {
		return nil, err
	}
	tearDowns = append(tearDowns, tearDownDB)

	return tf, nil

}
