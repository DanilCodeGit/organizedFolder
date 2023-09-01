package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Функция для проверки наличия ошибки
func check(err error) {
	if err != nil {
		fmt.Printf("Произошла ошибка: %s \n", err)
		os.Exit(1)
	}
}

// Функция для создания стандартных папок, таких как Images, Music, Docs, Others, Videos
func createDefaultFolders(targetFolder string) {
	defaultFolders := []string{"Music", "Videos", "Docs", "Images", "Others"}

	for _, folder := range defaultFolders {
		_, err := os.Stat(folder)
		if os.IsNotExist(err) {
			// Создаем папку, если она не существует
			os.Mkdir(filepath.Join(targetFolder, folder), 0755)
		}
	}
}

// Функция для организации файлов
func organizeFolder(targetFolder string) {
	// Считываем содержимое директории
	filesAndFolders, err := os.ReadDir(targetFolder)
	check(err)

	// Счетчик перемещенных файлов
	noOfFiles := 0

	for _, filesAndFolder := range filesAndFolders {
		// Проверяем, является ли элемент файлом (а не директорией)
		if !filesAndFolder.IsDir() {
			fileInfo, err := filesAndFolder.Info()
			check(err)

			// Получаем полный путь к файлу
			oldPath := filepath.Join(targetFolder, fileInfo.Name())
			fileExt := filepath.Ext(oldPath)

			// Используем конструкцию switch для перемещения файлов в зависимости от их расширения
			switch fileExt {
			case ".png", ".jpg", ".jpeg":
				newPath := filepath.Join(targetFolder, "Images", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp4", ".mov", ".avi", ".amv":
				newPath := filepath.Join(targetFolder, "Videos", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".pdf", ".docx", ".csv", ".xlsx":
				newPath := filepath.Join(targetFolder, "Docs", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			case ".mp3", ".wav", ".aac":
				newPath := filepath.Join(targetFolder, "Music", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			default:
				newPath := filepath.Join(targetFolder, "Others", fileInfo.Name())
				err = os.Rename(oldPath, newPath)
				check(err)
				noOfFiles++
			}
		}
	}

	// Выводим информацию о количестве перемещенных файлов
	if noOfFiles > 0 {
		fmt.Printf("%v файлов перемещено\n", noOfFiles)
	} else {
		fmt.Printf("Файлы не перемещены")
	}
}

// Основная функция
func main() {

	// Получаем ввод пользователя - целевая папка, которую нужно организовать
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Какую папку вы хотите организовать? - ")
	scanner.Scan()

	targetFolder := scanner.Text()

	// Проверяем существование папки
	_, err := os.Stat(targetFolder)
	if os.IsNotExist(err) {
		fmt.Println("Папка не существует.")
		os.Exit(1)
	} else {
		// Создаем стандартные папки, такие как Images, Music, Docs, Others, Videos
		createDefaultFolders(targetFolder)

		// Организуем файлы
		organizeFolder(targetFolder)
	}
}
