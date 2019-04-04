package tempfiles

import (
	"fmt"
	//Creating temporary files and directories is done using this package:

	"io/ioutil"
	"os"
)


func WorkWithTemp() error {

	//t is an os.File object
	t, err := ioutil.TempDir("", "tmp")
	if err != nil {
		return err
	}

	//you must delete the files on your own, as shown here:
	defer os.RemoveAll(t)

	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}

	fmt.Println(tf.Name())
	return nil
}

