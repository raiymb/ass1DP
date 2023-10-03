package abstractFactory

type IFurnitureFactory interface {
	CreateChair() IChair
	CreateTable() ITable
}
