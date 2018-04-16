package main

import (
	"net"
	"google.golang.org/grpc"
	"github.com/ievgen-ma/tcp-chat/protocol"
	"log"
	"context"
	"github.com/twinj/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type reportService struct{}

type mongo struct {
	Tasks *mgo.Collection
}

var DB *mongo

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	DB = &mongo{session.DB("report").C("tasks")}
}

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("could not listen to :3000: %v", err)
	}

	s := reportService{}
	grpcServer := grpc.NewServer()
	protocol.RegisterReportServer(grpcServer, s)

	if err := grpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func (reportService) Create(ctx context.Context, task *protocol.Task) (*protocol.ID, error) {
	task.Id = uuid.NewV4().String()
	return &protocol.ID{Id:task.Id}, DB.Tasks.Insert(task)
}

func (reportService) FindOne(ctx context.Context, protocolID *protocol.ID) (*protocol.Task, error) {
	task := protocol.Task{}
	return &task, DB.Tasks.Find(bson.M{"id":protocolID.Id}).One(&task)
}

func (reportService) FindAll(ctx context.Context, void *protocol.Void) (*protocol.Tasks, error) {
	tasks := protocol.Tasks{}
	return &tasks, DB.Tasks.Find(nil).All(&tasks.Task)
}

func (reportService) Update(ctx context.Context, task *protocol.Task) (*protocol.ID, error) {
	return &protocol.ID{Id:task.Id}, DB.Tasks.Update(bson.M{"id":task.Id}, bson.M{"$set":bson.M{"name":task.Name}})
}

func (reportService) Delete(ctx context.Context, protocolID *protocol.ID) (*protocol.ID, error) {
	return &protocol.ID{Id:protocolID.Id}, DB.Tasks.Remove(bson.M{"id":protocolID.Id})
}
