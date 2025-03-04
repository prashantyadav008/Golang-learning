package fileHandling

import (
	"fmt"
	"io"
	"os"
)

func FileHandling() {
	//-------------------------------------- Create the File ----------------------------------------

	// Create a File
	file, err := os.Create("19.fileHandling/example.txt")
	if err != nil {
		fmt.Println("Error while Creating File: ", err)
		return
	}

	defer file.Close() // free the resources and occurs resource leak
	fmt.Println("File Created Successfully")

	//-------------------------------------- Write into the File ----------------------------------------

	// Write some Content into the Write Content into File
	content := "You Successfully Create a File using Golang!"
	byte1, err1 := io.WriteString(file, content)
	if err1 != nil {
		fmt.Println("Error while Write into the File: ", err)
		return
	}

	fmt.Println("File Content Updated Successfully and Bytes is:", byte1)

	//-------------------------------------- Read into the File ----------------------------------------
	// Read the Content of the Reading File ***************without buffer***************
	fileContent, err := os.ReadFile("19.fileHandling/example.txt")
	if err != nil {
		fmt.Println("Error while Reading the File using without buffer: ", err)
		return
	}

	fmt.Println("Read File Content without Buffer: ", string(fileContent))

	// Read the Content of the Reading File ***************with buffer***************
	fileContent1, err := os.Open("19.fileHandling/example1_buffer.txt")
	if err != nil {
		fmt.Println("Error while Reading the File using Buffer: ", err)
		return
	}
	defer fileContent1.Close() // free the resources and occurs resource leak

	// create a buffer and it is temporary storage
	buffer := make([]byte, 1024)
	// read the content of the file
	for {
		n, err := fileContent1.Read(buffer)
		if err == io.EOF { // EOF end of file
			break
		}
		if err != nil {
			fmt.Println("Error while Reading the File using Buffer: ", err)
			return
		}

		// process the read content
		fmt.Println("Read File Content with Buffer: ", string(buffer[:n]))
	}

}
