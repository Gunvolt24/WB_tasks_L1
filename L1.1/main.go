//go:build ignore

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Speak(msg string) string {
	return fmt.Sprintf("%s говорит %s", h.Name, msg)
}

func (h Human) Think(msg string) string {
	return fmt.Sprintf("%s думает о %s", h.Name, msg)
}

func (h *Human) Birthday() {
	h.Age++
}

type Action struct {
	Human
	Dream string
}

func main() {

	a := Action{
		Human: Human{
			Name: "Вася",
			Age:  28,
		},
		Dream: "стать разработчиком на Go",
	}

	fmt.Println("Имя:", a.Name)
	fmt.Printf("%s мечтает %s\n", a.Name, a.Dream)
	fmt.Println(a.Speak("привет, уважаемый проверяющий!"))
	a.Birthday()
	fmt.Printf("Васе вот вот исполнится %d лет\n", a.Age)
	fmt.Println(a.Think("встраивании структур"))
}
