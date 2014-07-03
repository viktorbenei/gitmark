package main

type Configs struct {
}

func (c *Configs) tryConfigFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func readConfigFile() {
	err := tryConfigFile(".gitmarkrc.json")
	if !err {
		return
	}
}
