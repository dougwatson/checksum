package main

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	"github.com/dougwatson/go-checksum/checksum"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("needs parameters: go.mod or dir")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 1024)
			n := runtime.Stack(buf, false)
			fmt.Printf("panic: %v\n%s", r, buf[:n])
		}
	}()

	file := os.Args[1]

	println("000000000000000000000000000")
	println("000000000000000000000000000")
	println("000000000000000000000000000")
	fi, err := os.Stat(file)
	if err != nil {
		fmt.Println("ERROR=", err)
		return
	}
	fmt.Printf("fi=%+v\n", fi)

	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("directory: " + file)
		doDir(file)
	case mode.IsRegular():
		fmt.Println("file: " + file)
		doGoMod(file)
	}
}

func doDir(dir string) {
	if len(os.Args) < 3 {
		fmt.Println("needs parameters: module path with version like: github.com/gin-gonic/gin@v1.4.0")
		return
	}

	prefix := os.Args[2]

	h, err := checksum.HashDir(dir, prefix)
	if err != nil {
		println("error=", err.Error())
		return
	}
	fmt.Println(PrettyPrint(h))

}

func doGoMod(file string) {
	h, err := checksum.HashGoMod(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(PrettyPrint(h))

}

// PrettyPrint convert struct to pretty string
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

/*
// -------------------
func unzip(content_zipped []byte, dest string) error {
	//fmt.Println("unzip func dest=", dest)
	println("unzip func dest=", dest)
	r, err := zip.NewReader(bytes.NewReader(content_zipped), int64(len(content_zipped)))
	if err != nil {
		println("error unzipping=", err)
		return err
	}
	for _, f := range r.File {
		rc, err := f.Open() //reading from zip file
		if err != nil {
			println("Zip f.Open error:", err)
		}
		body, err := ioutil.ReadAll(rc)
		if err != nil {
			println("Zip ioutil.ReadAll read error:", err)
		}
		rc.Close()

		if f.Name[:4] == ".git" {
			continue
		}
		println("writing f.Name=", dest+"/"+f.Name)
		println("//////////////////")
		if dest != "" {
			dest = dest + "/"
		}
		if dest == "./" {
			dest = "" //to write the zip in current directory, dest value should be blank
		}
		fp, err := os.Create(dest + f.Name)
		if err != nil {
			println("Zip os.Create error:", err)
		}
		println("^^^^^^^^^^^^^^^")
		n, err := fp.Write(body)
		println("---------------")
		if err != nil {
			println("Zip f.Write error:", err)
		}
		println("n=", n)
		println()
		fp.Close()

		//err = os.WriteFile(dest+"/"+f.Name, body, 0666)
		//if err != nil {
		//		println("Zip os.WriteFile error:", err)
		//	}

		//read with os
		b, err := os.ReadFile(dest + "/" + f.Name)
		if err != nil {
			println("Zip os.ReadFile error:", err)
		}
		println("len(b)=", len(b))

		//Read with easyfs
		//b, err = dfs.ReadFile(dest + "/" + f.Name)
		//if err != nil {
		//	fmt.Println("ERROR in unzip ReadFile=", err)
		//}
		//fmt.Printf("unzip len(b)=%v\n", len(b))

	}
	return nil
}
*/
//println("--------------MKDIR WORKS but not needed----------------")
//err := os.Mkdir("home/go-checksum", 0777)
//if err != nil {
//	fmt.Println("err=", err)
//}
//fmt.Printf("dfs=%+v\n", dfs)
//println("--------------READDIR WORKS----------------")
//rd, err := fs.ReadDir(dfs, "home/go-checksum")
//if err != nil {
//	fmt.Println("err=", err)
//}
//fmt.Printf("rd=%v\n", rd)
//for i, f := range rd {
//	fmt.Printf("%d data=%v\n\n", i, f.Name())
//}

//println("---------- ReadFile by name works--------------------")
//b, err := dfs.ReadFile(file + "/go.mod")
//if err != nil {
//	fmt.Println("errrrrr=", err)
//}
//fmt.Printf("b=%v\n", string(b))
//fmt.Printf("dfs=%+v\n", dfs)
/*
		d, err := srcFile.ReadDir(".")
		if err != nil {
			fmt.Println("errrrrr1=", err)
		}
		fmt.Printf("d=%v\n", d)
		fi, err := d[0].Info()
		if err != nil {
			fmt.Println("errrrrr2=", err)
		}
		fmt.Printf("fi.IsDir()=%v\n", fi.IsDir())

		b1, err := srcFile.ReadFile("src.zip")
		if err != nil {
			fmt.Println("errrrrr3=", err)
		}
		fmt.Printf("len(b1)=%v\n", len(b1))

		println("fi.Size()=", fi.Size())

	if os.WASM {
		r := bytes.NewReader(srcZip)
		zr, err := zip.NewReader(r, int64(len(srcZip)))
		if err != nil {
			println("err=", err)
		}
		for _, f := range zr.File {
			println("f.Name=", f.Name)
		}
		fmt.Printf("dfs1=%+v\n", dfs)

		err = dfs.WriteFile("go-checksum/go.mod2", []byte("mod2 cool"), 0666)
		if err != nil {
			fmt.Println("errrrrr4=", err)
		}
		//b52.Write([]byte("hello"))

		unzip(srcZip, ".") //load initial filesystem

		fmt.Printf("dfs2=%+v\n", dfs)
		fmt.Printf("file=%v\n", file)
		println()

		fmt.Printf("dfs[\"go-checksum/go.mod\"]=%+v\n", string(dfs["go-checksum/go.mod"].Data))

		bf, err := os.Open("go-checksum/go.mod2")
		if err != nil {
			fmt.Println("errrrrr5=", err)
		}
		bb := make([]byte, 100)

		println("read=====================")
		n, err := bf.Read(bb)
		if err != nil {
			println("errrrrr6=", err.Error())
		}
		println("read.....................")
		fmt.Printf("n=%d bb=%v\n", n, string(bb))
	}

		finfo, err := dfs.ReadDir(".")
		if err != nil {
			fmt.Println("ERROR os.ReadDir=", err) // PROBLEMO
			return
		}
		//fmt.Printf("finfo os.ReadDir()=%+v\n", finfo)
		for i, f := range finfo {
			fmt.Printf("%d data=%v\n\n", i, f.Name())
		}
		//	fmt.Printf("finfo os.ReadDir()=%+v\n", finfo.IsDir())
*/

//return

//	defer b.Close()
//	fi, err := b.Stat()
//	if err != nil {
//		fmt.Println("ERROR stat=", err)
//		return
//	}
//	fmt.Printf("fi=%+v\n", fi)
