package tasks

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

// Интерфейс с описанным методом
type Playing interface {
	Play(song string)
}

//Структура
type LivePlaying struct {
	Band string
}

//Метода структуры, который связан с интерфейсом
func (m LivePlaying) Play(song string) {
	fmt.Printf(m.Band+" is playing: %s\n", song)
}

//Другая структура не имеющая отношения к первой
type Walkman struct {
	Model string
	vol   int
}

// Методы второй структуры
func (p Walkman) SetVolume(i int) {
	p.vol = i
	fmt.Printf("Walkman set volume on: %d\n", p.vol)

}

func (p Walkman) WalkmanPlay(model, song string) {
	p.Model = model
	fmt.Printf("Walkman model: %s is playing: %s\n", model, song)
}

// Структура адаптер

type WalkmanPlaying struct {
	Walkman Walkman
}

//Метод структуры адаптера принадлежащий интерфейсу Playing
func (p WalkmanPlaying) Play(song string) {
	s := fmt.Sprintf("%s", song)
	p.Walkman.SetVolume(p.Walkman.vol)
	p.Walkman.WalkmanPlay(p.Walkman.Model, s)

}

// Метод принимающий на вход интерфейс
func Music(playing Playing) {
	song := "Вот как то так мы сделали адаптер\n"
	playing.Play(song)
}

func Task21() {
	m := LivePlaying{Band: "Guns`n`Roses"}
	Music(m)
	pp := WalkmanPlaying{Walkman: Walkman{
		Model: "Romantika",
		vol:   12,
	}}
	Music(pp)
}
