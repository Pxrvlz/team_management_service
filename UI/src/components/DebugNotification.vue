<script setup lang="ts">
import { ref, onMounted } from 'vue';

interface Notification {
  id: number;
  level: 'INFO' | 'WARNING' | 'ERROR';
  message: string;
  context: string;
  timestamp: string;
  data?: any;
}

const notifications = ref<Notification[]>([]);
const autoHideTimeout = 5000; // 5 seconds

const addNotification = (notification: Omit<Notification, 'id'>) => {
  const id = Date.now();
  notifications.value.push({ ...notification, id });
  
  // Auto-hide after timeout
  setTimeout(() => {
    removeNotification(id);
  }, autoHideTimeout);
};

const removeNotification = (id: number) => {
  notifications.value = notifications.value.filter(n => n.id !== id);
};

// Expose the addNotification method globally
onMounted(() => {
  window.debugNotify = addNotification;
});

// Color mapping for different levels
const levelColors = {
  INFO: 'bg-blue-100 border-blue-500 text-blue-700',
  WARNING: 'bg-yellow-100 border-yellow-500 text-yellow-700',
  ERROR: 'bg-red-100 border-red-500 text-red-700'
};
</script>

<template>
  <div class="fixed top-4 right-4 z-50 space-y-2 max-w-md">
    <transition-group name="notification">
      <div
        v-for="notification in notifications"
        :key="notification.id"
        :class="[
          'p-4 rounded border-l-4 shadow-md cursor-pointer',
          levelColors[notification.level]
        ]"
        @click="removeNotification(notification.id)"
      >
        <div class="flex justify-between items-start">
          <div class="font-bold">{{ notification.context }}</div>
          <div class="text-sm opacity-75">{{ notification.level }}</div>
        </div>
        <div class="mt-1">{{ notification.message }}</div>
        <div v-if="notification.data" class="mt-2 text-sm opacity-75">
          <pre class="whitespace-pre-wrap">{{ JSON.stringify(notification.data, null, 2) }}</pre>
        </div>
        <div class="text-xs mt-2 opacity-50">
          {{ new Date(notification.timestamp).toLocaleTimeString() }}
        </div>
      </div>
    </transition-group>
  </div>
</template>

<style scoped>
.notification-enter-active,
.notification-leave-active {
  transition: all 0.3s ease;
}

.notification-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.notification-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>