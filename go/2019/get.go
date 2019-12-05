package main

type solver func(chan string, chan string)

func getSolver(name string) solver {
	functions := map[string]solver{
		"11": U11,
		"12": U12,
		"21": U21,
		"22": U22,
		"31": U31,
	}

	return functions[name]
}

func getTestfile(name string) string {
	textfiles := map[string]string{
		"11": "input/u11.txt",
		"12": "input/u11.txt",
		"21": "input/u21.txt",
		"22": "input/u21.txt",
		"31": "input/u31.txt",
	}

	return textfiles[name]
}
