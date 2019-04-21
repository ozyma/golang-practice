package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	//grab a directory name for today's project
	dir := flag.String("dir", "", "Input a name of today's project")
	flag.Parse()

	if *dir == "" {
		fmt.Println("Please provide a directory name")
		os.Exit(1)
	}

	//get today's date
	t := time.Now()

	//get today's date and fuse it with the project name for our directory path
	parentPath := fmt.Sprintf("%s-%s", t.Format("2006-01-02"), *dir)
	fmt.Println("Parent Path is the following: ", parentPath)
	os.MkdirAll(parentPath, 0777)

	srcPath := fmt.Sprintf("%s/%s", parentPath, "src")
	fmt.Println("Created src Path: ", srcPath)
	os.MkdirAll(srcPath, 0777)

	assetPath := fmt.Sprintf("%s/%s", parentPath, "assets")
	fmt.Println("Created assets Path: ", assetPath)
	os.MkdirAll(assetPath, 0777)

	buildPath := fmt.Sprintf("%s/%s", parentPath, "build")
	fmt.Println("Created a build path at: ", buildPath)
	os.MkdirAll(buildPath, 0777)

	READMEPath := fmt.Sprintf("%s/%s", parentPath, "README.md")
	fmt.Println("Created a README file at: ", READMEPath)
	os.Create(READMEPath)

	
}
