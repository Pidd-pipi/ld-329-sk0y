<template>
  <header class="app-header">
    <div>
      <p class="eyebrow">Skill Swap Campus</p>
      <h1>{{ APP_NAME }}</h1>
      <p>{{ APP_TAGLINE }}</p>
    </div>
    <div class="app-header__side">
      <el-select
        :model-value="currentUser"
        class="actor-select"
        @change="onActorChange"
      >
        <el-option v-for="name in users" :key="name" :label="`当前身份：${name}`" :value="name" />
      </el-select>
      <el-badge :value="unread" class="message-badge">
        <el-button type="primary">站内消息</el-button>
      </el-badge>
    </div>
  </header>
</template>

<script setup lang="ts">
import { APP_NAME, APP_TAGLINE } from '../constants/app.constants';

defineProps<{ unread: number; currentUser: string; users: string[] }>();
const emit = defineEmits<{ (e: 'change-actor', name: string): void }>();

function onActorChange(name: string) {
  emit('change-actor', name);
}
</script>

<style scoped>
.app-header__side { display: flex; align-items: center; gap: 14px; }
.actor-select { width: 168px; }
</style>
