package data

import (
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type SubscriptionFeed struct {
	FeedKey *datastore.Key
	Created time.Time
}

type Subscription struct {
	UserID   string
	FeedList []SubscriptionFeed

	Created time.Time
	Updated time.Time
}

func SubscribeFeed(c appengine.Context, u *user.User, feed *datastore.Key) (*datastore.Key, *Subscription, error) {
	s := new(Subscription)
	s.UserID = u.ID

	key := s.NewKey(c)
	err := datastore.Get(c, key, s)

	item := SubscriptionFeed{}
	item.FeedKey = feed
	item.Created = time.Now()

	if err == datastore.ErrNoSuchEntity {
		s.FeedList = []SubscriptionFeed{item}
		s.Created = time.Now()
	} else if err == nil {
		if !s.IsSubscribed(c, feed) {
			s.FeedList = append(s.FeedList, item)
		}
	} else {
		return nil, nil, err
	}

	s.UserID = u.ID
	s.Updated = time.Now()

	key, err = datastore.Put(c, key, s)
	return key, s, err
}

func (s *Subscription) NewKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Subscription", s.UserID, 0, nil)
}

func (s *Subscription) IsSubscribed(c appengine.Context, feed *datastore.Key) bool {
	id := feed.StringID()
	for _, item := range s.FeedList {
		if id == item.FeedKey.StringID() {
			return true
		}
	}
	return false
}

func GetSubscription(c appengine.Context, u *user.User) (*datastore.Key, *Subscription, error) {
	s := new(Subscription)
	s.UserID = u.ID

	key := s.NewKey(c)
	err := datastore.Get(c, key, s)
	if err != nil {
		return nil, nil, err
	}

	return key, s, nil
}
