package rank

import (
	"time"

	mdb "github.com/nikkerella/hitotose/gin/db/mongo"
	"github.com/nikkerella/hitotose/gin/model/rank"
	mgo "github.com/nikkerella/monggo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	// Rank
	ByID(rID any) rank.Rank
	ByTitle(title string) []rank.Rank
	Create(title string) error
	Delete(rID any) error
	Query() []rank.Rank
	Update(t rank.Rank) error

	// Members
	Add(rID any, mName string) error
	Rank(rID any, mOrder int, mName string)
	Remove(rID any, mName string) error
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) Add(rID any, mName string) error {
	r := s.ByID(rID)
	newMember := rank.Member{
		Order: 0,
		Name:  mName,
	}
	r.Members = append(r.Members, newMember)
	return s.Update(r)
}

func (s *service) ByID(rID any) rank.Rank {
	var rank rank.Rank
	mgo.FindID(mdb.Ranks, rID).Decode(&rank)
	return rank
}

func (s *service) ByTitle(title string) []rank.Rank {
	var ranks []rank.Rank
	filter := bson.D{primitive.E{Key: "title", Value: primitive.Regex{Pattern: title, Options: "i"}}}
	sort := bson.D{primitive.E{Key: "title", Value: 1}}
	mgo.FindMany(mdb.Ranks, &ranks, filter, sort)
	return ranks
}

func (s *service) Create(title string) error {
	t := rank.Rank{
		ID:        primitive.NewObjectIDFromTimestamp(time.Now()),
		Title:     title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return mgo.Insert(mdb.Ranks, t)
}

func (s *service) Delete(rID any) error {
	return mgo.DeleteID(mdb.Ranks, rID)
}

func (s *service) Query() []rank.Rank {
	return s.ByTitle("")
}

func (s *service) Rank(rID any, mOrder int, mName string) {
	panic("unimplemented")
}

func (s *service) Remove(rID any, mName string) error {
	rank := s.ByID(rID)
	rank.Remove(mName)
	return s.Update(rank)
}

func (s *service) Update(r rank.Rank) error {
	update := bson.D{primitive.E{
		Key: "$set", Value: bson.D{
			primitive.E{Key: "title", Value: r.Title},
			primitive.E{Key: "members", Value: r.Members},
			primitive.E{Key: "updated_at", Value: time.Now()},
		}},
	}
	return mgo.Update(mdb.Ranks, r.ID, update)
}
