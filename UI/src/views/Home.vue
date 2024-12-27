<script setup lang="ts">
import { RouterLink } from 'vue-router';
</script>

<template>
  <div class="container mt-5">
    <h1 class="text-center mb-4">مدیریت کارمند ها و تیم ها</h1>
    <div class="d-flex justify-content-center gap-3">
      <RouterLink to="/employees" class="btn btn-primary">مدیریت کارمند ها</RouterLink>
      <RouterLink to="/teams" class="btn btn-primary">مدیریت تیم ها</RouterLink>
    </div>
  </div>
</template>