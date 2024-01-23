package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"time"
)


type IndicatorValueRequest struct {
	AccessKey int64  `json:"access_key"`
	IndicatorValue int `json:"value"`
}

type Request struct {
	IndicatorId int64 `json:"indicator_id"`
	EstimateId int64 `json:"estimate_id"`
}


func (h *Handler) issueIndicatorValue(c *gin.Context) {
	var input Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issueIndicatorValue:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(4 * time.Second)
		sendIndicatorValueRequest(input)
	}()
}

func sendIndicatorValueRequest(request Request) {

	var value = -1
	if rand.Intn(10) % 10 >= 2 {
	 value = rand.Intn(10000)
	}

	answer := IndicatorValueRequest{
		AccessKey: 123,
		IndicatorValue: value,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/estimates/%d/update_indicator/%d/", request.EstimateId, request.IndicatorId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}
