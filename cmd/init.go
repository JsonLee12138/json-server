package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const (
	githubURL = "https://github.com/JsonLee12138/jsonix-kit"
	giteeURL  = "https://gitee.com/jsonlee_lee/jsonix-kit.git"

	airURL      = "github.com/air-verse/air@latest"
	easyjsonURL = "github.com/mailru/easyjson/...@latest"
	swaggoURL   = "github.com/swaggo/swag/cmd/swag@latest"

	sourcePackage = "jsonix-kit"
)

func initRun(cmd *cobra.Command, args []string) error {
	return utils.TryCatchVoid(func() {
		namePrompt := promptui.Prompt{
			Label: "Please enter the project name",
			Validate: func(input string) error {
				if input == "" {
					return errors.New("project name cannot be empty")
				}
				return nil
			},
		}
		name, err := namePrompt.Run()
		if err != nil {
			panic(fmt.Errorf("❌ Error entering project name: %s", err))
		}
		prompt := promptui.Select{
			Label: "Please select the template source (github/gitee)",
			Items: []string{"github", "gitee"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			panic(fmt.Errorf("❌ Error selecting template source: %s", err))
		}
		switch result {
		case "github":
			cmd.Println("✅ Pulling templates from github...")
			utils.RaiseVoid(exec.Command("git", "clone", githubURL).Run())
		case "gitee":
			cmd.Println("✅ Pulling templates from gitee...")
			utils.RaiseVoid(exec.Command("git", "clone", giteeURL).Run())
		default:
			panic(fmt.Errorf("❌ Invalid template source: %s", result))
		}
		cmd.Println("✅ Renaming project...")
		cwd, err := os.Getwd()
		if err != nil {
			panic(fmt.Errorf("❌ Error getting current working directory: %s", err))
		}
		if sourcePackage != name {
			utils.RaiseVoid(os.Rename(filepath.Join(cwd, sourcePackage), filepath.Join(cwd, name)))
			utils.RaiseVoid(filepath.Walk(filepath.Join(cwd, name), func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.IsDir() {
					return nil
				}
				if info.Mode().IsRegular() {
					content, err := os.ReadFile(path)
					if err != nil {
						return err
					}
					contentStr := string(content)
					if strings.Contains(contentStr, sourcePackage) {
						newContent := strings.ReplaceAll(string(content), sourcePackage, name)
						err = os.WriteFile(path, []byte(newContent), info.Mode().Perm())
						if err != nil {
							return err
						}
					}
				}
				return nil
			}))
		}
		var wg sync.WaitGroup
		errorsChan := make(chan error, 3)

		download := func(name, url string) {
			defer wg.Done()
			err := exec.Command("go", "install", url).Run()
			if err != nil {
				errorsChan <- fmt.Errorf("failed to download %s: %w", name, err)
			}
		}

		wg.Add(3)
		go download("air", airURL)
		go download("easyjson", easyjsonURL)
		go download("swaggo", swaggoURL)
		wg.Wait()
		close(errorsChan)

		for err := range errorsChan {
			if err != nil {
				panic(err)
			}
		}
		utils.RaiseVoid(exec.Command("rm", "-rf", fmt.Sprintf("%s/.git", name)).Run())
		cmd.Println("✅ Project initialized successfully!")
		cmd.Printf("✅ Please run `cd %s && jsonix server` to start the server!\n", name)
	}, utils.DefaultErrorHandler)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "A init for Jsonix",
	Args:  cobra.NoArgs,
	RunE:  initRun,
}
