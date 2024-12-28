import axios from 'axios';
import type { Employee, Team } from '../types';

const api = axios.create({
  baseURL: 'http://localhost:3000'
});

export const employeeService = {
  getAll: () => api.get<Employee[]>('/employees'),
  create: (employee: Omit<Employee, 'id'>) => api.post<Employee>('/employees', employee),
  update: (id: number, employee: Partial<Employee>) => api.put<Employee>(`/employees/${id}`, employee),
  delete: (id: number) => api.delete(`/employees/${id}`)
};

export const teamService = {
  getAll: () => api.get<Team[]>('/teams'),
  create: (team: Omit<Team, 'id' | 'member_count'>) => api.post<Team>('/teams', team),
  update: (id: number, team: Partial<Team>) => api.put<Team>(`/teams/${id}`, team),
  delete: (id: number) => api.delete(`/teams/${id}`)
};