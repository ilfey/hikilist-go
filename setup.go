package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Printf("Enter package name (default: main): ")

	var package_name string
	fmt.Scanln(&package_name)

	if package_name == "" {
		package_name = "main"
	}

	fmt.Println("Setting up project...")

	if err := setPackageName(".", package_name); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	fmt.Printf("Remove markdown files? (Y/n): ")

	var remove_markdown_files string
	fmt.Scanln(&remove_markdown_files)

	if strings.ToLower(remove_markdown_files) != "n" {
		err := removeMarkdownFiles(".")
		if err != nil {
			os.Exit(1)
		}
	}

	fmt.Println("Initializing project...")

	if err := modInit(package_name); err != nil {
		os.Exit(1)
	}

	fmt.Println("Tidying project...")

	if err := modTidy(); err != nil {
		os.Exit(1)
	}

	fmt.Println("Removing setup files...")

	if err := os.Remove("setup.go"); err != nil {
		os.Exit(1)
	}

	fmt.Println("Done!")
}

func modInit(package_name string) error {
	var stderr bytes.Buffer

	cmd := exec.Command("go", "mod", "init", package_name)
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		fmt.Print(stderr.String())

		return err
	}

	return nil
}

func modTidy() error {
	var stderr bytes.Buffer

	cmd := exec.Command("go", "mod", "tidy")

	if err := cmd.Run(); err != nil {
		fmt.Print(stderr.String())

		return err
	}

	return nil
}

func removeMarkdownFiles(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := fmt.Sprintf("%s/%s", path, entry.Name())

		if entry.IsDir() {
			err := removeMarkdownFiles(entryPath)
			if err != nil {
				return err
			}

			continue
		}

		if strings.HasSuffix(entry.Name(), ".md") {
			err := os.Remove(entryPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func setPackageName(path string, package_name string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := fmt.Sprintf("%s/%s", path, entry.Name())

		if entry.IsDir() {
			err := setPackageName(entryPath, package_name)
			if err != nil {
				return err
			}

			continue
		}

		err := replaceStringInFile(entryPath, package_name)
		if err != nil {
			return err
		}
	}

	return nil
}

func replaceStringInFile(path string, package_name string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(content), "<"+"package_name"+">", package_name)

	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
