package fisiks

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
	"log"
)

type Simulation struct {
	count     int64
	objects   map[string]*Object
	resources map[string]*ebiten.Image
	paused    bool

	// 1px = ... m
	Scale float64

	width  int
	height int
}

func NewSimulation(width int, height int) Simulation {
	return Simulation{
		objects:   make(map[string]*Object),
		resources: make(map[string]*ebiten.Image),
		width:     width,
		height:    height,
	}
}

func (s *Simulation) RegisterResource(id string, path string) {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatal(err)
	}
	s.resources[id] = img
}

func (s *Simulation) GetResource(id string) (*ebiten.Image, bool) {
	img, found := s.resources[id]
	return img, found
}

func (s *Simulation) RemoveResource(id string) {
	resource, found := s.resources[id]
	if !found {
		log.Fatalf("cannot remove an unexisting resource")
	}
	resource.Dispose()
	delete(s.resources, id)
}

func (s *Simulation) RegisterObject(obj *Object, id string) {
	if _, found := s.resources[obj.resource]; !found {
		log.Fatalf("unknown resource %s (while registering object %s)", obj.resource, id)
	}
	obj.x.pos = obj.OriginalX
	obj.y.pos = obj.OriginalY
	s.objects[id] = obj
}

func (s *Simulation) pixelHeight(obj *Object) float64 {
	return obj.Height / s.Scale
}

func (s *Simulation) pixelWidth(obj *Object) float64 {
	return obj.Height / s.Scale
}

func (s *Simulation) Time() float64 {
	return float64(s.count) / 60.
}

func verlet(f float64, ord *Ordinate, mass float64, scale float64) {
	dy := ord.velocity*dt + dt*dt*ord.acceleration/2
	ord.pos += dy / scale

	newAy := f / mass
	ord.velocity += (newAy + ord.acceleration) * dt / 2
}

func (s *Simulation) Update() error {
	keys := inpututil.AppendJustPressedKeys([]ebiten.Key{})
	if lo.Contains(keys, ebiten.KeySpace) {
		s.paused = !s.paused
	}

	if s.paused {
		return nil
	}

	for _, obj := range s.objects {
		if lo.Contains(keys, ebiten.KeyR) {
			obj.x = Ordinate{}
			obj.x.pos = obj.OriginalX
			obj.y = Ordinate{}
			obj.y.pos = obj.OriginalY
		}

		if s.pixelHeight(obj)+obj.y.pos >= float64(s.height) && obj.y.velocity > 0 {
			obj.y.velocity *= e
			obj.y.pos = float64(s.height) - s.pixelHeight(obj)
		}

		fx, fy := obj.ComputeForces()

		verlet(fx, &obj.x, obj.Mass, s.Scale)
		verlet(fy, &obj.y, obj.Mass, s.Scale)
	}

	s.count++
	return nil
}

func (s *Simulation) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%.2f", s.objects["croissant1"].y.pos))
	op := ebiten.DrawImageOptions{}
	for _, obj := range s.objects {
		resource := s.resources[obj.resource]
		op.GeoM.Reset()
		scaleX := s.pixelWidth(obj) / float64(resource.Bounds().Dx())
		scaleY := s.pixelHeight(obj) / (float64(resource.Bounds().Dy()))
		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(obj.x.pos, obj.y.pos)
		screen.DrawImage(s.resources[obj.resource], &op)
	}
}

func (s *Simulation) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return s.width, s.height
}
