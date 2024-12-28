<script setup lang="ts">
import { ref, defineEmits, defineProps } from 'vue';
import type { Team } from '../types';
import { translations } from '../utils/translations';

const props = defineProps<{
  team?: Team;
}>();

const emit = defineEmits(['submit']);
const t = translations;

const formData = ref<Team>({
  name: props.team?.name ?? '',
  description: props.team?.description ?? '',
  leader_id: props.team?.leader_id
});

const submitForm = () => {
  emit('submit', { ...formData.value });
};
</script>

<template>
  <form @submit.prevent="submitForm" class="p-3">
    <div class="mb-3">
      <label class="form-label">{{ t.teams.fields.name }}</label>
      <input v-model="formData.name" type="text" class="form-control" required>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.teams.fields.description }}</label>
      <textarea v-model="formData.description" class="form-control" required></textarea>
    </div>
    <div class="mb-3">
      <label class="form-label">{{ t.teams.fields.leaderId }}</label>
      <input v-model.number="formData.leader_id" type="number" class="form-control">
    </div>
    <button type="submit" class="btn btn-primary">
      {{ props.team ? t.common.update : t.common.add }}
    </button>
  </form>
</template>