import flet as ft
import requests

API_URL = "http://localhost:8080/api/employees"

def main(page: ft.Page):
    page.title = "Sistema de Gestión de Empleados"
    page.theme_mode = ft.ThemeMode.LIGHT
    page.padding = 30

    table = ft.DataTable(
        columns=[
            ft.DataColumn(ft.Text("ID", weight=ft.FontWeight.BOLD)),
            ft.DataColumn(ft.Text("Nombre", weight=ft.FontWeight.BOLD)),
            ft.DataColumn(ft.Text("Apellido", weight=ft.FontWeight.BOLD)),
            ft.DataColumn(ft.Text("Email", weight=ft.FontWeight.BOLD)),
            ft.DataColumn(ft.Text("Puesto", weight=ft.FontWeight.BOLD)),
            ft.DataColumn(ft.Text("Acciones", weight=ft.FontWeight.BOLD)),
        ],
        rows=[]
    )

    def load_employees(e=None):
        try:
            response = requests.get(API_URL)
            if response.status_code == 200:
                employees = response.json()
                table.rows.clear()
                
                for emp in employees:
                    table.rows.append(
                        ft.DataRow(
                            cells=[
                                ft.DataCell(ft.Text(str(emp["employeeNumber"]))),
                                ft.DataCell(ft.Text(emp["firstName"])),
                                ft.DataCell(ft.Text(emp["lastName"])),
                                ft.DataCell(ft.Text(emp["email"])),
                                ft.DataCell(ft.Container(
                                    content=ft.Text(emp["jobTitle"], color=ft.Colors.WHITE),
                                    bgcolor=ft.Colors.BLUE_600,
                                    padding=5,
                                    border_radius=5
                                )),
                                ft.DataCell(
                                    ft.IconButton(
                                        icon=ft.Icons.DELETE,
                                        icon_color=ft.Colors.RED_400,
                                        tooltip="Eliminar",
                                        on_click=lambda e, emp_id=emp["employeeNumber"]: delete_employee(emp_id)
                                    )
                                ),
                            ]
                        )
                    )
                page.update()
        except Exception as ex:
            print(f"Error conectando con el servidor Go: {ex}")

    def delete_employee(emp_id):
        try:
            res = requests.delete(f"{API_URL}/{emp_id}")
            if res.status_code == 200:
                load_employees() 
        except Exception as ex:
            print(f"Error al eliminar: {ex}")

    header = ft.Row([
        ft.Text("🏢 Directorio de Empleados", size=28, weight=ft.FontWeight.BOLD, color=ft.Colors.BLUE_900),
        ft.ElevatedButton("🔄 Actualizar Lista", on_click=load_employees, bgcolor=ft.Colors.BLUE_600, color=ft.Colors.WHITE)
    ], alignment=ft.MainAxisAlignment.SPACE_BETWEEN)

    page.add(
        header,
        ft.Divider(height=20, color=ft.Colors.TRANSPARENT),
        table
    )
    load_employees()

if __name__ == "__main__":
    ft.app(target=main)