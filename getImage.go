package ASCII

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetImage(imageUrl string) {
	// URL of the image you want to download
	//imageUrl := "https://images.pexels.com/photos/17929271/pexels-photo-17929271/free-photo-of-woman-standing-on-vineyard.jpeg"
	//imageUrl := "https://wallpapercave.com/wp/wp2941777.png"

	// Create an HTTP GET request
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return
	}

	// Create a new file to save the image
	outputFile, err := os.Create("downloaded_image.png")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return
	}

	fmt.Println("Image successfully downloaded and saved as 'downloaded_image.png'")
}
