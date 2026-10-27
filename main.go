package main

import "fmt"

type Producto struct {
	ID       int
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(id int, nombre string, precio float64, cantidad int) Producto {
	return Producto{
		ID:       id,
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
}

func (p *Producto) ActualizarPrecio(nuevoPrecio float64) {
	if nuevoPrecio >= 0 {
		p.Precio = nuevoPrecio
		fmt.Println("Precio actualizado correctamente.")
	} else {
		fmt.Println("El precio no puede ser negativo.")
	}
}

func (p *Producto) ActualizarCantidad(nuevaCantidad int) {
	if nuevaCantidad >= 0 {
		p.Cantidad = nuevaCantidad
		fmt.Println("Cantidad actualizada correctamente.")
	} else {
		fmt.Println("La cantidad no puede ser negativa.")
	}
}

func (p Producto) MostrarInfo() {
	fmt.Println("----------------------------")
	fmt.Println("ID:", p.ID)
	fmt.Println("Nombre:", p.Nombre)
	fmt.Printf("Precio: $%.2f\n", p.Precio)
	fmt.Println("Cantidad disponible:", p.Cantidad)
	fmt.Println("----------------------------")
}

type Inventario struct {
	Productos map[int]Producto
}

func NuevoInventario() Inventario {
	return Inventario{
		Productos: make(map[int]Producto),
	}
}

func (i *Inventario) AgregarProducto(p Producto) {
	if _, existe := i.Productos[p.ID]; existe {
		fmt.Println("Ya existe un producto con ese ID.")
		return
	}

	i.Productos[p.ID] = p
	fmt.Println("Producto agregado correctamente.")
}

func (i *Inventario) EliminarProducto(id int) {
	if _, existe := i.Productos[id]; existe {
		delete(i.Productos, id)
		fmt.Println("Producto eliminado correctamente.")
	} else {
		fmt.Println("Producto no encontrado.")
	}
}

func (i Inventario) BuscarProducto(id int) (Producto, bool) {
	producto, existe := i.Productos[id]
	return producto, existe
}

func (i Inventario) ListarProductos() {
	if len(i.Productos) == 0 {
		fmt.Println("El inventario esta vacio.")
		return
	}

	fmt.Println("\nLISTA DE PRODUCTOS")

	for _, producto := range i.Productos {
		producto.MostrarInfo()
	}
}

func main() {
	inventario := NuevoInventario()

	p1 := NuevoProducto(1, "Arroz", 1.25, 50)
	p2 := NuevoProducto(2, "Leche", 0.95, 30)
	p3 := NuevoProducto(3, "Pan", 0.50, 20)

	inventario.AgregarProducto(p1)
	inventario.AgregarProducto(p2)
	inventario.AgregarProducto(p3)

	inventario.ListarProductos()

	fmt.Println("\nACTUALIZAR PRODUCTO")
	if producto, existe := inventario.BuscarProducto(1); existe {
		producto.ActualizarPrecio(1.50)
		producto.ActualizarCantidad(60)

		inventario.Productos[producto.ID] = producto
		producto.MostrarInfo()
	}

	fmt.Println("\nBUSCAR PRODUCTO")
	if producto, existe := inventario.BuscarProducto(2); existe {
		producto.MostrarInfo()
	} else {
		fmt.Println("Producto no encontrado.")
	}

	fmt.Println("\nELIMINAR PRODUCTO")
	inventario.EliminarProducto(3)

	inventario.ListarProductos()
}
