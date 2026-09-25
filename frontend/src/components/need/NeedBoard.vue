<template>
  <div class="need-board">
    <div class="need-board__head">
      <h2>需求与响应</h2>
      <el-radio-group v-model="filter" size="small">
        <el-radio-button value="all">全部</el-radio-button>
        <el-radio-button value="open">招募中</el-radio-button>
        <el-radio-button value="mine">与我相关</el-radio-button>
      </el-radio-group>
    </div>

    <NeedBoardCard
      v-for="need in filteredNeeds"
      :key="need.id"
      :need="need"
      :slots="store.slots"
      :current-user="store.actor"
      :busy-key="store.busyKey"
      @respond-click="openRespond(need.id)"
      @select="onSelect(need.id, $event.id)"
      @confirm="store.confirm($event)"
      @cancel="store.cancel($event)"
    />

    <el-empty v-if="filteredNeeds.length === 0" description="暂无符合条件的需求" />

    <NeedResponseDialog
      v-model="dialogVisible"
      :slots="store.slots"
      :submitting="store.busyKey !== ''"
      @submit="onSubmit"
    />
  </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue';
import { ElMessageBox } from 'element-plus';
import NeedBoardCard from './NeedBoardCard.vue';
import NeedResponseDialog from './NeedResponseDialog.vue';
import { NEED_STATUS_OPEN } from '../../constants/need.constants';
import { useNeedStore } from '../../stores/need.store';

const store = useNeedStore();
const filter = ref<'all' | 'open' | 'mine'>('all');
const dialogVisible = ref(false);
const activeNeedId = ref<number | null>(null);

const filteredNeeds = computed(() => {
  const needs = store.overview?.needs ?? [];
  if (filter.value === 'open') {
    return needs.filter((n) => n.status === NEED_STATUS_OPEN);
  }
  if (filter.value === 'mine') {
    return needs.filter(
      (n) =>
        n.requester === store.actor ||
        n.myResponse !== undefined ||
        n.appointment?.requester === store.actor ||
        n.appointment?.respondent === store.actor,
    );
  }
  return needs;
});

function openRespond(needId: number) {
  activeNeedId.value = needId;
  dialogVisible.value = true;
}

function onSubmit(payload: { offerNote: string; freeSlots: string[] }) {
  if (activeNeedId.value === null) return;
  const needId = activeNeedId.value;
  void store.submitResponse(needId, payload, () => {
    dialogVisible.value = false;
  });
}

async function onSelect(needId: number, responseId: number) {
  try {
    await ElMessageBox.confirm(
      '选定后将按需求时间生成待双方确认的预约，并停止接收其他响应。确定选择该同学吗？',
      '确认选定候选人',
      { type: 'warning', confirmButtonText: '确定选定', cancelButtonText: '再看看' },
    );
  } catch {
    return;
  }
  void store.select(needId, responseId);
}
</script>

<style scoped>
.need-board { display: flex; flex-direction: column; gap: 14px; }
.need-board__head { display: flex; align-items: center; justify-content: space-between; gap: 12px; }
.need-board__head h2 { margin: 0; font-size: 20px; }
</style>
