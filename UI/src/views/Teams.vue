<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { teamService } from '../services/api';
import type { Team } from '../types';
import TeamForm from '../components/TeamForm.vue';
import { RouterLink } from 'vue-router';
import { translations } from '../utils/translations';
import { handleApiError } from '../utils/errorHandler';

const teams = ref<Team[]>([]);
const selectedTeam = ref<Team | null>(null);
const showForm = ref(false);
const t = translations;

const loadTeams = async () => {
  try {
    const response = await teamService.getAll();
    teams.value = response.data;
  } catch (error) {
    alert(handleApiError(error, 'loadTeams'));
  }
};

const handleSubmit = async (team: Team) => {
  try {
    if (selectedTeam.value) {
      await teamService.update(selectedTeam.value.id!, team);
    } else {
      const response = await teamService.create(team);
      if (!response.data.id) {
        throw new Error('خطا در ایجاد تیم جدید');
      }
    }
    await loadTeams();
    showForm.value = false;
    selectedTeam.value = null;
  } catch (error) {
    alert(handleApiError(error, 'handleSubmit'));
  }
};

const editTeam = (team: Team) => {
  selectedTeam.value = team;
  showForm.value = true;
};

const deleteTeam = async (id: number) => {
  if (confirm(t.teams.deleteConfirm)) {
    try {
      await teamService.delete(id);
      await loadTeams();
    } catch (error) {
      alert(handleApiError(error, 'deleteTeam'));
    }
  }
};

onMounted(loadTeams);
</script>

<template>
  <div class="container mt-4" dir="rtl">
    <div class="d-flex justify-content-between align-items-center mb-4">
      <h2>{{ t.teams.title }}</h2>
      <RouterLink to="/" class="btn btn-secondary">{{ t.common.home }}</RouterLink>
    </div>

    <button @click="showForm = !showForm" class="btn btn-primary mb-3">
      {{ showForm ? t.teams.hideForm : t.teams.add }}
    </button>
    
    <div v-if="showForm">
      <TeamForm 
        :team="selectedTeam"
        @submit="handleSubmit"
      />
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>{{ t.teams.fields.id }}</th>
          <th>{{ t.teams.fields.name }}</th>
          <th>{{ t.teams.fields.description }}</th>
          <th>{{ t.teams.fields.leader }}</th>
          <th>{{ t.teams.fields.members }}</th>
          <th>{{ t.common.actions }}</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="team in teams" :key="team.id">
          <td>{{ team.id }}</td>
          <td>{{ team.name }}</td>
          <td>{{ team.description }}</td>
          <td>{{ team.leader_name }}</td>
          <td>{{ team.member_count }}</td>
          <td>
            <button @click="editTeam(team)" class="btn btn-sm btn-warning me-2">{{ t.common.edit }}</button>
            <button @click="deleteTeam(team.id!)" class="btn btn-sm btn-danger">{{ t.common.delete }}</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>