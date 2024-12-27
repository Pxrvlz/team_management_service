<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { teamService } from '../services/api';
import type { Team } from '../types';
import TeamForm from '../components/TeamForm.vue';

const teams = ref<Team[]>([]);
const selectedTeam = ref<Team | null>(null);
const showForm = ref(false);

const loadTeams = async () => {
  try {
    const response = await teamService.getAll();
    teams.value = response.data;
  } catch (error) {
    alert('Failed to load teams');
  }
};

const handleSubmit = async (team: Team) => {
  try {
    if (selectedTeam.value) {
      await teamService.update(selectedTeam.value.id!, team);
    } else {
      await teamService.create(team);
    }
    await loadTeams();
    showForm.value = false;
    selectedTeam.value = null;
  } catch (error) {
    alert('Failed to save team');
  }
};

const editTeam = (team: Team) => {
  selectedTeam.value = team;
  showForm.value = true;
};

const deleteTeam = async (id: number) => {
  if (confirm('Are you sure you want to delete this team?')) {
    try {
      await teamService.delete(id);
      await loadTeams();
    } catch (error) {
      alert('Failed to delete team');
    }
  }
};

onMounted(loadTeams);
</script>

<template>
  <div class="container mt-4">
    <h2>Team Management</h2>
    <button @click="showForm = !showForm" class="btn btn-primary mb-3">
      {{ showForm ? 'Hide Form' : 'Add Team' }}
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
          <th>Name</th>
          <th>Description</th>
          <th>Leader</th>
          <th>Members</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="team in teams" :key="team.id">
          <td>{{ team.name }}</td>
          <td>{{ team.description }}</td>
          <td>{{ team.leader_name }}</td>
          <td>{{ team.member_count }}</td>
          <td>
            <button @click="editTeam(team)" class="btn btn-sm btn-warning me-2">Edit</button>
            <button @click="deleteTeam(team.id!)" class="btn btn-sm btn-danger">Delete</button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>