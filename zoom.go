package fractal

import (
	"fmt"
)

func (p Plane) FromSlidingFrames(zoomFactor float64, slideStep int, raiter func(Plane) float64) Plane {
	if zoomFactor <= 0.0 || zoomFactor > 1.0 {
		panic(fmt.Sprintf("invalid zoomFactor: %f", zoomFactor))
	}
	frameWith := int(float64(p.width) * zoomFactor)
	frameHeight := int(float64(p.height) * zoomFactor)
	bestFrame := p
	bestRaiting := 0.0
	for x := 0; x < (p.width - frameWith); x = x + slideStep {
		for y := 0; y < (p.height - frameHeight); y = y + slideStep {
			currentFrame := p.Crop(Box{x, y, frameWith, frameHeight})
			currentRaiting := raiter(currentFrame)
			if currentRaiting > bestRaiting {
				bestRaiting = currentRaiting
				bestFrame = currentFrame
			}
		}
	}
	return bestFrame.Scale(p.width, p.height)
}

func (p Plane) AutoZoom(zoomFactor float64, slideStep int) Plane {
	return p.FromSlidingFrames(zoomFactor, slideStep, Plane.Deviation)
}
