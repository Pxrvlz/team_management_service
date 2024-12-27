export interface Employee {
  id?: number;
  name: string;
  email: string;
  phone: string;
  position: string;
  salary: number;
  age: number;
  hire_date?: string;
  team_id?: number;
  team_name?: string;
}

export interface Team {
  id?: number;
  name: string;
  description: string;
  leader_id?: number;
  leader_name?: string;
  creation_date?: string;
  member_count?: number;
}