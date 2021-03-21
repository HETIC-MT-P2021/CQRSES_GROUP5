package services

import (
	"context"
	"fmt"
	"github.com/HETIC-MT-P2021/gocqrs/database"
)

//Document represents an Elastic Search(ES) document
type Document struct {
	ID   string
	Body interface{}
}

//NewIndex creates a new index in ES.
func NewIndex(ctx context.Context, index string) error {
	conn := database.EsConn
	exists, err := conn.IndexExists(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("error checking index : %v", err)
	}
	if exists {
		return fmt.Errorf("index already exists")
	}

	indexed, err := conn.CreateIndex(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("cannot create new index: %s", index)
	}

	if !indexed.Acknowledged {
		return fmt.Errorf("cannot create new index: %s", index)
	}

	return nil
}

//NewDocument creates a new document in ES (by Index)
func NewDocument(index string, document *Document) error {
	documentIndexed, err := database.EsConn.Index().
		Index(index).
		BodyJson(document.Body).
		Do(context.Background())

	if err != nil {
		return fmt.Errorf("cannot add resource in index %s", index)
	}

	document.ID = documentIndexed.Id

	return nil
}

//GetDocumentByIndexAndID retrieves a document by its index and its ID from ES
func GetDocumentByIndexAndID(index string, documentID string) (*Document, error) {
	conn := database.EsConn
	ctx := context.Background()
	docRetrieved, err := conn.Get().
		Index(index).
		Id(documentID).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve document id %s, err : %v", documentID, err)
	}
	if !docRetrieved.Found {
		return nil, fmt.Errorf("document not found with id: %s", documentID)
	}
	_, err = conn.Flush().Index(index).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while writing document: %v", err)
	}

	return &Document{
		ID:   docRetrieved.Id,
		Body: docRetrieved.Source,
	}, nil
}
