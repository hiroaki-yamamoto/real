package message

import (
	"github.com/hiroaki-yamamoto/real/backend/rpc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server implements MessageServiceServer interface.
type Server struct {
	db *mongo.Database
}

// Subscribe handles subscribtions from users
func (me *Server) Subscribe(
	req *rpc.MessageRequest, stream rpc.MessageService_SubscribeServer,
) (err error) {
	chstream, err := me.db.Collection("messages").Watch(stream.Context(), bson.M{
		"$match": bson.M{"topicId": req.TopicId},
	})
	defer chstream.Close(stream.Context())
	return
}