package main

func main() {
	app := newApp()
	if _, err := app.Run(); err != nil {
		panic(err)
	}
}
