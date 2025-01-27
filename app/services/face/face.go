package face

import (
	"errors"
	"fmt"
	"log"
	"math"
	"sync"

	"github.com/Kagami/go-face"
)

type (
	FaceMatcher struct{ r *face.Recognizer }
	MatchResult struct {
		IsMatch       bool    // Whether it's considered a match
		Confidence    float32 // Confidence score (0-1, where 1 is perfect match)
		DistanceScore float32 // Raw distance score (lower is better)
	}
)

var matcher *FaceMatcher
var once sync.Once

func GetFaceMatcher() *FaceMatcher {

	once.Do(func() {
		m, err := newFaceMatcher("app/services/face/models")
		if err != nil {
			log.Fatal(err.Error())
		}

		matcher = m
	})

	return matcher
}

func newFaceMatcher(modelsDir string) (m *FaceMatcher, err error) {
	r, err := face.NewRecognizer(modelsDir)
	if err != nil {
		return
	}

	m = &FaceMatcher{r}

	return
}

// CompareFaceImages compares faces in two images and returns match results
func (fm *FaceMatcher) CompareFaceImages(refImg, targetImg []byte) (MatchResult, error) {
	d1, err := fm.getFaceDescriptor(refImg)
	if err != nil {
		return MatchResult{}, fmt.Errorf("error processing reference image: %v", err)
	}

	d2, err := fm.getFaceDescriptor(targetImg)
	if err != nil {
		return MatchResult{}, fmt.Errorf("error processing target image: %v", err)
	}

	return fm.compareDescriptors(d1, d2), nil
}

func (fm *FaceMatcher) Close() {
	if fm.r != nil {
		fm.r.Close()
	}
}

func (fm *FaceMatcher) getFaceDescriptor(imageData []byte) (d *face.Descriptor, err error) {
	face, err := fm.r.RecognizeSingle(imageData)
	if err != nil {
		return
	}

	if face == nil {
		return nil, errors.New("no face detected in image")
	}

	return &face.Descriptor, nil
}

// compareDescriptors compares two face descriptors and returns detailed match results
func (fm *FaceMatcher) compareDescriptors(desc1, desc2 *face.Descriptor) MatchResult {
	// Calculate the Euclidean distance between the face descriptors
	var sum float32
	for i := 0; i < 128; i++ {
		diff := desc1[i] - desc2[i]
		sum += diff * diff
	}
	distance := float32(math.Sqrt(float64(sum)))

	// Convert distance to confidence score (0-1)
	// These thresholds are based on common practices but may need adjustment
	const (
		perfectMatch   = 0.0
		noMatch        = 1.0
		matchThreshold = 0.6 // Threshold for considering it a match
	)

	// Convert distance to confidence (inverse relationship)
	// When distance is 0, confidence is 1 (perfect match)
	// When distance is 1 or greater, confidence approaches 0
	confidence := 1.0 - math.Min(float64(distance), 1.0)

	return MatchResult{
		IsMatch:       confidence >= matchThreshold,
		Confidence:    float32(confidence),
		DistanceScore: distance,
	}
}
