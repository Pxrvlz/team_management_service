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
    alert('خطا در نمایش تیم ها');
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
    alert('خطا در ذخیره تیم');
  }
};

const editTeam = (team: Team) => {
  selectedTeam.value = team;
  showForm.value = true;
};

const deleteTeam = async (id: number) => {
  if (confirm('آیا مطمئن هستید که می خواهید این تیم را حذف کنید؟')) {
    try {
      await teamService.delete(id);
      await loadTeams();
    } catch (error) {
      alert('خطا در حذف تیم');
    }
  }
};

onMounted(loadTeams);
</script>

<template>
  <div class="container mt-4">
    <h2>مدیریت تیم ها</h2>
    <button @click="showForm = !showForm" class="btn btn-primary mb-3">
      {{ showForm ? 'مخفی کردن' : 'اضافه کردن تیم' }}
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
          <th>اسم</th>
          <th>توضیحات</th>
          <th>سرگروه</th>
          <th>عضوها</th>
          <th>فعالیت ها</th>
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