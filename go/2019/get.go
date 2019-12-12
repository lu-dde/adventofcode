package main

type solver func(chan string, chan string)

func getSolver(name string) solver {
	solver, ok := map[string]solver{
		"11": U11,
		"12": U12,
		"21": U21,
		"22": U22,
		"31": U31,
		"41": U41,
		"42": U42,
		"51": U51,
		"52": U52,
		"61": U61,
		"62": U62,
		"71": U71,
	}[name]

	if !ok {
		panic("could not find a solver function")
	}

	return solver
}

func getTestfile(name string) string {
	textfile, ok := map[string]string{
		"11": "input/u11.txt",
		"12": "input/u11.txt",
		"21": "input/u21.txt",
		"22": "input/u21.txt",
		"31": "input/u31.txt",
		"41": "input/u41.txt",
		"51": "input/u51.txt",
		"52": "input/u51.txt",
		"61": "input/u61.txt",
		"62": "input/u61.txt",
		"71": "input/u71.test.txt",
	}[name]

	if !ok {
		panic("could not find a solver function")
	}

	return textfile
}
