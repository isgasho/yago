package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GenDir(srcPath string, destPath, app string) error {
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			return errors.New("srcPath 不是一个正确的目录！")
		}
	}
	if destInfo, err := os.Stat(destPath); err != nil {
		return err
	} else {
		if !destInfo.IsDir() {
			return errors.New("destPath 不是一个正确的目录！")
		}
	}

	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			if err := GenFile(path, destNewPath, app); err != nil {
				log.Println(fmt.Sprintf("create file %s error:", destNewPath), err.Error())
				return err
			}
		}
		return nil
	})
	return err
}

func GenFile(src, dest, app string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destSplitPathDirs := strings.Split(dest, "/")

	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := pathExists(destSplitPath)
			if b == false {
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					return err
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	srcFileInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	content := make([]byte, srcFileInfo.Size())
	if _, err := srcFile.Read(content); err != nil {
		return err
	}

	contentStr := strings.ReplaceAll(string(content), "github.com/hulklab/yago/example/app", app)

	if _, err := dstFile.WriteString(contentStr); err != nil {
		return err
	}
	return nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init app",
	Long:  `Init a app named by input`,
	Run: func(cmd *cobra.Command, args []string) {
		useMod, _ := cmd.Flags().GetBool("mod")
		app, _ := cmd.Flags().GetString("app")

		log.Println("create app", app)
		if err := os.MkdirAll(app, 0755); err != nil {
			log.Println("create app dir error:", err.Error())
		}
		var src string
		fmt.Println(useMod)
		if useMod {
			src = fmt.Sprintf("%s/pkg/mod/github.com/hulklab/yago@%s/example/app", os.Getenv("GOPATH"), Version)
		} else {
			src = fmt.Sprintf("%s/src/github.com/hulklab/yago/example/app", os.Getenv("GOPATH"))
		}
		dest := app

		if err := GenDir(src, dest, app); err != nil {
			log.Println("create app error:", err.Error())
		}
	},
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "new module",
	Long:  `new a module named by input`,
	Run: func(cmd *cobra.Command, args []string) {
		app, _ := cmd.Flags().GetString("app")
		module, _ := cmd.Flags().GetString("module")

		log.Println("create module", module)
		dirs := []string{"cmd", "dao", "http", "model", "rpc", "task"}
		for _, d := range dirs {
			dirPath := fmt.Sprintf("modules/%s/%s", module, module+d)
			if err := os.MkdirAll(dirPath, 0755); err != nil {
				log.Println(fmt.Sprintf("create module dir %s error:", dirPath), err.Error())
				return
			}
			filePath := fmt.Sprintf("%s/%s.go", dirPath, module)
			fileBody := fmt.Sprintf("package %s%s", module, d)
			if err := ioutil.WriteFile(filePath, []byte(fileBody), 0644); err != nil {
				log.Println(fmt.Sprintf("create module file %s error:", filePath), err.Error())
				return
			}
		}

		routes := []string{"cmd", "http", "rpc", "task"}
		for _, d := range routes {
			routePath := fmt.Sprintf("routes/%sroute/%s.go", d, d)
			var routeBody []byte
			var err error
			if routeBody, err = ioutil.ReadFile(routePath); err != nil {
				log.Println(fmt.Sprintf("read route file %s error:", routePath), err.Error())
				return
			}
			newRoute := fmt.Sprintf("\t_ \"%s/modules/%s/%s%s\"\n)", app, module, module, d)
			contentStr := strings.ReplaceAll(string(routeBody), ")", newRoute)
			if err = ioutil.WriteFile(routePath, []byte(contentStr), 0644); err != nil {
				log.Println(fmt.Sprintf("write route file %s error:", routePath), err.Error())
				return
			}
			cmd := exec.Command("gofmt", "-w", routePath)
			if err := cmd.Run(); err != nil {
				log.Println(fmt.Sprintf("gofmt route file %s error:", routePath), err.Error())
				return
			}
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Print version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yago version", Version)
	},
}

var rootCmd = &cobra.Command{}

func main() {

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(newCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Println("cmd run error:", err.Error())
		os.Exit(1)
	}
}

func init() {
	// init cmd
	initCmd.Flags().BoolP("mod", "", true, "是否使用go mod")
	// init cmd
	initCmd.Flags().StringP("app", "a", "", "应用名称")
	_ = initCmd.MarkFlagRequired("app")

	// module cmd
	newCmd.Flags().StringP("app", "a", "", "应用名称")
	_ = newCmd.MarkFlagRequired("app")
	newCmd.Flags().StringP("module", "m", "", "模块名称")
	_ = newCmd.MarkFlagRequired("module")
}