package main

import (
	"fmt"
	"io/ioutil"
	"os"
)


func main() {
	fmt.Printf("Make sure you subscribed to the workshop item in steam!")
    inputPath := `C:\Program Files (x86)\Steam\steamapps\workshop\content\555160\2562926750\pakchunk_workshop_stage.scoreboard.pak`
    outputPath := `C:\Program Files (x86)\Steam\steamapps\common\PavlovVR\Pavlov\Content\Paks\pakchunk_workshop_stage.scoreboard.pak`
    
    // Part 1: open input file.
    inputFile, _ := os.Open(inputPath)
    
    // Part 2: call ReadAll to get contents of input file.
    data, _ := ioutil.ReadAll(inputFile)
    
    // Part 3: write data to copy file.
    ioutil.WriteFile(outputPath, data, 0)
}