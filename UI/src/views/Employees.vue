<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { employeeService } from '../services/api';
import type { Employee } from '../types';
import EmployeeForm from '../components/EmployeeForm.vue';

const employees = ref<Employee[]>([]);
const selectedEmployee = ref<Employee | null>(null);
const showForm = ref(false);

const loadEmployees = async () => {
  try {
    const response = await employeeService.getAll();
    employees.value = response.data;
  } catch (error) {
    alert('Failed to load employees');
  }
};

const handleSubmit = async (employee: Employee) => {
  try {
    if (selectedEmployee.value) {
      await employeeService.update(selectedEmployee.value.id!, employee);
    } else {
      await employeeService.create(employee);
    }
    await loadEmployees();
    showForm.value = false;
    selectedEmployee.value = null;
  } catch (error) {
    alert('Failed to save employee');
  }
};

const editEmployee = (employee: Employee) => {
  selectedEmployee.value = employee;
  showForm.value = true;
};

const deleteEmployee = async (id: number) => {
  if (confirm('Are you sure you want to delete this employee?')) {
    try {
      await employeeService.delete(id);
      await loadEmployees();
    } catch (error) {
      alert('Failed to delete employee');
    }
  }
};

onMounted(loadEmployees);
</script>

<template>
  <div class="container mt-4">
    <h2>Employee Management</h2>
    <button @click="showForm = !showForm" class="btn btn-primary mb-3">
      {{ showForm ? 'Hide Form' : 'Add Employee' }}
    </button>
    
    <div v-if="showForm">
      <EmployeeForm 
        :employee="selectedEmployee"
        @submit="handleSubmit"
      />
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>Name</th>
          <th>Email</th>
          <th>Position</th>
          <th>Team</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="employee in employees" :key="employee.id">
          <td>{{ employee.name }}</td>
          <td>{{ employee.email }}</td>
          <td>{{ employee.position }}</td>
          <td>{{ employee.team_name }}</td>
          <td>
            <button @click="editEmployee(employee)" class="btn btn-sm btn-warning me-2">Edit</button>
            <button @click="deleteEmployee(employee.id!)" class="btn btn-sm btn-danger">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>