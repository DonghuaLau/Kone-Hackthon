package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"testing"
	"time"
)

func TestJson2Map(t *testing.T) (mapResult map[string]interface{}) {
	jsonStr := `
    {
        "name":"liangyongxing",
        "age":12
    }
    `
	//var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		t.Fatal(err)
	}
	t.Log(mapResult)
	return mapResult
}

func Judge(confidence float64) (res int) {
	if confidence > 90 {
		res = 1
	}
	return res
}

func CheckFace(imgFile string) {

	pwdCmd := exec.Command("ffmpeg.exe", "-i", "rtmp://127.0.0.1/live", "-vframes", "1", "-y", "-f", "image2", "-t", "0.1", "-s", "480x866", imgFile)
	if err := pwdCmd.Run(); err != nil {
		fmt.Println("err: ", err)
	}

}

func Upload(url, file string, face_set string) (resp string, err error) {
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
	if err = w.WriteField("outer_id", face_set); err != nil {
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
		return "", err
	}
	resp = string(respBody)
	fmt.Printf("resp status: %s,resp body: %s", res.Status, resp)
	return resp, err
}

func get_value(result string) (float,string){
	
	var confidence float
	var face_token string

	var f interface{}
	err := json.Unmarshal(result, &f)
	m := f.(map[string]interface{})

	for k, v := range m {
    switch vv := v.(type) {
    case []interface{}:
        for i, u := range vv {
            if i=="confidence" {
            	confidence=u
            }else if i== "face_token" {
            	face_token=u
            }
        }
    }

    return confidence, face_token;
}

func main() {
	//var res string
	timestamp := int(time.Now().Unix())
	imgFile := "d:\\temp\\" + strconv.Itoa(timestamp) + ".jpg"
	fmt.Printf("%s", imgFile)
	CheckFace(imgFile)

	confidence1 := 0.1 // 员工可信度
	confidence2 := 0.1 // 访客可信度
	
	rc_1 := regexp.MustCompile(`[a-z]+`)  //
	rc_2 := regexp.MustCompile(`[a-z]+`)  // 
	
	// 检查是否为员工
	res1, err := Upload("https://api-cn.faceplusplus.com/facepp/v3/search", imgFile, "staff")
	if err != nil {
		fmt.Println("err: ", err)
	}
	

	if r := Judge(confidence1); r == 1 {
		fmt.Println("staff")
	}
	// 检查是否为访客
	res2, err := Upload("https://api-cn.faceplusplus.com/facepp/v3/search", imgFile, "staff")
	if err != nil {
		fmt.Println("err: ", err)
	}
	if r := Judge(confidence2); r == 1 {
		fmt.Println("guests")
	}

}
