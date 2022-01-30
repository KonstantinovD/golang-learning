package database

import (
	"context"
	"errors"
	"leaning/GO_WEBSITE_01/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SnippetModel struct {
	SnippetsCollection *mongo.Collection
}

var client *mongo.Client
var ctx context.Context

func InitDB(url string) (*SnippetModel, error) {
	ctx = context.TODO()
	clientOptions := options.Client().ApplyURI(url)
	cl, err := mongo.Connect(ctx, clientOptions)
	client = cl

	if err != nil {
		return nil, err
	}
	collection := client.Database("golang-learning").
		Collection("snippets")
	model := SnippetModel{SnippetsCollection: collection}
	return &model, nil
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	filter := bson.M{"snippetId": id}
	cur, err := m.SnippetsCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var snippet *models.Snippet

	if !cur.Next(ctx) {
		return nil, mongo.ErrNoDocuments
	}

	err = cur.Decode(&snippet)
	if err != nil {
		return nil, errors.New("cannot decode snippet")
	}
	if cur.Next(ctx) {
		return nil, errors.New("found multiple snippets")
	}
	return snippet, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	opts := options.Find()
	opts.SetLimit(10)

	filter := bson.D{
		{"expires", bson.D{
			{"$gt", time.Now()},
		}},
	}

	cur, err := m.SnippetsCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	var snippets []*models.Snippet

	for cur.Next(ctx) {
		var sn models.Snippet
		err := cur.Decode(&sn)
		if err != nil {
			return nil, errors.New("cannot decode snippet")
		}
		snippets = append(snippets, &sn)
	}
	return snippets, nil
}
