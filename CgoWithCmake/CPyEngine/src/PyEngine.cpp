// PyEmbedTest.cpp : Defines the entry point for the application.
//
#include <pybind11/embed.h> // everything needed for embedding
namespace py = pybind11;

#include "PyEngine.h"
#include "./CellLayout/PCell.hpp"
using namespace std;

PYBIND11_EMBEDDED_MODULE(PCellBase, m) {
	py::class_<PCell>(m, "PCell")
		.def(py::init<const std::string&>())
		.def("getName", &PCell::GetName)
		.def("regLayer", &PCell::RegLayer)
		.def("setBasePos", &PCell::SetBasePos);
	py::class_<Layer>(m, "Layer")
		.def(py::init<const std::string&>())
		.def("getName", &Layer::GetName);
	//.def("addShape", &Layer::AddShape);
	py::class_<Shape> shape(m, "Shape");
	py::enum_<Shape::EnumKind>(shape, "Kind")
		.value("Poly", Shape::EnumKind::Poly)
		.value("Rect", Shape::EnumKind::Rect)
		.export_values();

	shape
		.def(py::init<const std::string&, Shape::EnumKind>())
		.def("getName", &Shape::GetName);
}

PyEngine::PyEngine()
{
#ifdef DEBUG
	cout << "Hello from PyEngine" << endl;
#endif // DEBUG
	try {
		py::exec("import PCellBase as pcb");
	}
	catch (exception e) {
		cout << e.what() << endl;
	}
}

void PyEngine::StartEngine()
{
	string cmd;
	while (true) {
		cin >> cmd;
		if (cmd == "exit") {
			return;
		}
		try {
			py::exec(cmd);
		}
		catch (exception e) {
			cout << e.what() << endl;
		}
	}
}
