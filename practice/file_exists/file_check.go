package file_check

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CheckFile() {
	path := "C:\\Users\\ASWIN\\Desktop\\scripts\\python-scripts\\ext-finder\\input_coverage\\TC_verifyNfcdeInitPerformace\\nfc\\out\\soong\\.intermediates\\hardware\\nxp\\nfc\\snxxx\\android.hardware.nfc_snxxx@1.2-service\\android_vendor.33_arm64_armv8-a_cortex-a73\\obj\\hardware\\nxp\\nfc\\snxxx\\halimpl\\eseclients_extns\\src"
	if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("Directory does not exist.")
	} else {
			fmt.Println("Directory exists.")
	}

	filePath := path+"\\eSEClientExtns.gcda"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}