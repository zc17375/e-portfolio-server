package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Individual struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserUUID    string             `json:"user_uuid" bson:"user_uuid"`
	Name        string             `json:"name" bson:"name" example:"John Doe"`
	JobTitle    string             `json:"job_title" bson:"job_title" example:"Software Engineer"`
	HeadImgPath string             `json:"head_img_path" bson:"head_img_path" example:"/images/john_doe.jpg"`
	Email       string             `json:"email" bson:"email" example:"john@example.com"`
	SocialMedia SocialMedia        `json:"social_media" bson:"social_media"`
	Skills      []string           `json:"skills" bson:"skills" example:"Golang, JavaScript, Python"`
	ResumeLink  string             `json:"resume_link" bson:"resume_link" example:"https://example.com/john_doe_resume.pdf"`
	Projects    []Project          `json:"projects" bson:"projects"`
}

type Project struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        primitive.ObjectID `json:"user_id" bson:"user_id"`
	Name          string             `json:"name" bson:"name" example:"Project X"`
	Skills        []string           `json:"skills" bson:"skills" example:"['Golang', 'Docker', 'Kubernetes']"`
	Introduce     string             `json:"introduce" bson:"introduce" example:"A project for implementing microservices architecture."`
	DemoLink      string             `json:"demo_link" bson:"demo_link" example:"https://example.com/project_x_demo"`
	GithubRepLink string             `json:"github_rep_link" bson:"github_rep_link" example:"https://github.com/johndoe/project_x"`
}

type SocialMedia struct {
	UserID   primitive.ObjectID `json:"user_id" bson:"user_id"`
	Github   string             `json:"github" bson:"github" example:"johndoe"`
	Facebook string             `json:"facebook" bson:"facebook" example:"johndoe"`
	Linkedin string             `json:"linkedin" bson:"linkedin" example:"johndoe"`
	// 可以添加更多的社交媒体字段，比如Twitter等
}

// func (model Individual) CreateIndividual(ctx context.Context, user Individual) (Individual, error) {
// 	insert, err := global.EP_MongoDB.
// 		Collection("individual").
// 		InsertOne(ctx, user)
// 	if err != nil {
// 		return Individual{}, err
// 	}
// 	user.ID = insert.InsertedID.(primitive.ObjectID)
// 	return user, nil
// }
