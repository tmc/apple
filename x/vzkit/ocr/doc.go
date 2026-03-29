// Package ocr provides text recognition using Apple's Vision framework.
//
// It wraps VNRecognizeTextRequest to detect text in images and return
// observations with pixel coordinates, bounding boxes, and confidence
// scores. Search functions support region filtering and ranking
// preferences for locating specific text on screen.
//
// Basic usage:
//
//	svc := ocr.NewService(false)
//	obs, err := svc.RecognizeText(img)
//	for _, o := range obs {
//		fmt.Printf("%s at (%d,%d)\n", o.Text, o.Center.X, o.Center.Y)
//	}
//
// Finding text by content:
//
//	x, y, found := svc.FindText(img, "Continue")
package ocr
