package api

import (
	"github.com/daniil49926/FastTubeAnalytics/internal/app/storage"
	"io"
	"log"
	"net/http"
	"strings"
)

func (s *Server) handleHealthChecker() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		jsonData, err := makeOkResult()
		if err != nil {
			log.Fatal(err)
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusOK)
		if _, err = writer.Write(jsonData); err != nil {
			return
		}
	}
}

func (s *Server) handleSendAnalytics() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		body, err := io.ReadAll(request.Body)
		if err != nil {
			jsonData, err := makeBadDataResult()
			if err != nil {
				log.Fatal(err)
			}
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write(jsonData)
			return
		}

		dataList := strings.Split(string(body), ";")

		if len(dataList) != 2 {
			jsonData, err := makeFailSplitResult()
			if err != nil {
				log.Fatal(err)
			}
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write(jsonData)
			return
		}

		requestOnAnalise := strings.Split(dataList[0], "-")[1:]
		responseOnAnalise := strings.Split(dataList[1], "-")[1:]

		if len(requestOnAnalise) != 3 || len(responseOnAnalise) != 1 {
			jsonData, err := makeFailSplitResult()
			if err != nil {
				log.Fatal(err)
			}
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write(jsonData)
			return
		}

		err = storage.InsertStatements(requestOnAnalise, responseOnAnalise)
		if err != nil {
			jsonData, err := makeFailInsertResult()
			if err != nil {
				log.Fatal(err)
			}
			writer.WriteHeader(http.StatusBadRequest)
			_, _ = writer.Write(jsonData)
			return
		}
		jsonData, err := makeOkResult()
		if err != nil {
			log.Fatal(err)
		}

		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(jsonData)
		return

	}

}
