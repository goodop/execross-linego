package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var HostApi = "https://dev.execross.pw"
var APIKEY = "your_apikey_here"

type ResponseData struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func sendRequest(endpoint string, requestData map[string]interface{}) (*ResponseData, error) {
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response ResponseData
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func sendMultiImageWithUrl(token, to string, urls []string) (*ResponseData, error) {
	endpoint := HostApi + "/api/go/sendMultiImageWithUrl"
	requestData := map[string]interface{}{
		"apikey": APIKEY,
		"token":  token,
		"to":     to,
		"url":    urls,
	}
	return sendRequest(endpoint, requestData)
}

func sendImageWithUrl(token, to, url string) (*ResponseData, error) {
	endpoint := HostApi + "/api/go/sendImageWithUrl"
	requestData := map[string]string{
		"apikey": APIKEY,
		"token":  token,
		"to":     to,
		"url":    url,
	}
	return sendRequest(endpoint, requestData)
}

func sendReplyImageWithUrl(token, ids, to, url string) (*ResponseData, error) {
	endpoint := HostApi + "/api/go/sendReplyImageWithUrl"
	requestData := map[string]string{
		"apikey": APIKEY,
		"token":  token,
		"ids":    ids,
		"to":     to,
		"url":    url,
	}
	return sendRequest(endpoint, requestData)
}

func sendAudioWithUrl(token, to, url string) (*ResponseData, error) {
	endpoint := HostApi + "/api/go/sendReplyImageWithUrl"
	requestData := map[string]string{
		"apikey": APIKEY,
		"token":  token,
		"to":     to,
		"url":    url,
	}
	return sendRequest(endpoint, requestData)
}

func sendReplyImageWithUrl(token, ids, to, url string) (*ResponseData, error) {
	endpoint := HostApi + "/api/go/sendReplyAudioWithUrl"
	requestData := map[string]string{
		"apikey": APIKEY,
		"token":  token,
		"ids":    ids,
		"to":     to,
		"url":    url,
	}
	return sendRequest(endpoint, requestData)
}


func sendVideoWithUrl(token, to, url string) (*ResponseData, error) {
	endpoint := HostApi + "/api/go/sendReplyImageWithUrl"
	requestData := map[string]string{
		"apikey": APIKEY,
		"token":  token,
		"to":     to,
		"url":    url,
	}
	return sendRequest(endpoint, requestData)
}


func sendReplyVideoWithUrl(token, ids, to, url string) (*ResponseData, error) {
	endpoint := HostApi + "/api/go/sendReplyAudioWithUrl"
	requestData := map[string]string{
		"apikey": APIKEY,
		"token":  token,
		"ids":    ids,
		"to":     to,
		"url":    url,
	}
	return sendRequest(endpoint, requestData)
}


func sendFlex(token, to string, meta map[string]interface{}) (*ResponseData, error) {
    endpoint := HostApi + "/api/go/sendFlex"
    flexJSON, err := json.Marshal(meta)
    if err != nil {
        return nil, err
    }
    
    requestData := map[string]interface{}{
        "apikey": APIKEY,
        "token":  token,
        "to":     to,
        "flex":   string(flexJSON),
    }
    return sendRequest(endpoint, requestData)
}



func main() {
	baseAPIURL := "https://dev.execross.pw/api/js/"
	queryParams := url.Values{}
	queryParams.Set("gid", "your_gid_value")
	queryParams.Set("appName", "your_app_name")
	queryParams.Set("authToken", "your_auth_token")
	speedEndpoint := baseAPIURL + "speed?" + queryParams.Encode()
	makeGETRequest(speedEndpoint, "Speed Request")
	invites := []string{"invite1", "invite2", "invite3"} // this mid you want cancelled or like something like that :p
	members := []string{"member1", "member2", "member3"} // this mid you want kick from the group 
	invitesParam := strings.Join(invites, ",")
	membersParam := strings.Join(members, ",")


	// Making a GET request to /api/js/loopkick
	loopKickEndpoint := baseAPIURL + "loopkick?" + queryParams.Encode() + "&invites=" + invitesParam + "&members=" + membersParam
	makeGETRequest(loopKickEndpoint, "Loop Kick Request")


	// Making a GET request to /api/js/multi
	multiEndpoint = baseAPIURL + "multi?" + queryParams.Encode() + "&invites=" + invitesParam + "&members=" + membersParam
	makeGETRequest(multiEndpoint, "Multi Request")


	// Making a GET request to /api/js/cancel
	cancelEndpoint := baseAPIURL + "cancel?" + queryParams.Encode() + "&invites=" + invitesParam
	makeGETRequest(cancelEndpoint, "Cancel Request")

	// sending Media requests
	Token := "Your_auth_token"
	to := "groupid_or_user_id"
	url := "https://cdn.motor1.com/images/mgl/6ZAvXk/s1/lamborghini-invencible.jpg"
	urlAudio := "https://dev.execross.pw/assets/audio-url-test.mp3"
	urlVideo := "https://dev.execross.pw/assets/video-url-test.mp4"
	ids1 := "472433797585961272"
	MultiImg := []string{"https://upload.wikimedia.org/wikipedia/commons/1/19/Lamborghini_Aventador.jpg", "https://www.miaomiaopa.com/wp-content/uploads/2023/03/3f10ba5ca81d82cacef50fdcd6550b23.jpg", "https://www.miaomiaopa.com/wp-content/uploads/2023/03/374d1ded15ad5305d13a84118cbab938.jpg", "https://img-9gag-fun.9cache.com/photo/arG92Ld_460s.jpg"}

	// Send Image
	responseImg1, errImg1 := sendImageWithUrl(Token, to, url)
	if errImg1 != nil {
		fmt.Println("Error sending image with URL:", errImg1)
		return
	}
	fmt.Println(responseImg1)


	// Send Multi Image
	resImg, errMultiImg := sendMultiImageWithUrl(Token, to, MultiImg)
	if errMultiImg != nil {
		fmt.Println("Error sending image with URL:", errMultiImg)
		return
	}
	fmt.Println(resImg)

	// Send Reply Image
	responseImg2, errImg2 := sendReplyImageWithUrl(Token, ids1, to, url)
	if errImg2 != nil {
		fmt.Println("Error sending reply image with URL:", errImg2)
		return
	}
	fmt.Println(responseImg2)

	// Send Audio
	responseAud1, errAud1 := sendAudioWithUrl(Token, to, urlAudio)
	if errAud1 != nil {
		fmt.Println("Error sending audio with URL:", errAud1)
		return
	}
	fmt.Println(responseAud1)

	// Send Reply Audio
	responseAud2, errAud2 := sendReplyAudioWithUrl(Token, ids1, to, urlAudio)
	if errAud2 != nil {
		fmt.Println("Error sending reply audio with URL:", errAud2)
		return
	}
	fmt.Println(responseAud2)

	// Send Video
	responseVid1, errVid1 := sendVideoWithUrl(Token, to, urlVideo)
	if errVid1 != nil {
		fmt.Println("Error sending video with URL:", errVid1)
		return
	}
	fmt.Println(responseVid1)

	// Send Reply Video
	responseVid2, errVid2 := sendReplyVideoWithUrl(Token, ids1, to, urlVideo)
	if errVid2 != nil {
		fmt.Println("Error sending reply video with URL:", errVid2)
		return
	}
	fmt.Println(responseVid2)


	// Send Flex Example (Image)
	background := "https://cdn.motor1.com/images/mgl/6ZAvXk/s1/lamborghini-invencible.jpg"
	meta := map[string]interface{}{
		"type":               "image",
		"originalContentUrl": url,
		"previewImageUrl":    url,
	}
	dataFlex := map[string]interface{}{
		"type":     "flex",
		"altText":  "Developer Execross",
		"contents": meta,
	}
	responseFlex, errFlex := sendFlex(Token, to, dataFlex)
	if errFlex != nil {
		fmt.Println("Error sending Flex (Image):", errFlex)
		return
	}
	fmt.Println(responseFlex)


	// Send Flex Example (Video)
	meta = map[string]interface{}{
		"type":               "video",
		"originalContentUrl": urlVideo,
		"previewImageUrl":    url,
	}
	dataFlex = map[string]interface{}{
		"type":     "flex",
		"altText":  "Developer Execross",
		"contents": meta,
	}
	responseFlex, errFlex = sendFlex(Token, to, dataFlex)
	if errFlex != nil {
		fmt.Println("Error sending Flex (Video):", errFlex)
		return
	}
	fmt.Println(responseFlex)

	// Send Flex Example (Audio)
	meta = map[string]interface{}{
		"type":               "audio",
		"originalContentUrl": urlAudio,
		"duration":           60000,
	}
	dataFlex = map[string]interface{}{
		"type":     "flex",
		"altText":  "Developer Execross",
		"contents": meta,
	}
	responseFlex, errFlex = sendFlex(Token, to, dataFlex)
	if errFlex != nil {
		fmt.Println("Error sending Flex (Audio):", errFlex)
		return
	}
	fmt.Println(responseFlex)
}

func makeGETRequest(apiURL, requestName string) {
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("%s: Error - %v\n", requestName, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("%s: Request Successful\n", requestName)
	} else {
		fmt.Printf("%s: Request failed with status code: %d\n", requestName, resp.StatusCode)
	}
}
