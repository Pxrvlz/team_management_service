<script setup lang="ts">
import { ref, defineEmits, defineProps } from 'vue';
import type { Employee } from '../types';
import { translations } from '../utils/translations';

const props = defineProps<{
  employee?: Employee;
}>();

const emit = defineEmits(['submit']);
const t = translations;

const formData = ref<Employee>({
  name: props.employee?.name ?? '',
  email: props.employee?.email ?? '',
  phone: props.employee?.phone ?? '',
  position: props.employee?.position ?? '',
  salary: props.employee?.salary ?? 0,
  age: props.employee?.age ?? 0,
  team_id: props.employee?.team_id
});

const submitForm = () => {
  emit('submit', { ...formData.value });
};
</script>

<template>
  <form @submit.prevent="submitForm" class="p-3">
    <div class="mb-3">
      <label class="form-label">{{ t.employees.fields.name }}</label>
      <input v-model="formData.name" type="text" class="form-control" required>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.employees.fields.email }}</label>
      <input v-model="formData.email" type="email" class="form-control" required>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.employees.fields.phone }}</label>
      <input v-model="formData.phone" type="tel" class="form-control" required>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.employees.fields.position }}</label>
      <input v-model="formData.position" type="text" class="form-control" required>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.employees.fields.salary }}</label>
      <input v-model.number="formData.salary" type="number" class="form-control" required>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.employees.fields.age }}</label>
      <input v-model.number="formData.age" type="number" class="form-control" required>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.employees.fields.teamId }}</label>
      <input v-model.number="formData.team_id" type="number" class="form-control">
    </div>
    <button type="submit" class="btn btn-primary">
      {{ props.employee ? t.common.update : t.common.add }}
    </button>
  </form>
</template>