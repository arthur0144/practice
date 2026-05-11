package main

import (
	"bufio"
	"fmt"
	mathrand "math/rand"
	"os"
	"path/filepath"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	root  *os.Root
	files map[byte]*os.File
	count int
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	return c.count
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	ch := name[0]
	writer := bufio.NewWriter(c.files[ch])
	writer.WriteString(name)
	writer.WriteByte('\n')

	err := writer.Flush()
	if err != nil {
		panic(err)
	}
	c.count++
}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
	for _, f := range c.files {
		f.Close()
	}
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	path := filepath.FromSlash("./census")
	root, err := os.OpenRoot(path)
	if err != nil {
		panic(err)
	}

	files := make(map[byte]*os.File)
	for _, ch := range alphabet {
		f, err := root.Create(string(ch))
		if err != nil {
			panic(err)
		}
		files[byte(ch)] = f
	}

	return &Census{
		root:  root,
		files: files,
		count: 0,
	}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

var rand = mathrand.New(mathrand.NewSource(0))

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
}
