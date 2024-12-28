<script setup lang="ts">
import { RouterLink } from 'vue-router';
import { translations } from '../utils/translations';

const t = translations;
</script>

<template>
  <div class="container mt-5" dir="rtl">
    <h1 class="text-center mb-4">{{ t.home.title }}</h1>
    <div class="d-flex justify-content-center gap-3">
      <RouterLink to="/employees" class="btn btn-primary">{{ t.employees.title }}</RouterLink>
      <RouterLink to="/teams" class="btn btn-primary">{{ t.teams.title }}</RouterLink>
    </div>
  </div>
</template>