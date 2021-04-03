#include <vector>
#include <string>

using namespace std;

union ParVal {
    int iVal;
    double uVal;
    char sVal[64];
};

class Shape {
public:
    enum class EnumKind
    {
        Poly,
        Rect
    };
    Shape(const string& name, EnumKind kind);
    ~Shape();
    string GetName();
private:
    vector<string> parNameLst;
    vector<ParVal> parValLst;
    string name;
    EnumKind kind;
    double xPos;
    double yPos;
};

class Layer {
private:
    string name;
public:
    Layer(const string& name);
    ~Layer();
    string GetName();
};

class CellLayer:Layer {
private:
    vector<Shape> shapes;
public:
    CellLayer(Layer &techLayer);
    void AddShape();
};

class TechLib {
private:
    string name;
    vector<Layer> Layers;
public:
    TechLib();
    Layer GetLayer(const string &name);
};

class PCell {
private:
    string name;
    double xPos;
    double yPos;
    vector<CellLayer> layers;
    vector<string> parNameLst;
    vector<ParVal> parValLst;
public:
    PCell(const string& name);
    //~PCell();
    string GetName();
    int RegLayer(Layer &layer);
    int AddShape(const string& layerName, Shape &shape);
    void SetBasePos(double x, double y);
};
