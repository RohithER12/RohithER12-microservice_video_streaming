package interfaces

import "github.com/RohithER12/video_service/pkg/pb"

type VideoRepo interface {
	CreateVideoid(string) error
	FindAllVideo() ([]*pb.VideoID, error)
}
