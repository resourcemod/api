package entities

/**
#include "../../../../game/go/Entities.h"
*/
import "C"
import (
	"github.com/ungerik/go3d/quaternion"
	"github.com/ungerik/go3d/vec3"
)

type CollisionType int

const (
	None     CollisionType = 0
	BSP                    = 1
	BBOX                   = 2
	Sphere                 = 4
	Point                  = 5
	VPhysics               = 6
	Capsule                = 7
)

type EntityType string

const (
	Player              EntityType = ""
	PropDynamicOverride            = "prop_dynamic_override"
	PropPhysicsOverride            = "prop_physics_override"
)

type IEntity interface {
	Spawn(key int, model string) bool
	Kill(key int) bool
	SetModel(key int, model string) bool
	GetCollisionType(key int) CollisionType
	SetCollisionType(key int, collision CollisionType) bool
	SetColor(key int, color Color) bool
	GetColor(key int) Color
	GetAngle(key int) quaternion.T
	SetAngle(key int, angle quaternion.T) bool
	GetPosition(key int) vec3.T
	SetPosition(key int, coords vec3.T) bool
}

type Entity struct {
	key int
}

func (e Entity) Spawn(key int, model string) bool {
	// todo: implement
	return false
}

func (e Entity) Kill(key int) bool {
	// todo: implement
	return false
}

func (e Entity) SetModel(key int, model string) bool {
	// todo: implement
	return false
}

func (e Entity) GetCollisionType(key int) CollisionType {
	// todo: implement
	return BBOX
}

func (e Entity) SetCollisionType(key int, collision CollisionType) bool {
	// todo: implement
	return false
}

func (e Entity) SetColor(key int, color Color) bool {
	// todo: implement
	return false
}

func (e Entity) GetColor(key int) Color {
	// todo: implement
	return Color{0, 0, 0, 0}
}

func (e Entity) GetAngle(key int) quaternion.T {
	//todo: implement
	return quaternion.T{}
}

func (e Entity) SetAngle(key int, angle quaternion.T) bool {
	//todo: implement
	return false
}
func (e Entity) GetPosition(key int) vec3.T {
	//todo: implement
	return vec3.T{}
}

func (e Entity) SetPosition(key int, coords vec3.T) bool {
	// todo: implement
	return false
}

func NewEntity(entityType EntityType) IEntity {
	var entity Entity
	key := C.Create(*C.char(entityType))
	entity.key = key
	return entity
}
