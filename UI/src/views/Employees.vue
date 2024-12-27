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
    alert('خطا در نمایش کارمند ها');
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
    alert('خطا در ذخیره کارمند');
  }
};

const editEmployee = (employee: Employee) => {
  selectedEmployee.value = employee;
  showForm.value = true;
};

const deleteEmployee = async (id: number) => {
  if (confirm('آیا مطمئن هستید که می خواهید این کارمند را حذف کنید؟')) {
    try {
      await employeeService.delete(id);
      await loadEmployees();
    } catch (error) {
      alert('خطا در حذف کارمند');
    }
  }
};

onMounted(loadEmployees);
</script>

<template>
  <div class="container mt-4">
    <h2>مدیریت کارمند ها</h2>
    <button @click="showForm = !showForm" class="btn btn-primary mb-3">
      {{ showForm ? 'مخفی کردن' : 'اضافه کردن کارمند' }}
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
          <th>اسم</th>
          <th>ایمیل</th>
          <th>جایگاه</th>
          <th>تیم</th>
          <th>فعالیت ها</th>
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