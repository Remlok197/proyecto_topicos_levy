import { EmployeesResponseSchema } from '../schemas/employee.schema';

const API_URL = import.meta.env?.VITE_API_URL || 'http://localhost:8080/api';

export const getEmployees = async (page: number = 1, limit: number = 50) => {
  const response = await fetch(`${API_URL}/employees?page=${page}&limit=${limit}`);
  
  if (!response.ok) {
    throw new Error('Error al obtener los empleados');
  }

  const data = await response.json();
  
  // Validate API response with Zod
  return EmployeesResponseSchema.parse(data);
};
