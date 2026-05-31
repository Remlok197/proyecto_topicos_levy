import { EmployeeAPI } from '../schemas/employee.schema';

export interface Employee {
  employeeNumber: number;
  firstName: string;
  lastName: string;
  jobTitle: string;
  salary: number;
  avatarColor: string;
}

// Function to generate a deterministic color based on employee number
const generateAvatarColor = (id: number): string => {
  const colors = [
    'bg-teal-100 text-teal-700',
    'bg-blue-100 text-blue-700',
    'bg-purple-100 text-purple-700',
    'bg-amber-100 text-amber-700',
    'bg-rose-100 text-rose-700',
    'bg-indigo-100 text-indigo-700',
    'bg-pink-100 text-pink-700',
    'bg-emerald-100 text-emerald-700'
  ];
  return colors[id % colors.length];
};

export const createEmployeeAdapter = (apiEmployee: EmployeeAPI): Employee => {
  return {
    employeeNumber: apiEmployee.employeeNumber,
    firstName: apiEmployee.firstName,
    lastName: apiEmployee.lastName,
    jobTitle: apiEmployee.jobTitle,
    salary: apiEmployee.salary,
    avatarColor: generateAvatarColor(apiEmployee.employeeNumber)
  };
};

export const createEmployeeListAdapter = (apiEmployees: EmployeeAPI[]): Employee[] => {
  return apiEmployees.map(createEmployeeAdapter);
};
