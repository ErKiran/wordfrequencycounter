package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"wordcount/api/response"
	"wordcount/utils"
)

type WordCountRequest struct {
	Text string `json:"text"`
}

type WordCountResponse struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Data    []utils.WordCount `json:"data"`
}

// WordCountFile godoc
// @Summary Count freuqent words in text file (.txt file)
// @Description Count freuqent words in text file (.txt file)
// @Tags count
// @Accept  json
// @Produce  json
// @Success 200 {object} WordCountResponse
// @Param myFile formData file true "text-file"
// @Router /api/v1/wordcount/file [post]
func (s *Server) WordCountFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	defer file.Close()
	fileName := strings.Split(handler.Filename, ".")

	if len(fileName) < 2 {
		response.ERROR(w, http.StatusBadRequest, fmt.Errorf("invalid file name"))
		return
	}
	if fileName[len(fileName)-1] != "txt" {
		response.ERROR(w, http.StatusBadRequest, fmt.Errorf("invalid file name"))
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	wordCounts, err := utils.TextToWordCount(string(fileBytes))
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if len(wordCounts) < 10 {
		response.JSON(w, http.StatusOK, WordCountResponse{
			Status:  "ok",
			Data:    wordCounts,
			Message: "Text is not long enough to count 10 most frequent words",
		})
		return
	}

	response.JSON(w, http.StatusOK, WordCountResponse{
		Status: "ok",
		Data:   wordCounts[:10],
	})
}

// WordCountText godoc
// @Summary Count freuqent words in text
// @Description Count freuqent words in text
// @Tags count
// @Accept  json
// @Produce  json
// @Success 200 {object} WordCountResponse
// @Param wordcountrequest body WordCountRequest true "wordcountrequest"
// @Router /api/v1/wordcount [post]
func (s *Server) WordCountText(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	wordCountRequest := &WordCountRequest{}
	err = json.Unmarshal(body, wordCountRequest)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	wordCounts, err := utils.TextToWordCount(wordCountRequest.Text)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if len(wordCounts) < 10 {
		response.JSON(w, http.StatusOK, WordCountResponse{
			Status:  "ok",
			Data:    wordCounts,
			Message: "Text is not long enough to count 10 most frequent words",
		})
		return
	}

	response.JSON(w, http.StatusOK, WordCountResponse{
		Status: "ok",
		Data:   wordCounts[:10],
	})
}
