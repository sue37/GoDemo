#include "PCell.hpp"

PCell::PCell(const string& name)
{
    this->name = name;
    this->xPos = 0;
    this->yPos = 0;
}

//PCell::~PCell()
//{
//}

string PCell::GetName()
{
    return this->name;
}

int PCell::RegLayer(Layer& techLayer)
{
    auto layer = new CellLayer(techLayer);
    this->layers.push_back(*layer);
    return 0;
}

int PCell::AddShape(const string& layerName, Shape& shape)
{
    return 0;
}

void PCell::SetBasePos(double x, double y)
{
    this->xPos = x;
    this->yPos = y;
}

Layer::Layer(const string& name)
{
    this->name = name;
}

Layer::~Layer()
{
}

string Layer::GetName()
{
    return this->name;
}

Shape::Shape(const string& name, EnumKind kind)
{
    this->xPos = 0.0;
    this->yPos = 0.0;
    this->kind = EnumKind::Rect;
}

Shape::~Shape()
{
}

string Shape::GetName()
{
    return this->name;
}

CellLayer::CellLayer(Layer &techLayer) : Layer(techLayer.GetName())
{
}

void CellLayer::AddShape()
{
}

TechLib::TechLib()
{
}

Layer TechLib::GetLayer(const string& name)
{
    //return Layer;
    return NULL;
}
