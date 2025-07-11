package mappers

import (
	"http-go/dtos"
	"http-go/ent"
)

type RollerCoasterMapper interface {
	RollerCoasterDTOToRollerCoasterEnt(dtos.RollerCoaster) *ent.RollerCoaster
	RollerCoasterEntToRollerCoasterDTO(*ent.RollerCoaster) dtos.RollerCoaster
}

func RollerCoasterDTOToRollerCoasterEnt(dto dtos.RollerCoaster) *ent.RollerCoaster {
	return &ent.RollerCoaster{
		Name:     dto.Name,
		Speed:    dto.Speed,
		Height:   dto.Height,
		Location: dto.Location,
	}
}

func RollerCoasterEntToRollerCoasterDTO(entity *ent.RollerCoaster) dtos.RollerCoaster {
	return dtos.RollerCoaster{
		Name:     entity.Name,
		Speed:    entity.Speed,
		Height:   entity.Height,
		Location: entity.Location,
	}
}
