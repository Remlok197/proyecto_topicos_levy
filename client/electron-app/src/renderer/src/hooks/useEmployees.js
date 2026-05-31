import { useState, useEffect, useCallback } from 'react';
import { getEmployees } from '../services/employee.service';
import { createEmployeeListAdapter } from '../adapters/employee.adapter';

export const useEmployees = (page = 1, limit = 50) => {
  const [employees, setEmployees] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchEmployees = useCallback(async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await getEmployees(page, limit);
      const adaptedEmployees = createEmployeeListAdapter(response.data);
      setEmployees(adaptedEmployees);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ocurrió un error desconocido');
      console.error('Error fetching employees:', err);
    } finally {
      setLoading(false);
    }
  }, [page, limit]);

  useEffect(() => {
    fetchEmployees();
  }, [fetchEmployees]);

  return { employees, loading, error, setEmployees };
};
