<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { employeeService } from '../services/api';
import type { Employee } from '../types';
import EmployeeForm from '../components/EmployeeForm.vue';
import { RouterLink } from 'vue-router';
import { translations } from '../utils/translations';
import { handleApiError } from '../utils/errorHandler';

const employees = ref<Employee[]>([]);
const selectedEmployee = ref<Employee | null>(null);
const showForm = ref(false);
const t = translations;

const loadEmployees = async () => {
  try {
    const response = await employeeService.getAll();
    employees.value = response.data;
  } catch (error) {
    alert(handleApiError(error, 'loadEmployees'));
  }
};

const handleSubmit = async (employee: Employee) => {
  try {
    if (selectedEmployee.value) {
      await employeeService.update(selectedEmployee.value.id!, employee);
    } else {
      const response = await employeeService.create(employee);
      if (!response.data.id) {
        throw new Error('خطا در ایجاد کارمند جدید');
      }
    }
    await loadEmployees();
    showForm.value = false;
    selectedEmployee.value = null;
  } catch (error) {
    alert(handleApiError(error, 'handleSubmit'));
  }
};

const editEmployee = (employee: Employee) => {
  selectedEmployee.value = employee;
  showForm.value = true;
};

const deleteEmployee = async (id: number) => {
  if (confirm(t.employees.deleteConfirm)) {
    try {
      await employeeService.delete(id);
      await loadEmployees();
    } catch (error) {
      alert(handleApiError(error, 'deleteEmployee'));
    }
  }
};

onMounted(loadEmployees);
</script>

<template>
  <div class="container mt-4" dir="rtl">
    <div class="d-flex justify-content-between align-items-center mb-4">
      <h2>{{ t.employees.title }}</h2>
      <RouterLink to="/" class="btn btn-secondary">{{ t.common.home }}</RouterLink>
    </div>

    <button @click="showForm = !showForm" class="btn btn-primary mb-3">
      {{ showForm ? t.employees.hideForm : t.employees.add }}
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
          <th>{{ t.employees.fields.id }}</th>
          <th>{{ t.employees.fields.name }}</th>
          <th>{{ t.employees.fields.email }}</th>
          <th>{{ t.employees.fields.position }}</th>
          <th>{{ t.employees.fields.team }}</th>
          <th>{{ t.common.actions }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="employee in employees" :key="employee.id">
          <td>{{ employee.id }}</td>
          <td>{{ employee.name }}</td>
          <td>{{ employee.email }}</td>
          <td>{{ employee.position }}</td>
          <td>{{ employee.team_name }}</td>
          <td>
            <button @click="editEmployee(employee)" class="btn btn-sm btn-warning me-2">{{ t.common.edit }}</button>
            <button @click="deleteEmployee(employee.id!)" class="btn btn-sm btn-danger">{{ t.common.delete }}</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>