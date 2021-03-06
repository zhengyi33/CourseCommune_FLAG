package main

import (
	"fmt"
	"reflect"

	"github.com/olivere/elastic"
	"github.com/pborman/uuid"
)

const (
	QUESTION_INDEX = "question"
)

type Question struct {
	Id          string `json:"id"`
	User        string `json:"user"`
	Description string `json:"description"`
	Answer      string `json:"answer"`
}

func searchQuestionsByKeywords(keywords string) ([]Question, error) {
	query := elastic.NewMatchQuery("description", keywords)
	query.Operator("AND")
	if keywords == "" {
		query.ZeroTermsQuery("all")
	}
	searchResult, err := readFromES(query, QUESTION_INDEX)
	if err != nil {
		return nil, err
	}
	return getQuestionFromSearchResult(searchResult), nil
}

func getQuestionFromSearchResult(searchResult *elastic.SearchResult) []Question {
	var ptype Question
	var questions []Question

	for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
		if p, ok := item.(Question); ok {
			questions = append(questions, p)
		}
	}
	return questions
}

func saveQuestion(question *Question) error {
	//random generate a unit string as file name
	id := uuid.New()
	//1. save to GCS


	//2.save to elastic search
	err = saveToES(question, QUESTION_INDEX, id)
	if err != nil {
		return err
	}
	fmt.Printf("question is saved to elastic earch: %s.\n", question.Message)
	return nil
}
