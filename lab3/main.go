package main

import (
	"fmt"
	"lab3/vfs"
)


func main() {
	vfs := vfs.NewVirtualFileSystem("root")

	docsDir, _ := vfs.CreateDirectory("docs", "/root")


	file, _ := vfs.CreateFile("notes.txt", docsDir.Path())


	data := []byte("I like planes!")
	n, err := vfs.Write("notes.txt", data)
	if err != nil {
		fmt.Println("Błąd zapisu:", err)
	} else {
		fmt.Printf("Zapisano %d bajtów do notes.txt\n", n)
	}

	readable, _ := vfs.Open("notes.txt")

	buf := make([]byte, 1024)
	n, err = readable.Read(buf)
	if err != nil {
		fmt.Println("Błąd odczytu:", err)
	} else {
		fmt.Printf("Odczytano z notes.txt: %s\n", string(buf[:n]))
	}

	readonlyFile, _ := vfs.CreateReadonlyFile("readme.md", docsDir.Path())


	_, err = readonlyFile.Write([]byte("Nie wolno pisać!"))
	if err != nil {
		fmt.Println("Błąd dla readonly: ", err)
	}

	symLink, _ := vfs.CreateSymLink("shortcut", docsDir.Path(), file)

	fmt.Printf("Dowiązanie symboliczne %s -> %s\n", symLink.Path(), file.Path())
	fmt.Printf("------------------------------------------------------\n")
	for _, item := range vfs.ListItems() {
		fmt.Printf("- %s (%s)\n", item.Name(), item.Path())
	}

	fmt.Printf("------------------------------------------------------")
	err = vfs.RemoveItem("notes.txt")
	if err != nil {
		fmt.Println("Błąd usuwania pliku:", err)
	} else {
		fmt.Println("\nPlik 'notes.txt' został usunięty.")
	}

	fmt.Printf("------------------------------------------------------\n")
	for _, item := range vfs.ListItems() {
		fmt.Printf("- %s (%s)\n", item.Name(), item.Path())
	}
}
