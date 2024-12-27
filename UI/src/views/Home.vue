<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
  <div class="container mt-5">
    <h1 class="text-center mb-4">Employee and Team Management</h1>
    <div class="d-flex justify-content-center gap-3">
      <RouterLink to="/employees" class="btn btn-primary">Employee Management</RouterLink>
      <RouterLink to="/teams" class="btn btn-primary">Team Management</RouterLink>
    </div>
  </div>
</template>