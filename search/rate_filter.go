package search

import (
	"time"

	"github.com/core-go/search"
)

type RateFilter struct {
	*search.Filter
	Id          string            `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"id" dynamodbav:"id" firestore:"id" match:"equal" validate:"max=255"`
	Author      string            `mapstructure:"author" json:"author,omitempty" gorm:"column:author;primary_key" bson:"author" dynamodbav:"author" firestore:"author" match:"equal" validate:"max=255"`
	Rate        string            `mapstructure:"rate" json:"rate,omitempty" gorm:"column:rate" bson:"rate" dynamodbav:"rate" firestore:"rate" match:"equal" validate:"max=10"`
	Review      string            `mapstructure:"review" json:"review" gorm:"column:review" bson:"review" dynamodbav:"review" firestore:"review"`
	Time        *search.TimeRange `mapstructure:"time" json:"time" gorm:"column:time" bson:"time" dynamodbav:"time" firestore:"time"`
	UsefulCount string            `mapstructure:"usefulCount" json:"usefulCount,omitempty" gorm:"column:usefulCount" bson:"usefulCount" dynamodbav:"usefulCount" firestore:"usefulCount"`
	ReplyCount  string            `mapstructure:"replyCount" json:"replyCount,omitempty" gorm:"column:replyCount" bson:"replyCount" dynamodbav:"replyCount" firestore:"replyCount"`
}

type RateCriteriaFilter struct {
	Id          string       `mapstructure:"id" json:"id,omitempty" gorm:"column:id;primary_key" bson:"id,omitempty" dynamodbav:"id,omitempty" firestore:"id,omitempty" validate:"required,max=255"`
	Author      string       `mapstructure:"author" json:"author,omitempty" gorm:"column:author;primary_key" bson:"author,omitempty" dynamodbav:"author,omitempty" firestore:"author,omitempty" validate:"required,max=255"`
	Rate        float32      `mapstructure:"rate" json:"rate,omitempty" gorm:"column:rate" bson:"rate,omitempty" dynamodbav:"rate,omitempty" firestore:"rate,omitempty" validate:"required,max=10"`
	Rates       []float32    `mapstructure:"rates" json:"rates,omitempty" gorm:"column:rates" bson:"rates,omitempty" dynamodbav:"rates,omitempty" firestore:"rates,omitempty"`
	Review      string       `mapstructure:"review" json:"review,omitempty" gorm:"column:review" bson:"review,omitempty" dynamodbav:"review,omitempty" firestore:"review,omitempty"`
	Time        *time.Time   `mapstructure:"time" json:"time,omitempty" gorm:"column:time" bson:"time,omitempty" dynamodbav:"time,omitempty" firestore:"time,omitempty"`
	UsefulCount int          `mapstructure:"usefulCount" json:"usefulCount" gorm:"column:usefulCount" bson:"usefulCount,omitempty" dynamodbav:"usefulCount,omitempty" firestore:"usefulCount,omitempty"`
	ReplyCount  int          `mapstructure:"replyCount" json:"replyCount" gorm:"column:replyCount" bson:"replyCount,omitempty" dynamodbav:"replyCount,omitempty" firestore:"replyCount,omitempty"`
	Histories   *[]Histories `mapstructure:"histories" json:"histories" gorm:"column:histories" bson:"histories,omitempty" dynamodbav:"histories,omitempty" firestore:"histories,omitempty"`
}
