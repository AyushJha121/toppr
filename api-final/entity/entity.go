package entity

//Subject Entity
type Subject struct {
	SubjectID int
	SubName   string
	Grade     int
}

//Topics Entity
type Topic struct {
	TopicID   int
	TopicName string
	SubjectID int
}

//SubTopic Entity
type SubTopic struct {
	SubtopicID int
	ST_Name    string
	TopicID    int
}

//Concept Entity
type Concept struct {
	ConceptID   int
	Name        string
	Description string
	SubTopicID  int
}

//Video Segment Entity
type VideoSegment struct {
	SegmentID int
	Name      string
	Duration  int
	ConceptID int
}

//Video Entity
type Video struct {
	SegmentID int
	URL       string
}

//Topic Child Entity
type TopicChild struct {
	TopicName   string
	ST_Name     string
	ConceptName string
	SegmentName string
	Duration    int
	URL         string
}
