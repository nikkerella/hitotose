package rank

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rank struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Members   []Member           `json:"members" bson:"members"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Member struct {
	Order int    `json:"order" bson:"order"`
	Name  string `json:"name" bson:"name"`
}

func (r *Rank) Clear() {
	r.Members = []Member{}
}

func (r *Rank) Remove(mName string) {
	for i, member := range r.Members {
		if member.Name == mName {
			// Remove the member by slicing
			r.Members = append(r.Members[:i], r.Members[i+1:]...)
			return
		}
	}
}

func (r *Rank) Update(mOrder int, mName string) {
	for i, member := range r.Members {
		if member.Name == mName {
			// Update the member's order
			r.Members[i].Order = mOrder
			break
		}
	}
}

func (r *Rank) UnrankedMembers() []Member {
	var unrankedMembers []Member
	for _, member := range r.Members {
		if member.Order == 0 {
			unrankedMembers = append(unrankedMembers, member)
		}
	}
	return unrankedMembers
}

func (r *Rank) RankedMembers() []Member {
	var rankedMembers []Member
	for _, member := range r.Members {
		if member.Order > 0 {
			rankedMembers = append(rankedMembers, member)
		}
	}
	return rankedMembers
}
