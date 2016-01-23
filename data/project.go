package data

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	ID         bson.ObjectId   `bson:"_id"`
	Name       string          `bson:"name"`
	OwnerID    bson.ObjectId   `bson:"owner_id"`
	MemberIDs  []bson.ObjectId `bson:"member_ids"`
	CreatedAt  time.Time       `bson:"created_at"`
	ModifiedAt time.Time       `bson:"modified_at"`
}

func GetProject(id bson.ObjectId) (*Project, error) {
	pro := Project{}
	err := sess.DB("").C(projectC).FindId(id).One(&pro)
	if err == mgo.ErrNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pro, nil
}

func (p *Project) Put() error {
	p.ModifiedAt = time.Now()

	if p.ID == "" {
		p.ID = bson.NewObjectId()
		p.CreatedAt = p.ModifiedAt
	}
	_, err := sess.DB("").C(organizationC).UpsertId(p.ID, p)
	return err
}
