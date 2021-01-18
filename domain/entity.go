package domain

type Entity struct {
	Id int
}

func (entity *Entity) SetId(id int) {
	entity.Id = id
}
