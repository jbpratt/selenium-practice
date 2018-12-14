package util

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/tebeka/selenium"
)

func Log() {
	fmt.Println("selenium util pkg")
}

func SetUp() {
	cmd := exec.Command("docker", "run", "-d", "--name=goselenium", "-p=4444:4444", "-v=/dev/shm:/dev/shm", "selenium/standalone-firefox")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Selenium container has been created.")
	time.Sleep(4000 * time.Millisecond)
}

func TearDown() {
	cmd := exec.Command("docker", "rm", "goselenium", "-f")
	cmd.CombinedOutput()
	fmt.Println("Container was removed")
}

func NewDriver() selenium.WebDriver {
	// Connect to the selenium server
	caps := selenium.Capabilities{"browserName": "firefox"}
	d, err := selenium.NewRemote(caps, "http://127.0.0.1:4444/wd/hub")
	defer d.Quit()
	if err != nil {
		fmt.Println(err)
	}
	return d
}
