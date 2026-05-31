import { z } from 'zod';

export const EmployeeSchema = z.object({
  employeeNumber: z.number(),
  firstName: z.string(),
  jobTitle: z.string(),
  lastName: z.string(),
  salary: z.number()
});

export const EmployeesResponseSchema = z.object({
  data: z.array(EmployeeSchema),
  limit: z.number(),
  page: z.number()
});

export type EmployeeAPI = z.infer<typeof EmployeeSchema>;
export type EmployeesResponseAPI = z.infer<typeof EmployeesResponseSchema>;
