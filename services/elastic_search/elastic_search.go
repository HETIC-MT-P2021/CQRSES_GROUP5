package elastic_search

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type EsConnector struct {
	client *elastic.Client
}

type EsService interface {
	NewIndex(context.Context, string) error
	NewDocument(context.Context, string, *Document) error
	GetDocumentByIndexAndID(context.Context, string, string) (*Document, error)
}

func NewEsConnector(client *elastic.Client) *EsConnector {
	return &EsConnector{client: client}
}

//Document represents an Elastic Search(ES) document
type Document struct {
	ID   string
	Body interface{}
}

//NewIndex creates a new index in ES.
func (connector *EsConnector) NewIndex(ctx context.Context, index string) error {
	exists, err := connector.client.IndexExists(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("error checking index : %v", err)
	}
	if exists {
		return fmt.Errorf("index already exists")
	}

	indexed, err := connector.client.CreateIndex(index).Do(ctx)
	if err != nil {
		return fmt.Errorf("cannot create new index: %s", index)
	}

	if !indexed.Acknowledged {
		return fmt.Errorf("cannot create new index: %s", index)
	}

	return nil
}


//NewDocument creates a new document in ES (by Index)
func (connector *EsConnector) NewDocument(ctx context.Context, index string, document *Document) error {
	exists, err := connector.client.IndexExists(index).Do(ctx)
	if !exists {
		if err := connector.NewIndex(ctx, index); err != nil {
			return fmt.Errorf("could not create index: %s", index)
		}
	}
	documentIndexed, err := connector.client.Index().
		Index(index).
		BodyJson(document.Body).
		Do(ctx)

	if err != nil {
		return fmt.Errorf("cannot add resource in index %s : %v", index, err)
	}

	document.ID = documentIndexed.Id

	return nil
}

func (connector *EsConnector) UpdateDocument(ctx context.Context, index string, document *Document) (*elastic.UpdateResponse, error){
	
	return connector.client.Update().Index(index).Id(document.ID).Doc(document).Do(ctx)
}

//GetDocumentByIndexAndID retrieves a document by its index and its ID from ES
func (connector *EsConnector) GetDocumentByIndexAndID(ctx context.Context, index string, documentID string) (*Document, error) {
	docRetrieved, err := connector.client.Get().
		Index(index).
		Id(documentID).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve document id %s, err : %v", documentID, err)
	}
	if !docRetrieved.Found {
		return nil, fmt.Errorf("document not found with id: %s", documentID)
	}
	_, err = connector.client.Flush().Index(index).Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("error while writing document: %v", err)
	}

	return &Document{
		ID:   docRetrieved.Id,
		Body: docRetrieved.Source,
	}, nil
}
