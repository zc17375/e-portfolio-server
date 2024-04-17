package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Individual struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserUUID    string             `json:"-" bson:"user_uuid,omitempty"`
	UserName    string             `json:"username" bson:"username,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty" example:"John Doe"`
	JobTitle    string             `json:"job_title" bson:"job_title,omitempty" example:"Software Engineer"`
	HeadImgPath string             `json:"head_img_path" bson:"head_img_path,omitempty" example:"/images/john_doe.jpg"`
	Email       string             `json:"email" bson:"email,omitempty" example:"john@example.com"`
	SocialMedia SocialMedia        `json:"social_media" bson:"social_media"`
	Skills      []string           `json:"skills" bson:"skills,omitempty" example:"Golang,JavaScript,Python"`
	ResumeLink  string             `json:"resume_link" bson:"resume_link,omitempty" example:"https://example.com/john_doe_resume.pdf"`
	Projects    []Project          `json:"projects" bson:"projects"`
	CreatedAt   time.Time          `json:"-" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"-" bson:"updated_at"`
}

type Project struct {
	ID            string             `json:"-" bson:"id"`
	UserID        primitive.ObjectID `json:"-" bson:"user_id,omitempty"`
	Name          string             `json:"name" bson:"name,omitempty" example:"Project X"`
	Skills        []string           `json:"skills" bson:"skills,omitempty" example:"['Golang', 'Docker', 'Kubernetes']"`
	Introduce     string             `json:"introduce" bson:"introduce,omitempty" example:"A project for implementing microservices architecture."`
	DemoLink      string             `json:"demo_link" bson:"demo_link,omitempty" example:"https://example.com/project_x_demo"`
	GithubRepLink string             `json:"github_rep_link" bson:"github_rep_link,omitempty" example:"https://github.com/johndoe/project_x"`
}

type SocialMedia struct {
	UserID   primitive.ObjectID `json:"-" bson:"user_id,omitempty"`
	Github   string             `json:"github" bson:"github,omitempty" example:"johndoe"`
	Facebook string             `json:"facebook" bson:"facebook,omitempty" example:"johndoe"`
	Linkedin string             `json:"linkedin" bson:"linkedin,omitempty" example:"johndoe"`
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
