package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
)

func CheckFace() {

	pwdCmd := exec.Command("ffmpeg.exe", "-i", "rtmp://127.0.0.1/live", "-vframes", "1", "-y", "-f", "image2", "-t", "0.2", "-s", "480x866", "d:/d.jpg")
	if err := pwdCmd.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

func Upload(url, file string) (err error) {
	var respBody []byte
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// Add your image file
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	fw, err := w.CreateFormFile("image_file", file)
	if err != nil {
		return
	}
	if _, err = io.Copy(fw, f); err != nil {
		return
	}
	// Add the other fields
	if err = w.WriteField("api_key", "dpoHjm6j5OILHQiFAfQK83cPxw5MQLU-"); err != nil {
		return
	}
	if err = w.WriteField("api_secret", "BpTIea2MpMzVTwhay3AQjWAIZJZXS4Q1"); err != nil {
		return
	}
	// if fw, err = w.CreateFormField("api_key"); err != nil {
	// 	return
	// }
	// if _, err = fw.Write([]byte("KEY")); err != nil {
	// 	return
	// }
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	respBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Printf("resp status: %s,resp body: %s", res.Status, string(respBody))
	return
}

func main() {
	CheckFace()
	Upload("https://api-cn.faceplusplus.com/facepp/v3/detect", "D:\\d.jpg")
}
