package main

type solver func(chan string, chan string)

func getSolver(name string) solver {
	functions := map[string]solver{
		"11": U11,
	}

	return functions[name]
}

func getTestfile(name string) string {
	textfiles := map[string]string{
		"11": "input/u11.txt",
	}

	return textfiles[name]
}
