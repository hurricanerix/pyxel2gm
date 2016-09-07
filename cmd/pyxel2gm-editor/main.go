package main

func main() {
	// os.Args[1] should be the full path to the image file to be edited.
	// Step1: parse filepath for project_dir and filename.
	// Step2: look in project_dir for .pyxel2gm.conf (contains assets_dir)
	// Step3: walk assets_dir looking for a pyxel file corresponding to the filename.
	// Step4: open pyxel edit passing in the pyxel file as a parameter
	// Step5: monitor pyxel file for changes, when file modified time is updated,
	//        call export just the modified file.
	/*
		f, err := os.Create("test.log")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println(os.Args)
	*/
}
