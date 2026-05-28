const API_URL = 'http://localhost:8080/api/employees';

async function loadEmployees() {
    const tableBody = document.getElementById('employeeTableBody');
    tableBody.innerHTML = '<tr><td colspan="6" class="text-center">Cargando datos...</td></tr>';

    try {
        const response = await fetch(API_URL);
        const employees = await response.json();

        tableBody.innerHTML = '';

        if (employees.length === 0) {
            tableBody.innerHTML = '<tr><td colspan="6" class="text-center">No hay empleados registrados.</td></tr>';
            return;
        }

        employees.forEach(emp => {
            const row = `
                <tr>
                    <td><strong>${emp.employeeNumber}</strong></td>
                    <td>${emp.firstName}</td>
                    <td>${emp.lastName}</td>
                    <td>${emp.email}</td>
                    <td><span class="badge bg-info text-dark">${emp.jobTitle}</span></td>
                    <td>
                        <button class="btn btn-sm btn-danger" onclick="deleteEmployee(${emp.employeeNumber})">Eliminar</button>
                    </td>
                </tr>
            `;
            tableBody.innerHTML += row;
        });
    } catch (error) {
        console.error('Error al cargar empleados:', error);
        tableBody.innerHTML = '<tr><td colspan="6" class="text-center text-danger">Error conectando con el servidor Go.</td></tr>';
    }
}

async function deleteEmployee(id) {
    if (!confirm(`¿Estás seguro de que deseas eliminar al empleado #${id}?`)) return;

    try {
        const response = await fetch(`${API_URL}/${id}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            alert('Empleado eliminado correctamente');
            loadEmployees();
        } else {
            alert('Error al eliminar el empleado');
        }
    } catch (error) {
        console.error('Error:', error);
        alert('Fallo de conexión con el servidor');
    }
}

window.onload = loadEmployees;